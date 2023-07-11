package types

import (
	"github.com/taenzeyang/gpu-device-filter/nvml"
)

// MigState stores the MIG state for a set of GPUs.
type MigState struct {
	Devices []DeviceState
}

// DeviceState stores the MIG state for a specific GPU.
type DeviceState struct {
	UUID         string
	GpuInstances []GpuInstanceState
}

// GpuInstanceState stores the MIG state for a specific GPUInstance.
type GpuInstanceState struct {
	ProfileID        int
	Placement        nvml.GpuInstancePlacement
	ComputeInstances []ComputeInstanceState
}

// ComputeInstanceState stores the MIG state for a specific ComputeInstance.
type ComputeInstanceState struct {
	ProfileID    int
	EngProfileID int
}
