// Code generated by Proto2Go.Net. Only used for MiraiGo. DO NOT EDIT.
// Source: oidb0xdad.proto

package oidb0xdad

import (
	"github.com/pkg/errors"
	"go.dedis.ch/protobuf"
)

type (
	ReqBody struct {
		Client    *int64    `protobuf:"1,opt"`
		ProductId *uint64   `protobuf:"2,opt"`
		Amount    *int64    `protobuf:"3,opt"`
		ToUin     *uint64   `protobuf:"4,opt"`
		Gc        *uint64   `protobuf:"5,opt"`
		Ip        *string   `protobuf:"6,opt"`
		Version   *string   `protobuf:"7,opt"`
		Sig       *LoginSig `protobuf:"8,opt"`
	}

	LoginSig struct {
		Type  *uint32 `protobuf:"1,opt"`
		Sig   []byte  `protobuf:"2,opt"`
		Appid *uint32 `protobuf:"3,opt"`
	}
)

func (x *ReqBody) GetClient() int64 {
	if x != nil && x.Client != nil {
		return *x.Client
	}
	return 0
}

func (x *ReqBody) GetProductId() uint64 {
	if x != nil && x.ProductId != nil {
		return *x.ProductId
	}
	return 0
}

func (x *ReqBody) GetAmount() int64 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *ReqBody) GetToUin() uint64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *ReqBody) GetGc() uint64 {
	if x != nil && x.Gc != nil {
		return *x.Gc
	}
	return 0
}

func (x *ReqBody) GetIp() string {
	if x != nil && x.Ip != nil {
		return *x.Ip
	}
	return ""
}

func (x *ReqBody) GetVersion() string {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return ""
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}

func (x *LoginSig) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *LoginSig) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

func (x *LoginSig) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}
