package back

import "github.com/Alice00021/test_api/internal/entity"

type ReagentType string

type Address string

type Container struct {
	Address     Address
	ReagentType ReagentType
	Volume      int64
}
type Command struct {
	entity.Entity
	Name             string
	SystemName       string
	Reagent          ReagentType
	AverageTime      int64
	VolumeWaste      int64
	VolumeDriveFluid int64
	VolumeContainer  int64
	DefaultAddress   Address
}
