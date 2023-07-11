package types

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	nvdev "gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvlib/device"
	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvml"
)

var nvmllib nvml.Interface
var nvdevlib nvdev.Interface

func tryNvmlShutdown(nvmlLib nvml.Interface) {
	ret := nvmlLib.Shutdown()
	if ret != nvml.SUCCESS {
		log.Warnf("Error shutting down NVML: %v", ret)
	}
}

func nvdevAssertValidMigProfileFormat(profile string) error {
	if nvmllib == nil {
		nvmllib = nvml.New()
	}
	if nvdevlib == nil {
		nvdevlib = nvdev.New(nvdev.WithNvml(nvmllib))
	}

	return nvdevlib.AssertValidMigProfileFormat(profile)
}

func nvdevParseMigProfile(profile string) (nvdev.MigProfile, error) {
	if nvmllib == nil {
		nvmllib = nvml.New()
	}
	if nvdevlib == nil {
		nvdevlib = nvdev.New(nvdev.WithNvml(nvmllib))
	}

	ret := nvmllib.Init()
	if ret != nvml.SUCCESS {
		return nil, fmt.Errorf("error initializing NVML: %v", ret)
	}
	defer tryNvmlShutdown(nvmllib)

	mp, err := nvdevlib.ParseMigProfile(profile)
	if err != nil {
		return nil, err
	}

	return mp, nil
}
