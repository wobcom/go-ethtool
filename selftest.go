package ethtool

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"unsafe"
)

const (
	testIoctl = 0x0000001a /* execute NIC self-test. */
)

const (
	testFlagOffline              = (1 << 0) /* if set perform online and offline tests, otherwise only online tests */
	testFlagFailed               = (1 << 1) /* Driver set this flag if test fails */
	testFlagExternalLoopback     = (1 << 2) /* Application request to perform external loopback test */
	testFlagExternalLoopbackDone = (1 << 3) /* Driver performed the external loopback test */
)

type ethtoolTest struct {
	cmd      uint32
	flags    uint32
	reserved uint32
	length   uint32
	data     [maxNumStrings]uint64
}

// PerformSelftest performs an interface selftest.
// Note that this may cause the NIC to fail until reboot.
// TODO implement external loopback test (hint: requires special external wiring and some lab setup)
func (i *Interface) PerformSelftest(offline bool) (*TestResult, error) {
	flags := uint32(testFlagOffline)
	if offline {
		flags = 0
	}
	test := ethtoolTest{
		cmd:   testIoctl,
		flags: flags,
		data:  [maxNumStrings]uint64{},
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&test))); err != nil {
		return nil, errors.Wrapf(err, "Error performing ioctl testIoctl: %v", err)
	}

	return i.newTestResult(test)
}

// TestResult the return type for the PerformSelftest function
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
		Success:    result.flags&testFlagFailed == 0,
		TestValues: make(map[string]uint64),
	}
	resultNames, err := i.GetStringSet(StringSetTest)
	if err != nil {
		return nil, err
	}
	for i := uint32(0); i < result.length; i++ {
		testResult.TestValues[resultNames[i]] = result.data[i]
	}
	return testResult, nil
}
