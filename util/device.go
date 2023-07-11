package util

import (
	"fmt"

	"github.com/taenzeyang/gpu-device-filter/nvml"
	"github.com/taenzeyang/gpu-device-filter/types"

	"gitlab.com/nvidia/cloud-native/go-nvlib/pkg/nvpci"
)

func GetGPUDeviceIDs() ([]types.DeviceID, error) {
	nvidiaModuleLoaded, err := IsNvidiaModuleLoaded()
	if err != nil {
		return nil, fmt.Errorf("error checking if nvidia module loaded: %v", err)
	}
	if nvidiaModuleLoaded {
		return nvmlGetGPUDeviceIDs()
	}
	return pciGetGPUDeviceIDs()
}

func pciVisitGPUs(visit func(*nvpci.NvidiaPCIDevice) error) error {
	nvpci := nvpci.New()
	gpus, err := nvpci.GetGPUs()
	if err != nil {
		return fmt.Errorf("error enumerating GPUs: %v", err)
	}
	for _, gpu := range gpus {
		err := visit(gpu)
		if err != nil {
			return err
		}
	}
	return nil
}

func pciGetGPUDeviceIDs() ([]types.DeviceID, error) {
	var ids []types.DeviceID
	err := pciVisitGPUs(func(gpu *nvpci.NvidiaPCIDevice) error {
		ids = append(ids, types.NewDeviceID(gpu.Device, gpu.Vendor))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func nvmlGetGPUDeviceIDs() ([]types.DeviceID, error) {
	nvmlLib := nvml.New()
	err := NvmlInit(nvmlLib)
	if err != nil {
		return nil, fmt.Errorf("error initializing NVML: %v", err)
	}
	defer TryNvmlShutdown(nvmlLib)

	var ids []types.DeviceID
	err = pciVisitGPUs(func(gpu *nvpci.NvidiaPCIDevice) error {
		_, ret := nvmlLib.DeviceGetHandleByPciBusId(gpu.Address)
		if ret.Value() != nvml.SUCCESS {
			return nil
		}

		ids = append(ids, types.NewDeviceID(gpu.Device, gpu.Vendor))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ids, nil
}
