package ethtool

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"unsafe"
)

const (
	ETHTOOL_TEST = 0x0000001a /* execute NIC self-test. */
)

const (
	ETH_TEST_FL_OFFLINE          = (1 << 0) /* if set perform online and offline tests, otherwise only online tests */
	ETH_TEST_FL_FAILED           = (1 << 1) /* Driver set this flag if test fails */
	ETH_TEST_FL_EXTERNAL_LB      = (1 << 2) /* Application request to perform external loopback test */
	ETH_TEST_FL_EXTERNAL_LB_DONE = (1 << 3) /* Driver performed the external loopback test */
)

type ethtoolTest struct {
	cmd      uint32
	flags    uint32
	reserved uint32
	length   uint32
	data     [MAX_GSTRINGS]uint64
}

// !! Use this only with absolute precaution - I have had to reboot systems after testing a NIC !!
//
// TODO implement external loopback test (hint: requires special external wiring and some lab setup)
func (i *Interface) PerformSelftest(offline bool) (*TestResult, error) {
	flags := uint32(ETH_TEST_FL_OFFLINE)
	if offline {
		flags = 0
	}
	test := ethtoolTest{
		cmd:   ETHTOOL_TEST,
		flags: flags,
		data:  [MAX_GSTRINGS]uint64{},
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&test))); err != nil {
		return nil, errors.Wrapf(err, "Error performing ioctl ETHTOOL_TEST: %v", err)
	}

	return i.newTestResult(test)
}

type TestResult struct {
	Success    bool
	TestValues map[string]uint64
}

func (t *TestResult) String() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "Test success: %t\n", t.Success)
	fmt.Fprintf(&builder, "Test results:\n")
	for resultName, resultValue := range t.TestValues {
		fmt.Fprintf(&builder, " * %s = %v\n", resultName, resultValue)
	}
	return builder.String()
}

func (i *Interface) newTestResult(result ethtoolTest) (*TestResult, error) {
	testResult := &TestResult{
		Success:    result.flags&ETH_TEST_FL_FAILED == 0,
		TestValues: make(map[string]uint64),
	}
	resultNames, err := i.GetStringSet(ETH_SS_TEST)
	if err != nil {
		return nil, err
	}
	for i := uint32(0); i < result.length; i++ {
		testResult.TestValues[resultNames[i]] = result.data[i]
	}
	return testResult, nil
}
