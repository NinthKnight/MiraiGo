// Code generated by Proto2Go.Net. Only used for MiraiGo. DO NOT EDIT.
// Source: oidb0x769.proto

package oidb0x769

import (
	"github.com/pkg/errors"
	"go.dedis.ch/protobuf"
)

type (
	CPU struct {
		Model     *string `protobuf:"1,opt"`
		Cores     *uint32 `protobuf:"2,opt"`
		Frequency *uint32 `protobuf:"3,opt"`
	}

	Camera struct {
		Primary   *uint64 `protobuf:"1,opt"`
		Secondary *uint64 `protobuf:"2,opt"`
		Flash     *bool   `protobuf:"3,opt"`
	}

	ConfigSeq struct {
		Type    *uint32 `protobuf:"1,opt"`
		Version *uint32 `protobuf:"2,opt"`
	}

	Content struct {
		TaskId   *uint32 `protobuf:"1,opt"`
		Compress *uint32 `protobuf:"2,opt"`
		Content  []byte  `protobuf:"10,opt"`
	}

	DeviceInfo struct {
		Brand   *string  `protobuf:"1,opt"`
		Model   *string  `protobuf:"2,opt"`
		Os      *OS      `protobuf:"3,opt"`
		Cpu     *CPU     `protobuf:"4,opt"`
		Memory  *Memory  `protobuf:"5,opt"`
		Storage *Storage `protobuf:"6,opt"`
		Screen  *Screen  `protobuf:"7,opt"`
		Camera  *Camera  `protobuf:"8,opt"`
	}

	Memory struct {
		Total   *uint64 `protobuf:"1,opt"`
		Process *uint64 `protobuf:"2,opt"`
	}

	OS struct {
		Type    *uint32 `protobuf:"1,opt"`
		Version *string `protobuf:"2,opt"`
		Sdk     *string `protobuf:"3,opt"`
		Kernel  *string `protobuf:"4,opt"`
		Rom     *string `protobuf:"5,opt"`
	}

	QueryUinPackageUsageReq struct {
		Type        *uint32 `protobuf:"1,opt"`
		UinFileSize *uint64 `protobuf:"2,opt"`
	}

	QueryUinPackageUsageRsp struct {
		Status             *uint32               `protobuf:"1,opt"`
		LeftUinNum         *uint64               `protobuf:"2,opt"`
		MaxUinNum          *uint64               `protobuf:"3,opt"`
		Proportion         *uint32               `protobuf:"4,opt"`
		UinPackageUsedList []*UinPackageUsedInfo `protobuf:"10"`
	}

	ReqBody struct {
		ConfigList              []*ConfigSeq             `protobuf:"1"`
		DeviceInfo              *DeviceInfo              `protobuf:"2,opt"`
		Info                    *string                  `protobuf:"3,opt"`
		Province                *string                  `protobuf:"4,opt"`
		City                    *string                  `protobuf:"5,opt"`
		ReqDebugMsg             *int32                   `protobuf:"6,opt"`
		QueryUinPackageUsageReq *QueryUinPackageUsageReq `protobuf:"101,opt"`
	}

	RspBody struct {
		Result                  *uint32                  `protobuf:"1,opt"`
		ConfigList              []*ConfigSeq             `protobuf:"2"`
		QueryUinPackageUsageRsp *QueryUinPackageUsageRsp `protobuf:"101,opt"`
	}

	Screen struct {
		Model      *string `protobuf:"1,opt"`
		Width      *uint32 `protobuf:"2,opt"`
		Height     *uint32 `protobuf:"3,opt"`
		Dpi        *uint32 `protobuf:"4,opt"`
		MultiTouch *bool   `protobuf:"5,opt"`
	}

	Storage struct {
		Builtin  *uint64 `protobuf:"1,opt"`
		External *uint64 `protobuf:"2,opt"`
	}

	UinPackageUsedInfo struct {
		RuleId *uint32 `protobuf:"1,opt"`
		Author *string `protobuf:"2,opt"`
		Url    *string `protobuf:"3,opt"`
		UinNum *uint64 `protobuf:"4,opt"`
	}
)

func (x *CPU) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *CPU) GetCores() uint32 {
	if x != nil && x.Cores != nil {
		return *x.Cores
	}
	return 0
}

func (x *CPU) GetFrequency() uint32 {
	if x != nil && x.Frequency != nil {
		return *x.Frequency
	}
	return 0
}

func (x *CPU) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *Camera) GetPrimary() uint64 {
	if x != nil && x.Primary != nil {
		return *x.Primary
	}
	return 0
}

func (x *Camera) GetSecondary() uint64 {
	if x != nil && x.Secondary != nil {
		return *x.Secondary
	}
	return 0
}

func (x *Camera) GetFlash() bool {
	if x != nil && x.Flash != nil {
		return *x.Flash
	}
	return false
}

func (x *Camera) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *ConfigSeq) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *ConfigSeq) GetVersion() uint32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *ConfigSeq) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *Content) GetTaskId() uint32 {
	if x != nil && x.TaskId != nil {
		return *x.TaskId
	}
	return 0
}

func (x *Content) GetCompress() uint32 {
	if x != nil && x.Compress != nil {
		return *x.Compress
	}
	return 0
}

func (x *Content) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *DeviceInfo) GetBrand() string {
	if x != nil && x.Brand != nil {
		return *x.Brand
	}
	return ""
}

func (x *DeviceInfo) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *DeviceInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *Memory) GetTotal() uint64 {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return 0
}

func (x *Memory) GetProcess() uint64 {
	if x != nil && x.Process != nil {
		return *x.Process
	}
	return 0
}

func (x *Memory) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *OS) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *OS) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *OS) GetSdk() string {
	if x != nil && x.Sdk != nil {
		return *x.Sdk
	}
	return ""
}

func (x *OS) GetKernel() string {
	if x != nil && x.Kernel != nil {
		return *x.Kernel
	}
	return ""
}

func (x *OS) GetRom() string {
	if x != nil && x.Rom != nil {
		return *x.Rom
	}
	return ""
}

func (x *OS) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *QueryUinPackageUsageReq) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *QueryUinPackageUsageReq) GetUinFileSize() uint64 {
	if x != nil && x.UinFileSize != nil {
		return *x.UinFileSize
	}
	return 0
}

func (x *QueryUinPackageUsageReq) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *QueryUinPackageUsageRsp) GetStatus() uint32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) GetLeftUinNum() uint64 {
	if x != nil && x.LeftUinNum != nil {
		return *x.LeftUinNum
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) GetMaxUinNum() uint64 {
	if x != nil && x.MaxUinNum != nil {
		return *x.MaxUinNum
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) GetProportion() uint32 {
	if x != nil && x.Proportion != nil {
		return *x.Proportion
	}
	return 0
}

func (x *QueryUinPackageUsageRsp) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *ReqBody) GetInfo() string {
	if x != nil && x.Info != nil {
		return *x.Info
	}
	return ""
}

func (x *ReqBody) GetProvince() string {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return ""
}

func (x *ReqBody) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *ReqBody) GetReqDebugMsg() int32 {
	if x != nil && x.ReqDebugMsg != nil {
		return *x.ReqDebugMsg
	}
	return 0
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *RspBody) GetResult() uint32 {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return 0
}

func (x *RspBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *Screen) GetModel() string {
	if x != nil && x.Model != nil {
		return *x.Model
	}
	return ""
}

func (x *Screen) GetWidth() uint32 {
	if x != nil && x.Width != nil {
		return *x.Width
	}
	return 0
}

func (x *Screen) GetHeight() uint32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

func (x *Screen) GetDpi() uint32 {
	if x != nil && x.Dpi != nil {
		return *x.Dpi
	}
	return 0
}

func (x *Screen) GetMultiTouch() bool {
	if x != nil && x.MultiTouch != nil {
		return *x.MultiTouch
	}
	return false
}

func (x *Screen) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *Storage) GetBuiltin() uint64 {
	if x != nil && x.Builtin != nil {
		return *x.Builtin
	}
	return 0
}

func (x *Storage) GetExternal() uint64 {
	if x != nil && x.External != nil {
		return *x.External
	}
	return 0
}

func (x *Storage) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *UinPackageUsedInfo) GetRuleId() uint32 {
	if x != nil && x.RuleId != nil {
		return *x.RuleId
	}
	return 0
}

func (x *UinPackageUsedInfo) GetAuthor() string {
	if x != nil && x.Author != nil {
		return *x.Author
	}
	return ""
}

func (x *UinPackageUsedInfo) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *UinPackageUsedInfo) GetUinNum() uint64 {
	if x != nil && x.UinNum != nil {
		return *x.UinNum
	}
	return 0
}

func (x *UinPackageUsedInfo) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}
