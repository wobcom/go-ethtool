package ethtool

import "unsafe"

// identify the NIC
const identifyNicIoctl = 0x0000001c /* identify the NIC */

// Identify makes the interface LED blink for the given time in seconds
func (i *Interface) Identify(time uint32) error {
    cmd := ethtoolArbitraryCommand{
        cmd:   identifyNicIoctl,
        value: time,
    }   
    return i.performIoctl(uintptr(unsafe.Pointer(&cmd)))
}
