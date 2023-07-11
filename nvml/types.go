package nvml

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

type Interface interface {
	Init() Return
	Shutdown() Return
	SystemGetNVMLVersion() (string, Return)
	DeviceGetCount() (int, Return)
	DeviceGetHandleByIndex(Index int) (Device, Return)
	DeviceGetHandleByUUID(UUID string) (Device, Return)
	DeviceGetHandleByPciBusId(busID string) (Device, Return)
}

type Device interface {
	GetIndex() (int, Return)
	GetUUID() (string, Return)
	GetMemoryInfo() (Memory, Return)
	GetPciInfo() (PciInfo, Return)
	SetMigMode(Mode int) (Return, Return)
	GetMigMode() (int, int, Return)
	GetGpuInstanceProfileInfo(Profile int) (GpuInstanceProfileInfo, Return)
	CreateGpuInstance(Info *GpuInstanceProfileInfo) (GpuInstance, Return)
	CreateGpuInstanceWithPlacement(Info *GpuInstanceProfileInfo, Placement *GpuInstancePlacement) (GpuInstance, Return)
	GetGpuInstances(Info *GpuInstanceProfileInfo) ([]GpuInstance, Return)
}

type GpuInstance interface {
	GetInfo() (GpuInstanceInfo, Return)
	GetComputeInstanceProfileInfo(Profile int, EngProfile int) (ComputeInstanceProfileInfo, Return)
	CreateComputeInstance(Info *ComputeInstanceProfileInfo) (ComputeInstance, Return)
	GetComputeInstances(Info *ComputeInstanceProfileInfo) ([]ComputeInstance, Return)
	Destroy() Return
}

type ComputeInstance interface {
	GetInfo() (ComputeInstanceInfo, Return)
	Destroy() Return
}

type GpuInstanceInfo struct {
	Device    Device
	Id        uint32
	ProfileId uint32
	Placement GpuInstancePlacement
}

type ComputeInstanceInfo struct {
	Device      Device
	GpuInstance GpuInstance
	Id          uint32
	ProfileId   uint32
	Placement   ComputeInstancePlacement
}

type Memory nvml.Memory
type PciInfo nvml.PciInfo
type GpuInstanceProfileInfo nvml.GpuInstanceProfileInfo
type GpuInstancePlacement nvml.GpuInstancePlacement
type ComputeInstanceProfileInfo nvml.ComputeInstanceProfileInfo
type ComputeInstancePlacement nvml.ComputeInstancePlacement
