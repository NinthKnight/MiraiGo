// Code generated by Proto2Go.Net. Only used for MiraiGo. DO NOT EDIT.
// Source: oidb0xed3.proto

package oidb0xed3

import (
	"github.com/pkg/errors"
	"go.dedis.ch/protobuf"
)

type (
	ReqBody struct {
		ToUin     *int64 `protobuf:"1,opt"`
		GroupCode *int64 `protobuf:"2,opt"`
		MsgSeq    *int32 `protobuf:"3,opt"`
		MsgRand   *int32 `protobuf:"4,opt"`
		AioUin    *int64 `protobuf:"5,opt"`
	}
)

func (x *ReqBody) GetToUin() int64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *ReqBody) GetGroupCode() int64 {
	if x != nil && x.GroupCode != nil {
		return *x.GroupCode
	}
	return 0
}

func (x *ReqBody) GetMsgSeq() int32 {
	if x != nil && x.MsgSeq != nil {
		return *x.MsgSeq
	}
	return 0
}

func (x *ReqBody) GetMsgRand() int32 {
	if x != nil && x.MsgRand != nil {
		return *x.MsgRand
	}
	return 0
}

func (x *ReqBody) GetAioUin() int64 {
	if x != nil && x.AioUin != nil {
		return *x.AioUin
	}
	return 0
}

func (x *ReqBody) Marshal() ([]byte, error) {
	if x == nil {
		return nil, errors.New("nil pointer error")
	}
	return protobuf.Encode(x)
}
