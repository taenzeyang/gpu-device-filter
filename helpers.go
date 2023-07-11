package device_filter

import (
	"github.com/taenzeyang/gpu-device-filter/types"
)

// MatchesDeviceFilter checks a 'MigConfigSpec' to see if its device filter matches the provided 'deviceID'.
func (ms *MigConfigSpec) MatchesDeviceFilter(deviceID types.DeviceID) bool {
	var deviceFilter []string
	switch df := ms.DeviceFilter.(type) {
	case string:
		if df != "" {
			deviceFilter = append(deviceFilter, df)
		}
	case []string:
		deviceFilter = df
	}

	if len(deviceFilter) == 0 {
		return true
	}

	for _, df := range deviceFilter {
		newDeviceID, _ := types.NewDeviceIDFromString(df)
		if newDeviceID == deviceID {
			return true
		}
	}

	return false
}

// MatchesAllDevices checks a 'MigConfigSpec' to see if it matches on 'all' devices.
func (ms *MigConfigSpec) MatchesAllDevices() bool {
	switch devices := ms.Devices.(type) {
	case string:
		return devices == "all"
	}
	return false
}

// MatchesDevices checks a 'MigConfigSpec' to see if it matches on a device at the specified 'index'.
func (ms *MigConfigSpec) MatchesDevices(index int) bool {
	switch devices := ms.Devices.(type) {
	case []int:
		for _, d := range devices {
			if index == d {
				return true
			}
		}
	}
	return ms.MatchesAllDevices()
}
