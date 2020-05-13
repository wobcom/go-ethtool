package ethtool

import (
	"bytes"
	"github.com/pkg/errors"
	"unsafe"
)

const (
	maxNumStrings   = 256
	maxStringLength = 32

	// Get string set info
	getStringSetInfoIoctl = 0x00000037
	// Get specified string set
	getStringSet = 0x0000001b
)

// StringSet a set of strings used by some ethtool ioctls
type StringSet uint32

const (
	// StringSetTest Self-test result names
	StringSetTest StringSet = 0
	// StringSetStats Statistic names
	StringSetStats StringSet = 1
	// StringSetPrivFlag Driver private flag names
	StringSetPrivFlag StringSet = 2
	// StringSetTupleFilters Depracted
	StringSetTupleFilters StringSet = 3
	// StringSetFeatures Device feature names
	StringSetFeatures StringSet = 4
	// StringSetRssHashFuncs RSS hash function names
	StringSetRssHashFuncs StringSet = 5
	// StringSetTunables Not documented in the kernel's source code
	StringSetTunables StringSet = 6
	// StringSetPhyStats PHY Statistic names
	StringSetPhyStats StringSet = 7
	// StringSetPhyTunables PPHY tunables names
	StringSetPhyTunables StringSet = 8
	// StringSetLinkModes link mode names
	StringSetLinkModes StringSet = 9
	// StringSetMsgClasses debug message class names
	StringSetMsgClasses StringSet = 10
	// StringSetWolModes wake-on-lan modes
	StringSetWolModes StringSet = 11
	// StringSetSofTimestamping SOF_TIMESTAMPING_* flags
	StringSetSofTimestamping StringSet = 12
	// StringSetTimestampTxTypes timestamping Tx types
	StringSetTimestampTxTypes StringSet = 13
	// StringSetTimestampRxFilters timestamping Rx filters
	StringSetTimestampRxFilters StringSet = 14
)

var existingStringSets = []StringSet{
	StringSetTest, StringSetStats, StringSetPrivFlag, StringSetTupleFilters, StringSetTupleFilters, StringSetFeatures,
	StringSetRssHashFuncs, StringSetTunables, StringSetPhyStats, StringSetPhyTunables, StringSetLinkModes,
	StringSetMsgClasses, StringSetWolModes, StringSetSofTimestamping, StringSetTimestampTxTypes, StringSetTimestampRxFilters}

type ethtoolGStrings struct {
	cmd       uint32
	stringSet uint32
	length    uint32
	data      [maxNumStrings * maxStringLength]byte
}

type ethtoolSsetInfo struct {
	cmd      uint32
	reserved uint32
	ssetMask uint64
	data     uint32
}

// GetStringSet retrieves the given StringSet and returns a string slice
func (i *Interface) GetStringSet(set StringSet) ([]string, error) {
	length, err := i.GetStringSetLength(set)
	if err != nil {
		return nil, errors.Wrapf(err, "Error retrieving string set length: %v", err)
	}

	if length == 0 {
		return []string{}, nil
	}

	gStrings := ethtoolGStrings{
		cmd:       getStringSet,
		stringSet: uint32(set),
		length:    length,
		data:      [maxNumStrings * maxStringLength]byte{},
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&gStrings))); err != nil {
		return nil, errors.Wrapf(err, "Error performing ioctl getStringSet: %v", err)
	}
	ret := make([]string, int(gStrings.length))
	for i := 0; i < int(gStrings.length); i++ {
		b := gStrings.data[i*maxStringLength : (i+1)*maxStringLength]
		ret[i] = string(bytes.Trim(b, "\x00"))
	}
	return ret, nil
}

// GetStringSetLength gets the length of a given StringSet
func (i *Interface) GetStringSetLength(set StringSet) (uint32, error) {
	setInfo := ethtoolSsetInfo{
		cmd:      getStringSetInfoIoctl,
		ssetMask: 1 << uint32(set),
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&setInfo))); err != nil {
		return 0, errors.Wrapf(err, "Error performing ioctl getStringSetInfoIoctl")
	}
	return uint32(setInfo.data), nil
}
