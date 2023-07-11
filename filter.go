package device_filter

import (
	"bufio"
	"fmt"
	"os"

	"github.com/taenzeyang/gpu-device-filter/nvml"
	"github.com/taenzeyang/gpu-device-filter/util"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/yaml"
)

type Flags struct {
	ConfigFile     string
	SelectedConfig string
}

type DeviceInfo struct {
	Index    int
	DeviceId string
}

func ParseConfigFile(f *Flags) (*Spec, error) {
	var err error
	var configYaml []byte

	if f.ConfigFile == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			configYaml = append(configYaml, scanner.Bytes()...)
			configYaml = append(configYaml, '\n')
		}
	} else {
		configYaml, err = os.ReadFile(f.ConfigFile)
		if err != nil {
			return nil, fmt.Errorf("read error: %v", err)
		}
	}

	var spec Spec
	err = yaml.Unmarshal(configYaml, &spec)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}
	return &spec, nil
}

func GetSelectedMigConfig(f *Flags, spec *Spec) (MigConfigSpecSlice, error) {
	if len(spec.MigConfigs) > 1 && f.SelectedConfig == "" {
		return nil, fmt.Errorf("missing required flag 'selected-config' when more than one config available")
	}

	if len(spec.MigConfigs) == 1 && f.SelectedConfig == "" {
		for c := range spec.MigConfigs {
			f.SelectedConfig = c
		}
	}

	if _, exists := spec.MigConfigs[f.SelectedConfig]; !exists {
		return nil, fmt.Errorf("selected mig-config not present: %v", f.SelectedConfig)
	}

	return spec.MigConfigs[f.SelectedConfig], nil
}

func DeviceFilter(migConfig MigConfigSpecSlice) ([]DeviceInfo, error) {
	n := nvml.New()
	err := util.NvmlInit(n)
	if err != nil {
		return nil, fmt.Errorf("error initializing NVML: %v", err)
	}
	defer util.TryNvmlShutdown(n)

	deviceIDs, err := util.GetGPUDeviceIDs()
	if err != nil {
		return nil, fmt.Errorf("Error enumerating GPU device IDs: %v", err)
	}

	deviceInfos := make([]DeviceInfo, 0)
	for _, mc := range migConfig {
		if mc.DeviceFilter == nil {
			log.Infof("Walking Config for (devices=%v)", mc.Devices)
		} else {
			log.Infof("Walking Config for (device-filter=%v, devices=%v)", mc.DeviceFilter, mc.Devices)
		}

		for i, deviceID := range deviceIDs {
			if !mc.MatchesDeviceFilter(deviceID) {
				continue
			}

			if !mc.MatchesDevices(i) {
				continue
			}

			deviceInfos = append(deviceInfos, DeviceInfo{
				Index:    i,
				DeviceId: deviceID.String(),
			})
		}
	}

	return deviceInfos, nil
}

func Apply() ([]DeviceInfo, error) {
	f := &Flags{ConfigFile: "/var/run/config.yaml", SelectedConfig: "test"}
	//TODO: Check flags

	log.Infof("Parsing config file...")
	spec, err := ParseConfigFile(f)
	if err != nil {
		log.Errorf("error parsing config file: %v", err)
	}

	log.Infof("Selecting specific config...")
	migConfig, err := GetSelectedMigConfig(f, spec)
	if err != nil {
		log.Errorf("error selecting config: %v", err)
	}

	log.Infof("Selecting specific devices...")
	deviceInfos, err := DeviceFilter(migConfig)
	if err != nil {
		log.Errorf("error selecting specific devices: %v", err)
	}

	return deviceInfos, nil
}
