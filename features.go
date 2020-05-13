package ethtool

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"unsafe"
)

const (
	// Get device offload settings
	getFeatures = 0x0000003a
)

const (
	maxFeatureBlocks = (maxNumStrings + 32 - 1) / 32
)

type ethtoolGetFeaturesBlock struct {
	available    uint32
	requested    uint32
	active       uint32
	neverChanged uint32
}

type ethtoolGfeatures struct {
	cmd    uint32
	size   uint32
	blocks [maxFeatureBlocks]ethtoolGetFeaturesBlock
}

// FeatureList maps the name of a network interfaces feature to a FeatureStatus
type FeatureList map[string]FeatureStatus

// FeatureStatus describes the status of network interface feature
type FeatureStatus struct {
	Available    bool
	Active       bool
	NeverChanged bool
}

// GetFeatures returns a FeatureList of features supported by the NIC
func (i *Interface) GetFeatures() (FeatureList, error) {
	names, err := i.GetStringSet(StringSetFeatures)
	if err != nil {
		return FeatureList{}, errors.Wrapf(err, "Could not retrieve list of feature names: %v", err)
	}

	features := ethtoolGfeatures{
		cmd:  getFeatures,
		size: uint32((len(names) + 31) / 32),
	}
	if err := i.performIoctl(uintptr(unsafe.Pointer(&features))); err != nil {
		return FeatureList{}, errors.Wrapf(err, "Error running ioctl getFeatures")
	}

	ret := make(FeatureList)
	for index, name := range names {
		val, err := getFeatureBit(features, index)
		if err != nil {
			return FeatureList{}, errors.Wrap(err, "Failed to retrieve feature information. ")
		}
		ret[name] = val
	}
	return ret, nil
}

func (f FeatureList) String() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "Feature List (active, available, never changed):\n")
	for key, status := range map[string]FeatureStatus(f) {
		fmt.Fprintf(&builder, " * %s: %t, %t, %t\n", key, status.Active, status.Available, status.NeverChanged)
	}
	return builder.String()
}

func getFeatureBit(features ethtoolGfeatures, index int) (FeatureStatus, error) {
	if index/32 > int(features.size) || index/32 > maxFeatureBlocks {
		return FeatureStatus{false, false, false}, errors.New(fmt.Sprintf("Index %d out of bound for retrieved feature list (size = %d * 32)", index, features.size))
	}
	return FeatureStatus{
		Available:    features.blocks[index/32].available&(1<<(index%32)) > 0,
		Active:       features.blocks[index/32].active&(1<<(index%32)) > 0,
		NeverChanged: features.blocks[index/32].neverChanged&(1<<(index%32)) > 0,
	}, nil
}
