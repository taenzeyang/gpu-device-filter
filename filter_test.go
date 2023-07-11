package device_filter

import (
	"testing"
)

func TestApply(t *testing.T) {
	f := &Flags{ConfigFile: "config.yaml", SelectedConfig: "test"}
	//TODO: Check flags

	t.Logf("Parsing config file...")
	spec, err := ParseConfigFile(f)
	if err != nil {
		t.Errorf("error parsing config file: %v", err)
	}

	t.Logf("Selecting specific config...")
	migConfig, err := GetSelectedMigConfig(f, spec)
	if err != nil {
		t.Errorf("error selecting config: %v", err)
	}

	t.Logf("Selecting specific devices...")
	deviceInfos, err := DeviceFilter(migConfig)
	if err != nil {
		t.Errorf("error selecting specific devices: %v", err)
	}

	if len(deviceInfos) == 0 {
		t.Logf("No device found")
	}
	for _, device := range deviceInfos {
		t.Logf("GPU %v: %v", device.Index, device.DeviceId)
	}

}
