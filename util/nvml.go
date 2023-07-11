package util

import (
	"fmt"
	"os"
	"strings"

	"github.com/taenzeyang/gpu-device-filter/nvml"

	log "github.com/sirupsen/logrus"
)

func IsNvidiaModuleLoaded() (bool, error) {
	modules, err := os.ReadFile("/proc/modules")
	if err != nil {
		return false, fmt.Errorf("unable to read /proc/modules: %v", err)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(modules)), "\n") {
		fields := strings.Fields(line)
		if fields[0] == "nvidia" {
			return true, nil
		}
	}
	return false, nil
}

func NvmlInit(nvmlLib nvml.Interface) error {
	if nvmlLib == nil {
		nvmlLib = nvml.New()
	}
	ret := nvmlLib.Init()
	if ret.Value() != nvml.SUCCESS {
		return ret
	}
	return nil
}

func TryNvmlShutdown(nvmlLib nvml.Interface) {
	if nvmlLib == nil {
		nvmlLib = nvml.New()
	}
	ret := nvmlLib.Shutdown()
	if ret.Value() != nvml.SUCCESS {
		log.Warnf("error shutting down NVML: %v", ret)
	}
}
