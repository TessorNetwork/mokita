// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mokita/gamm/pool-models/stableswap/tx.proto

package stableswap

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ===================== MsgCreatePool
type MsgCreateStableswapPool struct {
	Sender                  string                                   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolParams              *PoolParams                              `protobuf:"bytes,2,opt,name=pool_params,json=poolParams,proto3" json:"pool_params,omitempty" yaml:"pool_params"`
	InitialPoolLiquidity    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=initial_pool_liquidity,json=initialPoolLiquidity,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_pool_liquidity"`
	ScalingFactors          []uint64                                 `protobuf:"varint,4,rep,packed,name=scaling_factors,json=scalingFactors,proto3" json:"scaling_factors,omitempty" yaml:"stableswap_scaling_factor"`
	FuturePoolGovernor      string                                   `protobuf:"bytes,5,opt,name=future_pool_governor,json=futurePoolGovernor,proto3" json:"future_pool_governor,omitempty" yaml:"future_pool_governor"`
	ScalingFactorController string                                   `protobuf:"bytes,6,opt,name=scaling_factor_controller,json=scalingFactorController,proto3" json:"scaling_factor_controller,omitempty" yaml:"scaling_factor_controller"`
}

func (m *MsgCreateStableswapPool) Reset()         { *m = MsgCreateStableswapPool{} }
func (m *MsgCreateStableswapPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStableswapPool) ProtoMessage()    {}
func (*MsgCreateStableswapPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{0}
}
func (m *MsgCreateStableswapPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStableswapPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStableswapPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStableswapPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStableswapPool.Merge(m, src)
}
func (m *MsgCreateStableswapPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStableswapPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStableswapPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStableswapPool proto.InternalMessageInfo

func (m *MsgCreateStableswapPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgCreateStableswapPool) GetPoolParams() *PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return nil
}

func (m *MsgCreateStableswapPool) GetInitialPoolLiquidity() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialPoolLiquidity
	}
	return nil
}

func (m *MsgCreateStableswapPool) GetScalingFactors() []uint64 {
	if m != nil {
		return m.ScalingFactors
	}
	return nil
}

func (m *MsgCreateStableswapPool) GetFuturePoolGovernor() string {
	if m != nil {
		return m.FuturePoolGovernor
	}
	return ""
}

func (m *MsgCreateStableswapPool) GetScalingFactorController() string {
	if m != nil {
		return m.ScalingFactorController
	}
	return ""
}

// Returns a poolID with custom poolName.
type MsgCreateStableswapPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateStableswapPoolResponse) Reset()         { *m = MsgCreateStableswapPoolResponse{} }
func (m *MsgCreateStableswapPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateStableswapPoolResponse) ProtoMessage()    {}
func (*MsgCreateStableswapPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{1}
}
func (m *MsgCreateStableswapPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateStableswapPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateStableswapPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateStableswapPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateStableswapPoolResponse.Merge(m, src)
}
func (m *MsgCreateStableswapPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateStableswapPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateStableswapPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateStableswapPoolResponse proto.InternalMessageInfo

func (m *MsgCreateStableswapPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

// Sender must be the pool's scaling_factor_governor in order for the tx to
// succeed. Adjusts stableswap scaling factors.
type MsgStableSwapAdjustScalingFactors struct {
	Sender         string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolID         uint64   `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	ScalingFactors []uint64 `protobuf:"varint,3,rep,packed,name=scaling_factors,json=scalingFactors,proto3" json:"scaling_factors,omitempty" yaml:"stableswap_scaling_factor"`
}

func (m *MsgStableSwapAdjustScalingFactors) Reset()         { *m = MsgStableSwapAdjustScalingFactors{} }
func (m *MsgStableSwapAdjustScalingFactors) String() string { return proto.CompactTextString(m) }
func (*MsgStableSwapAdjustScalingFactors) ProtoMessage()    {}
func (*MsgStableSwapAdjustScalingFactors) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{2}
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStableSwapAdjustScalingFactors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactors.Merge(m, src)
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_Size() int {
	return m.Size()
}
func (m *MsgStableSwapAdjustScalingFactors) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactors.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStableSwapAdjustScalingFactors proto.InternalMessageInfo

func (m *MsgStableSwapAdjustScalingFactors) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgStableSwapAdjustScalingFactors) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

func (m *MsgStableSwapAdjustScalingFactors) GetScalingFactors() []uint64 {
	if m != nil {
		return m.ScalingFactors
	}
	return nil
}

type MsgStableSwapAdjustScalingFactorsResponse struct {
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) Reset() {
	*m = MsgStableSwapAdjustScalingFactorsResponse{}
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) String() string {
	return proto.CompactTextString(m)
}
func (*MsgStableSwapAdjustScalingFactorsResponse) ProtoMessage() {}
func (*MsgStableSwapAdjustScalingFactorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b7c8a0f24de97c, []int{3}
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse.Merge(m, src)
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStableSwapAdjustScalingFactorsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateStableswapPool)(nil), "mokita.gamm.poolmodels.stableswap.v1beta1.MsgCreateStableswapPool")
	proto.RegisterType((*MsgCreateStableswapPoolResponse)(nil), "mokita.gamm.poolmodels.stableswap.v1beta1.MsgCreateStableswapPoolResponse")
	proto.RegisterType((*MsgStableSwapAdjustScalingFactors)(nil), "mokita.gamm.poolmodels.stableswap.v1beta1.MsgStableSwapAdjustScalingFactors")
	proto.RegisterType((*MsgStableSwapAdjustScalingFactorsResponse)(nil), "mokita.gamm.poolmodels.stableswap.v1beta1.MsgStableSwapAdjustScalingFactorsResponse")
}

func init() {
	proto.RegisterFile("mokita/gamm/pool-models/stableswap/tx.proto", fileDescriptor_46b7c8a0f24de97c)
}

var fileDescriptor_46b7c8a0f24de97c = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x6f, 0xd6, 0x52, 0x84, 0x27, 0x40, 0x44, 0xd5, 0x96, 0x15, 0x29, 0x29, 0x81, 0x43, 0x07,
	0x2c, 0x66, 0x9b, 0x84, 0x04, 0xb7, 0xa5, 0x68, 0x68, 0x82, 0x4a, 0x23, 0x13, 0x17, 0x38, 0x14,
	0x27, 0xf1, 0x82, 0x21, 0x89, 0x43, 0xec, 0x76, 0xeb, 0x91, 0x6f, 0xc0, 0xc7, 0x40, 0x7c, 0x07,
	0x2e, 0x1c, 0xd0, 0x8e, 0x3b, 0x72, 0x0a, 0xa8, 0xfd, 0x06, 0xbd, 0x71, 0x43, 0x89, 0xd3, 0x7f,
	0x52, 0xbb, 0xad, 0x68, 0xa7, 0x38, 0xcf, 0xbf, 0xf7, 0xfb, 0xbd, 0xf7, 0x7b, 0xb6, 0xc1, 0x43,
	0xca, 0x02, 0xca, 0x08, 0x83, 0x1e, 0x0a, 0x02, 0x18, 0x51, 0xea, 0x6f, 0x04, 0xd4, 0xc5, 0x3e,
	0x83, 0x8c, 0x23, 0xdb, 0xc7, 0xec, 0x08, 0x45, 0x90, 0x1f, 0x1b, 0x51, 0x4c, 0x39, 0x95, 0xef,
	0xe7, 0x68, 0x23, 0x45, 0x1b, 0x29, 0x5a, 0x80, 0x8d, 0x31, 0xd8, 0xe8, 0x6c, 0xda, 0x98, 0xa3,
	0xcd, 0xaa, 0xea, 0x64, 0x60, 0x68, 0x23, 0x86, 0x61, 0x1e, 0x84, 0x0e, 0x25, 0xa1, 0xe0, 0xaa,
	0x56, 0x3c, 0xea, 0xd1, 0x6c, 0x09, 0xd3, 0x55, 0x1e, 0x7d, 0x72, 0x91, 0x7a, 0xc6, 0xcb, 0x56,
	0x8a, 0x10, 0xa9, 0xfa, 0xf7, 0x12, 0x58, 0x6d, 0x32, 0xaf, 0x11, 0x63, 0xc4, 0xf1, 0xc1, 0x08,
	0xb2, 0x4f, 0xa9, 0x2f, 0xaf, 0x83, 0x32, 0xc3, 0xa1, 0x8b, 0x63, 0x45, 0xaa, 0x49, 0xf5, 0x6b,
	0xe6, 0xad, 0x41, 0xa2, 0x5d, 0xef, 0xa2, 0xc0, 0x7f, 0xaa, 0x8b, 0xb8, 0x6e, 0xe5, 0x00, 0x99,
	0x82, 0xe5, 0x94, 0xb4, 0x15, 0xa1, 0x18, 0x05, 0x4c, 0x59, 0xaa, 0x49, 0xf5, 0xe5, 0xad, 0xc7,
	0xc6, 0xc5, 0x3b, 0x37, 0x52, 0xc5, 0xfd, 0x2c, 0xdb, 0x5c, 0x19, 0x24, 0x9a, 0x2c, 0x74, 0x26,
	0x48, 0x75, 0x0b, 0x44, 0x23, 0x8c, 0xfc, 0x59, 0x02, 0x2b, 0x24, 0x24, 0x9c, 0x20, 0x3f, 0x6b,
	0xa7, 0xe5, 0x93, 0x4f, 0x6d, 0xe2, 0x12, 0xde, 0x55, 0x8a, 0xb5, 0x62, 0x7d, 0x79, 0x6b, 0xcd,
	0x10, 0x56, 0x1a, 0xa9, 0x95, 0x23, 0x95, 0x06, 0x25, 0xa1, 0xf9, 0xe8, 0x24, 0xd1, 0x0a, 0xdf,
	0x7e, 0x6b, 0x75, 0x8f, 0xf0, 0xf7, 0x6d, 0xdb, 0x70, 0x68, 0x00, 0x73, 0xdf, 0xc5, 0x67, 0x83,
	0xb9, 0x1f, 0x21, 0xef, 0x46, 0x98, 0x65, 0x09, 0xcc, 0xaa, 0xe4, 0x52, 0x69, 0x91, 0x2f, 0x87,
	0x42, 0x72, 0x13, 0xdc, 0x64, 0x0e, 0xf2, 0x49, 0xe8, 0xb5, 0x0e, 0x91, 0xc3, 0x69, 0xcc, 0x94,
	0x52, 0xad, 0x58, 0x2f, 0x99, 0xf7, 0x06, 0x89, 0x56, 0xcb, 0x8d, 0x1a, 0xbb, 0x3e, 0x8d, 0xd5,
	0xad, 0x1b, 0x79, 0x60, 0x57, 0xe4, 0xca, 0xaf, 0x40, 0xe5, 0xb0, 0xcd, 0xdb, 0x31, 0x16, 0x0d,
	0x79, 0xb4, 0x83, 0xe3, 0x90, 0xc6, 0xca, 0x95, 0xcc, 0x7c, 0x6d, 0x90, 0x68, 0xb7, 0x05, 0xe7,
	0x2c, 0x94, 0x6e, 0xc9, 0x22, 0x9c, 0x96, 0xf8, 0x3c, 0x0f, 0xca, 0xef, 0xc0, 0xda, 0xb4, 0x6a,
	0xcb, 0xa1, 0x21, 0x8f, 0xa9, 0xef, 0xe3, 0x58, 0x29, 0x67, 0xbc, 0x93, 0xb5, 0xce, 0x83, 0xea,
	0xd6, 0xea, 0x54, 0xad, 0x8d, 0xf1, 0xce, 0x2e, 0xd0, 0xe6, 0x1c, 0x1f, 0x0b, 0xb3, 0x88, 0x86,
	0x0c, 0xcb, 0x77, 0xc1, 0xd5, 0xac, 0x54, 0xe2, 0x66, 0xe7, 0xa8, 0x64, 0x82, 0x5e, 0xa2, 0x95,
	0x53, 0xc8, 0xde, 0x33, 0xab, 0x9c, 0x6e, 0xed, 0xb9, 0xfa, 0x0f, 0x09, 0xdc, 0x69, 0x32, 0x4f,
	0x50, 0x1c, 0x1c, 0xa1, 0x68, 0xc7, 0xfd, 0xd0, 0x66, 0xfc, 0x60, 0xda, 0xa2, 0x05, 0x4e, 0xe4,
	0x84, 0xea, 0xd2, 0x3c, 0xd5, 0x59, 0x13, 0x2c, 0xfe, 0xff, 0x04, 0xf5, 0x07, 0x60, 0xfd, 0xdc,
	0x1e, 0x86, 0xb6, 0x6c, 0xfd, 0x5d, 0x02, 0xc5, 0x26, 0xf3, 0xe4, 0xaf, 0x12, 0xa8, 0xcc, 0xbc,
	0x7e, 0x8d, 0x45, 0xae, 0xcf, 0x9c, 0x21, 0x54, 0x5f, 0x5c, 0x02, 0xc9, 0x68, 0x92, 0x3f, 0x25,
	0xa0, 0x9e, 0x33, 0xa1, 0xe6, 0x82, 0x7a, 0x67, 0xd3, 0x55, 0x5f, 0x5f, 0x2a, 0xdd, 0xb0, 0x11,
	0xf3, 0xed, 0x49, 0x4f, 0x95, 0x4e, 0x7b, 0xaa, 0xf4, 0xa7, 0xa7, 0x4a, 0x5f, 0xfa, 0x6a, 0xe1,
	0xb4, 0xaf, 0x16, 0x7e, 0xf5, 0xd5, 0xc2, 0x9b, 0x9d, 0x89, 0x37, 0x21, 0x97, 0xde, 0xf0, 0x91,
	0xcd, 0x86, 0x3f, 0xb0, 0xb3, 0xb9, 0x0d, 0x8f, 0xcf, 0x7a, 0x68, 0xed, 0x72, 0xf6, 0xb2, 0x6e,
	0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc6, 0x5b, 0x3f, 0xf5, 0x26, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateStableswapPool(ctx context.Context, in *MsgCreateStableswapPool, opts ...grpc.CallOption) (*MsgCreateStableswapPoolResponse, error)
	StableSwapAdjustScalingFactors(ctx context.Context, in *MsgStableSwapAdjustScalingFactors, opts ...grpc.CallOption) (*MsgStableSwapAdjustScalingFactorsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateStableswapPool(ctx context.Context, in *MsgCreateStableswapPool, opts ...grpc.CallOption) (*MsgCreateStableswapPoolResponse, error) {
	out := new(MsgCreateStableswapPoolResponse)
	err := c.cc.Invoke(ctx, "/mokita.gamm.poolmodels.stableswap.v1beta1.Msg/CreateStableswapPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StableSwapAdjustScalingFactors(ctx context.Context, in *MsgStableSwapAdjustScalingFactors, opts ...grpc.CallOption) (*MsgStableSwapAdjustScalingFactorsResponse, error) {
	out := new(MsgStableSwapAdjustScalingFactorsResponse)
	err := c.cc.Invoke(ctx, "/mokita.gamm.poolmodels.stableswap.v1beta1.Msg/StableSwapAdjustScalingFactors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateStableswapPool(context.Context, *MsgCreateStableswapPool) (*MsgCreateStableswapPoolResponse, error)
	StableSwapAdjustScalingFactors(context.Context, *MsgStableSwapAdjustScalingFactors) (*MsgStableSwapAdjustScalingFactorsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateStableswapPool(ctx context.Context, req *MsgCreateStableswapPool) (*MsgCreateStableswapPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStableswapPool not implemented")
}
func (*UnimplementedMsgServer) StableSwapAdjustScalingFactors(ctx context.Context, req *MsgStableSwapAdjustScalingFactors) (*MsgStableSwapAdjustScalingFactorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StableSwapAdjustScalingFactors not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateStableswapPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateStableswapPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateStableswapPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mokita.gamm.poolmodels.stableswap.v1beta1.Msg/CreateStableswapPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateStableswapPool(ctx, req.(*MsgCreateStableswapPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StableSwapAdjustScalingFactors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStableSwapAdjustScalingFactors)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StableSwapAdjustScalingFactors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mokita.gamm.poolmodels.stableswap.v1beta1.Msg/StableSwapAdjustScalingFactors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StableSwapAdjustScalingFactors(ctx, req.(*MsgStableSwapAdjustScalingFactors))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mokita.gamm.poolmodels.stableswap.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStableswapPool",
			Handler:    _Msg_CreateStableswapPool_Handler,
		},
		{
			MethodName: "StableSwapAdjustScalingFactors",
			Handler:    _Msg_StableSwapAdjustScalingFactors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mokita/gamm/pool-models/stableswap/tx.proto",
}

func (m *MsgCreateStableswapPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStableswapPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStableswapPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ScalingFactorController) > 0 {
		i -= len(m.ScalingFactorController)
		copy(dAtA[i:], m.ScalingFactorController)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ScalingFactorController)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.FuturePoolGovernor) > 0 {
		i -= len(m.FuturePoolGovernor)
		copy(dAtA[i:], m.FuturePoolGovernor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FuturePoolGovernor)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ScalingFactors) > 0 {
		dAtA2 := make([]byte, len(m.ScalingFactors)*10)
		var j1 int
		for _, num := range m.ScalingFactors {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTx(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.InitialPoolLiquidity) > 0 {
		for iNdEx := len(m.InitialPoolLiquidity) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialPoolLiquidity[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateStableswapPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateStableswapPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateStableswapPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgStableSwapAdjustScalingFactors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStableSwapAdjustScalingFactors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStableSwapAdjustScalingFactors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ScalingFactors) > 0 {
		dAtA5 := make([]byte, len(m.ScalingFactors)*10)
		var j4 int
		for _, num := range m.ScalingFactors {
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		i -= j4
		copy(dAtA[i:], dAtA5[:j4])
		i = encodeVarintTx(dAtA, i, uint64(j4))
		i--
		dAtA[i] = 0x1a
	}
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateStableswapPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.InitialPoolLiquidity) > 0 {
		for _, e := range m.InitialPoolLiquidity {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.ScalingFactors) > 0 {
		l = 0
		for _, e := range m.ScalingFactors {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	l = len(m.FuturePoolGovernor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ScalingFactorController)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateStableswapPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func (m *MsgStableSwapAdjustScalingFactors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	if len(m.ScalingFactors) > 0 {
		l = 0
		for _, e := range m.ScalingFactors {
			l += sovTx(uint64(e))
		}
		n += 1 + sovTx(uint64(l)) + l
	}
	return n
}

func (m *MsgStableSwapAdjustScalingFactorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateStableswapPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateStableswapPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStableswapPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolParams == nil {
				m.PoolParams = &PoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPoolLiquidity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialPoolLiquidity = append(m.InitialPoolLiquidity, types.Coin{})
			if err := m.InitialPoolLiquidity[len(m.InitialPoolLiquidity)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ScalingFactors = append(m.ScalingFactors, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ScalingFactors) == 0 {
					m.ScalingFactors = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ScalingFactors = append(m.ScalingFactors, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ScalingFactors", wireType)
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGovernor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FuturePoolGovernor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScalingFactorController", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ScalingFactorController = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateStableswapPoolResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateStableswapPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateStableswapPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgStableSwapAdjustScalingFactors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ScalingFactors = append(m.ScalingFactors, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTx
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTx
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTx
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.ScalingFactors) == 0 {
					m.ScalingFactors = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTx
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ScalingFactors = append(m.ScalingFactors, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ScalingFactors", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgStableSwapAdjustScalingFactorsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStableSwapAdjustScalingFactorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)