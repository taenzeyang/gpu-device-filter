package types

import (
	"fmt"
	"strconv"
)

// DeviceID represents a GPU Device ID as read from a GPUs PCIe config space.
type DeviceID uint32

// NewDeviceID constructs a new 'DeviceID' from the device and vendor values pulled from a GPUs PCIe config space.
func NewDeviceID(device, vendor uint16) DeviceID {
	return DeviceID((uint32(device) << 16) | uint32(vendor))
}

// NewDeviceIDFromString constructs a 'DeviceID' from its string representation.
func NewDeviceIDFromString(str string) (DeviceID, error) {
	deviceID, err := strconv.ParseInt(str, 0, 32)
	if err != nil {
		return 0, fmt.Errorf("unable to create DeviceID from string '%v': %v", str, err)
	}
	return DeviceID(deviceID), nil
}

// String returns a 'DeviceID' as a string.
func (d DeviceID) String() string {
	return fmt.Sprintf("0x%X", uint32(d))
}

// GetVendor returns the 'vendor' portion of a 'DeviceID'.
func (d DeviceID) GetVendor() uint16 {
	return uint16(d)
}

// GetDevice returns the 'device' portion of a 'DeviceID'.
func (d DeviceID) GetDevice() uint16 {
	return uint16(d >> 16)
}
