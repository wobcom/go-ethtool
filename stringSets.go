package ethtool

import (
	"bytes"
	"github.com/pkg/errors"
	"unsafe"
)

const (
	MAX_GSTRINGS    = 256
	ETH_GSTRING_LEN = 32

	ETHTOOL_GSSET_INFO = 0x00000037 /* Get string set info */
	ETHTOOL_GSTRINGS   = 0x0000001b /* get specified string set */
)

type StringSet uint32

const (
	ETH_SS_TEST             StringSet = 0
	ETH_SS_STATS            StringSet = 1
	ETH_SS_PRIV_FLAGS       StringSet = 2
	ETH_SS_NTUPLE_FILTERS   StringSet = 3
	ETH_SS_FEATURES         StringSet = 4
	ETH_SS_RSS_HASH_FUNCS   StringSet = 5
	ETH_SS_TUNABLES         StringSet = 6
	ETH_SS_PHY_STATS        StringSet = 7
	ETH_SS_PHY_TUNABLES     StringSet = 8
	ETH_SS_LINK_MODES       StringSet = 9
	ETH_SS_MSG_CLASSES      StringSet = 10
	ETH_SS_WOL_MODES        StringSet = 11
	ETH_SS_SOF_TIMESTAMPING StringSet = 12
	ETH_SS_TS_TX_TYPES      StringSet = 13
	ETH_SS_TS_RX_FILTERS    StringSet = 14
)

var existingStringSets = []StringSet{
	ETH_SS_TEST, ETH_SS_STATS, ETH_SS_PRIV_FLAGS, ETH_SS_NTUPLE_FILTERS, ETH_SS_NTUPLE_FILTERS, ETH_SS_FEATURES,
	ETH_SS_RSS_HASH_FUNCS, ETH_SS_TUNABLES, ETH_SS_PHY_STATS, ETH_SS_PHY_TUNABLES, ETH_SS_LINK_MODES,
	ETH_SS_MSG_CLASSES, ETH_SS_WOL_MODES, ETH_SS_SOF_TIMESTAMPING, ETH_SS_TS_TX_TYPES, ETH_SS_TS_RX_FILTERS}

type ethtoolGStrings struct {
	cmd        uint32
	string_set uint32
	length     uint32
	data       [MAX_GSTRINGS * ETH_GSTRING_LEN]byte
}

type ethtoolSsetInfo struct {
	cmd       uint32
	reserved  uint32
	sset_mask uint64
	data      uint32
}

func (i *Interface) GetStringSet(set StringSet) ([]string, error) {
	length, err := i.GetStringSetLength(set)
	if err != nil {
		return nil, errors.Wrapf(err, "Error retrieving string set length: %v", err)
	}

	if length == 0 {
		return []string{}, nil
	}

	gStrings := ethtoolGStrings{
		cmd:        ETHTOOL_GSTRINGS,
		string_set: uint32(set),
		length:     length,
		data:       [MAX_GSTRINGS * ETH_GSTRING_LEN]byte{},
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&gStrings))); err != nil {
		return nil, errors.Wrapf(err, "Error performing ioctl ETHTOOL_GSTRINGS: %v", err)
	}
	ret := make([]string, int(gStrings.length))
	for i := 0; i < int(gStrings.length); i++ {
		b := gStrings.data[i*ETH_GSTRING_LEN : (i+1)*ETH_GSTRING_LEN]
		ret[i] = string(bytes.Trim(b, "\x00"))
	}
	return ret, nil
}

func (i *Interface) GetStringSetLength(set StringSet) (uint32, error) {
	setInfo := ethtoolSsetInfo{
		cmd:       ETHTOOL_GSSET_INFO,
		sset_mask: 1 << uint32(set),
	}

	if err := i.performIoctl(uintptr(unsafe.Pointer(&setInfo))); err != nil {
		return 0, errors.Wrapf(err, "Error performing ioctl ETHTOOL_GSSET_INFO")
	}
	return uint32(setInfo.data), nil
}
