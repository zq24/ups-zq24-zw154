// Code generated by protoc-gen-go. DO NOT EDIT.
// source: amazon/amazon_ups.proto

package pb_ua

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AWarehouse struct {
	Id                   *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	X                    *int32   `protobuf:"varint,2,req,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,3,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AWarehouse) Reset()         { *m = AWarehouse{} }
func (m *AWarehouse) String() string { return proto.CompactTextString(m) }
func (*AWarehouse) ProtoMessage()    {}
func (*AWarehouse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{0}
}

func (m *AWarehouse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AWarehouse.Unmarshal(m, b)
}
func (m *AWarehouse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AWarehouse.Marshal(b, m, deterministic)
}
func (m *AWarehouse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AWarehouse.Merge(m, src)
}
func (m *AWarehouse) XXX_Size() int {
	return xxx_messageInfo_AWarehouse.Size(m)
}
func (m *AWarehouse) XXX_DiscardUnknown() {
	xxx_messageInfo_AWarehouse.DiscardUnknown(m)
}

var xxx_messageInfo_AWarehouse proto.InternalMessageInfo

func (m *AWarehouse) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *AWarehouse) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *AWarehouse) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type ASendTruck struct {
	Whs                  *AWarehouse     `protobuf:"bytes,1,req,name=whs" json:"whs,omitempty"`
	Packages             []*APackageInfo `protobuf:"bytes,2,rep,name=packages" json:"packages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ASendTruck) Reset()         { *m = ASendTruck{} }
func (m *ASendTruck) String() string { return proto.CompactTextString(m) }
func (*ASendTruck) ProtoMessage()    {}
func (*ASendTruck) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{1}
}

func (m *ASendTruck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ASendTruck.Unmarshal(m, b)
}
func (m *ASendTruck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ASendTruck.Marshal(b, m, deterministic)
}
func (m *ASendTruck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ASendTruck.Merge(m, src)
}
func (m *ASendTruck) XXX_Size() int {
	return xxx_messageInfo_ASendTruck.Size(m)
}
func (m *ASendTruck) XXX_DiscardUnknown() {
	xxx_messageInfo_ASendTruck.DiscardUnknown(m)
}

var xxx_messageInfo_ASendTruck proto.InternalMessageInfo

func (m *ASendTruck) GetWhs() *AWarehouse {
	if m != nil {
		return m.Whs
	}
	return nil
}

func (m *ASendTruck) GetPackages() []*APackageInfo {
	if m != nil {
		return m.Packages
	}
	return nil
}

type USendWorldID struct {
	Worldid              *int64   `protobuf:"varint,1,req,name=worldid" json:"worldid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *USendWorldID) Reset()         { *m = USendWorldID{} }
func (m *USendWorldID) String() string { return proto.CompactTextString(m) }
func (*USendWorldID) ProtoMessage()    {}
func (*USendWorldID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{2}
}

func (m *USendWorldID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_USendWorldID.Unmarshal(m, b)
}
func (m *USendWorldID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_USendWorldID.Marshal(b, m, deterministic)
}
func (m *USendWorldID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_USendWorldID.Merge(m, src)
}
func (m *USendWorldID) XXX_Size() int {
	return xxx_messageInfo_USendWorldID.Size(m)
}
func (m *USendWorldID) XXX_DiscardUnknown() {
	xxx_messageInfo_USendWorldID.DiscardUnknown(m)
}

var xxx_messageInfo_USendWorldID proto.InternalMessageInfo

func (m *USendWorldID) GetWorldid() int64 {
	if m != nil && m.Worldid != nil {
		return *m.Worldid
	}
	return 0
}

type UTruckSent struct {
	Truckid              *int32          `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	Packages             []*APackageInfo `protobuf:"bytes,2,rep,name=packages" json:"packages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UTruckSent) Reset()         { *m = UTruckSent{} }
func (m *UTruckSent) String() string { return proto.CompactTextString(m) }
func (*UTruckSent) ProtoMessage()    {}
func (*UTruckSent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{3}
}

func (m *UTruckSent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTruckSent.Unmarshal(m, b)
}
func (m *UTruckSent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTruckSent.Marshal(b, m, deterministic)
}
func (m *UTruckSent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTruckSent.Merge(m, src)
}
func (m *UTruckSent) XXX_Size() int {
	return xxx_messageInfo_UTruckSent.Size(m)
}
func (m *UTruckSent) XXX_DiscardUnknown() {
	xxx_messageInfo_UTruckSent.DiscardUnknown(m)
}

var xxx_messageInfo_UTruckSent proto.InternalMessageInfo

func (m *UTruckSent) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *UTruckSent) GetPackages() []*APackageInfo {
	if m != nil {
		return m.Packages
	}
	return nil
}

type UTruckArrived struct {
	Truckid              *int32   `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UTruckArrived) Reset()         { *m = UTruckArrived{} }
func (m *UTruckArrived) String() string { return proto.CompactTextString(m) }
func (*UTruckArrived) ProtoMessage()    {}
func (*UTruckArrived) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{4}
}

func (m *UTruckArrived) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTruckArrived.Unmarshal(m, b)
}
func (m *UTruckArrived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTruckArrived.Marshal(b, m, deterministic)
}
func (m *UTruckArrived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTruckArrived.Merge(m, src)
}
func (m *UTruckArrived) XXX_Size() int {
	return xxx_messageInfo_UTruckArrived.Size(m)
}
func (m *UTruckArrived) XXX_DiscardUnknown() {
	xxx_messageInfo_UTruckArrived.DiscardUnknown(m)
}

var xxx_messageInfo_UTruckArrived proto.InternalMessageInfo

func (m *UTruckArrived) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

type APackageInfo struct {
	Packageid            *int64   `protobuf:"varint,1,req,name=packageid" json:"packageid,omitempty"`
	X                    *int32   `protobuf:"varint,2,req,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,3,req,name=y" json:"y,omitempty"`
	UpsUserName          *string  `protobuf:"bytes,4,opt,name=ups_user_name,json=upsUserName" json:"ups_user_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *APackageInfo) Reset()         { *m = APackageInfo{} }
func (m *APackageInfo) String() string { return proto.CompactTextString(m) }
func (*APackageInfo) ProtoMessage()    {}
func (*APackageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{5}
}

func (m *APackageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_APackageInfo.Unmarshal(m, b)
}
func (m *APackageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_APackageInfo.Marshal(b, m, deterministic)
}
func (m *APackageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_APackageInfo.Merge(m, src)
}
func (m *APackageInfo) XXX_Size() int {
	return xxx_messageInfo_APackageInfo.Size(m)
}
func (m *APackageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_APackageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_APackageInfo proto.InternalMessageInfo

func (m *APackageInfo) GetPackageid() int64 {
	if m != nil && m.Packageid != nil {
		return *m.Packageid
	}
	return 0
}

func (m *APackageInfo) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *APackageInfo) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *APackageInfo) GetUpsUserName() string {
	if m != nil && m.UpsUserName != nil {
		return *m.UpsUserName
	}
	return ""
}

type AStartDelivery struct {
	Truckid              *int32          `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	Packages             []*APackageInfo `protobuf:"bytes,2,rep,name=packages" json:"packages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AStartDelivery) Reset()         { *m = AStartDelivery{} }
func (m *AStartDelivery) String() string { return proto.CompactTextString(m) }
func (*AStartDelivery) ProtoMessage()    {}
func (*AStartDelivery) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{6}
}

func (m *AStartDelivery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AStartDelivery.Unmarshal(m, b)
}
func (m *AStartDelivery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AStartDelivery.Marshal(b, m, deterministic)
}
func (m *AStartDelivery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AStartDelivery.Merge(m, src)
}
func (m *AStartDelivery) XXX_Size() int {
	return xxx_messageInfo_AStartDelivery.Size(m)
}
func (m *AStartDelivery) XXX_DiscardUnknown() {
	xxx_messageInfo_AStartDelivery.DiscardUnknown(m)
}

var xxx_messageInfo_AStartDelivery proto.InternalMessageInfo

func (m *AStartDelivery) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *AStartDelivery) GetPackages() []*APackageInfo {
	if m != nil {
		return m.Packages
	}
	return nil
}

type UDelivered struct {
	Shipid               *int64   `protobuf:"varint,1,req,name=shipid" json:"shipid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UDelivered) Reset()         { *m = UDelivered{} }
func (m *UDelivered) String() string { return proto.CompactTextString(m) }
func (*UDelivered) ProtoMessage()    {}
func (*UDelivered) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{7}
}

func (m *UDelivered) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UDelivered.Unmarshal(m, b)
}
func (m *UDelivered) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UDelivered.Marshal(b, m, deterministic)
}
func (m *UDelivered) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UDelivered.Merge(m, src)
}
func (m *UDelivered) XXX_Size() int {
	return xxx_messageInfo_UDelivered.Size(m)
}
func (m *UDelivered) XXX_DiscardUnknown() {
	xxx_messageInfo_UDelivered.DiscardUnknown(m)
}

var xxx_messageInfo_UDelivered proto.InternalMessageInfo

func (m *UDelivered) GetShipid() int64 {
	if m != nil && m.Shipid != nil {
		return *m.Shipid
	}
	return 0
}

type AtoUCommands struct {
	Sendtrucks           []*ASendTruck     `protobuf:"bytes,1,rep,name=sendtrucks" json:"sendtrucks,omitempty"`
	Startdelivery        []*AStartDelivery `protobuf:"bytes,2,rep,name=startdelivery" json:"startdelivery,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AtoUCommands) Reset()         { *m = AtoUCommands{} }
func (m *AtoUCommands) String() string { return proto.CompactTextString(m) }
func (*AtoUCommands) ProtoMessage()    {}
func (*AtoUCommands) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{8}
}

func (m *AtoUCommands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AtoUCommands.Unmarshal(m, b)
}
func (m *AtoUCommands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AtoUCommands.Marshal(b, m, deterministic)
}
func (m *AtoUCommands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AtoUCommands.Merge(m, src)
}
func (m *AtoUCommands) XXX_Size() int {
	return xxx_messageInfo_AtoUCommands.Size(m)
}
func (m *AtoUCommands) XXX_DiscardUnknown() {
	xxx_messageInfo_AtoUCommands.DiscardUnknown(m)
}

var xxx_messageInfo_AtoUCommands proto.InternalMessageInfo

func (m *AtoUCommands) GetSendtrucks() []*ASendTruck {
	if m != nil {
		return m.Sendtrucks
	}
	return nil
}

func (m *AtoUCommands) GetStartdelivery() []*AStartDelivery {
	if m != nil {
		return m.Startdelivery
	}
	return nil
}

type UtoAResponses struct {
	Trucksent            []*UTruckSent    `protobuf:"bytes,1,rep,name=trucksent" json:"trucksent,omitempty"`
	Arrived              []*UTruckArrived `protobuf:"bytes,2,rep,name=arrived" json:"arrived,omitempty"`
	Delivered            []*UDelivered    `protobuf:"bytes,3,rep,name=delivered" json:"delivered,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UtoAResponses) Reset()         { *m = UtoAResponses{} }
func (m *UtoAResponses) String() string { return proto.CompactTextString(m) }
func (*UtoAResponses) ProtoMessage()    {}
func (*UtoAResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{9}
}

func (m *UtoAResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UtoAResponses.Unmarshal(m, b)
}
func (m *UtoAResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UtoAResponses.Marshal(b, m, deterministic)
}
func (m *UtoAResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UtoAResponses.Merge(m, src)
}
func (m *UtoAResponses) XXX_Size() int {
	return xxx_messageInfo_UtoAResponses.Size(m)
}
func (m *UtoAResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_UtoAResponses.DiscardUnknown(m)
}

var xxx_messageInfo_UtoAResponses proto.InternalMessageInfo

func (m *UtoAResponses) GetTrucksent() []*UTruckSent {
	if m != nil {
		return m.Trucksent
	}
	return nil
}

func (m *UtoAResponses) GetArrived() []*UTruckArrived {
	if m != nil {
		return m.Arrived
	}
	return nil
}

func (m *UtoAResponses) GetDelivered() []*UDelivered {
	if m != nil {
		return m.Delivered
	}
	return nil
}

// Probable Message
type ATraceShip struct {
	Shipid               *int64   `protobuf:"varint,1,req,name=shipid" json:"shipid,omitempty"`
	Seqnum               *int64   `protobuf:"varint,2,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ATraceShip) Reset()         { *m = ATraceShip{} }
func (m *ATraceShip) String() string { return proto.CompactTextString(m) }
func (*ATraceShip) ProtoMessage()    {}
func (*ATraceShip) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{10}
}

func (m *ATraceShip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ATraceShip.Unmarshal(m, b)
}
func (m *ATraceShip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ATraceShip.Marshal(b, m, deterministic)
}
func (m *ATraceShip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ATraceShip.Merge(m, src)
}
func (m *ATraceShip) XXX_Size() int {
	return xxx_messageInfo_ATraceShip.Size(m)
}
func (m *ATraceShip) XXX_DiscardUnknown() {
	xxx_messageInfo_ATraceShip.DiscardUnknown(m)
}

var xxx_messageInfo_ATraceShip proto.InternalMessageInfo

func (m *ATraceShip) GetShipid() int64 {
	if m != nil && m.Shipid != nil {
		return *m.Shipid
	}
	return 0
}

func (m *ATraceShip) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UShipInfo struct {
	Shipid               *int64   `protobuf:"varint,1,req,name=shipid" json:"shipid,omitempty"`
	Info                 *string  `protobuf:"bytes,2,req,name=info" json:"info,omitempty"`
	Seqnum               *int64   `protobuf:"varint,3,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UShipInfo) Reset()         { *m = UShipInfo{} }
func (m *UShipInfo) String() string { return proto.CompactTextString(m) }
func (*UShipInfo) ProtoMessage()    {}
func (*UShipInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d02007b38b3077a6, []int{11}
}

func (m *UShipInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UShipInfo.Unmarshal(m, b)
}
func (m *UShipInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UShipInfo.Marshal(b, m, deterministic)
}
func (m *UShipInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UShipInfo.Merge(m, src)
}
func (m *UShipInfo) XXX_Size() int {
	return xxx_messageInfo_UShipInfo.Size(m)
}
func (m *UShipInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UShipInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UShipInfo proto.InternalMessageInfo

func (m *UShipInfo) GetShipid() int64 {
	if m != nil && m.Shipid != nil {
		return *m.Shipid
	}
	return 0
}

func (m *UShipInfo) GetInfo() string {
	if m != nil && m.Info != nil {
		return *m.Info
	}
	return ""
}

func (m *UShipInfo) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

func init() {
	proto.RegisterType((*AWarehouse)(nil), "AWarehouse")
	proto.RegisterType((*ASendTruck)(nil), "ASendTruck")
	proto.RegisterType((*USendWorldID)(nil), "USendWorldID")
	proto.RegisterType((*UTruckSent)(nil), "UTruckSent")
	proto.RegisterType((*UTruckArrived)(nil), "UTruckArrived")
	proto.RegisterType((*APackageInfo)(nil), "APackageInfo")
	proto.RegisterType((*AStartDelivery)(nil), "AStartDelivery")
	proto.RegisterType((*UDelivered)(nil), "UDelivered")
	proto.RegisterType((*AtoUCommands)(nil), "AtoUCommands")
	proto.RegisterType((*UtoAResponses)(nil), "UtoAResponses")
	proto.RegisterType((*ATraceShip)(nil), "ATraceShip")
	proto.RegisterType((*UShipInfo)(nil), "UShipInfo")
}

func init() { proto.RegisterFile("amazon/amazon_ups.proto", fileDescriptor_d02007b38b3077a6) }

var fileDescriptor_d02007b38b3077a6 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4b, 0x8b, 0xd4, 0x40,
	0x10, 0xc7, 0x49, 0xb2, 0xba, 0xa6, 0x66, 0x32, 0x42, 0x1f, 0x34, 0x07, 0x85, 0xa1, 0xf1, 0x90,
	0x45, 0x88, 0xb0, 0x20, 0x78, 0xf0, 0x12, 0xdc, 0xcb, 0x5e, 0x7c, 0xf4, 0x6c, 0xdc, 0xe3, 0xd0,
	0x4c, 0xd7, 0x3a, 0x61, 0x27, 0xdd, 0xb1, 0x3b, 0xd9, 0x9d, 0xf1, 0x4b, 0xf8, 0x95, 0xa5, 0x3b,
	0xcf, 0x39, 0x8c, 0x20, 0x78, 0x4a, 0xfe, 0x55, 0x5d, 0xbf, 0x7a, 0xc2, 0x4b, 0x5e, 0xf2, 0x5f,
	0x4a, 0xbe, 0x6b, 0x3f, 0xeb, 0xa6, 0x32, 0x69, 0xa5, 0x55, 0xad, 0xe8, 0x07, 0x80, 0xec, 0x96,
	0x6b, 0xdc, 0xaa, 0xc6, 0x20, 0x59, 0x80, 0x5f, 0x88, 0xd8, 0x5b, 0xfa, 0xc9, 0x13, 0xe6, 0x17,
	0x82, 0xcc, 0xc1, 0xdb, 0xc7, 0xbe, 0x93, 0xde, 0xde, 0xaa, 0x43, 0x1c, 0xb4, 0xea, 0x40, 0xbf,
	0x03, 0x64, 0x2b, 0x94, 0xe2, 0x46, 0x37, 0x9b, 0x7b, 0xf2, 0x1a, 0x82, 0xc7, 0xad, 0x71, 0xa1,
	0xb3, 0xcb, 0x59, 0x3a, 0x32, 0x99, 0xb5, 0x93, 0x0b, 0x78, 0x56, 0xf1, 0xcd, 0x3d, 0xff, 0x81,
	0x26, 0xf6, 0x97, 0x41, 0x32, 0xbb, 0x8c, 0xd2, 0xec, 0x6b, 0x6b, 0xb9, 0x96, 0x77, 0x8a, 0x0d,
	0x6e, 0x9a, 0xc0, 0x3c, 0xb7, 0xdc, 0x5b, 0xa5, 0x77, 0xe2, 0xfa, 0x8a, 0xc4, 0x70, 0xfe, 0x68,
	0x7f, 0xbb, 0xc2, 0x02, 0xd6, 0x4b, 0xfa, 0x0d, 0x20, 0x77, 0xd9, 0x57, 0x28, 0x6b, 0xfb, 0xae,
	0xb6, 0x62, 0x68, 0xa0, 0x97, 0xff, 0x92, 0xfc, 0x02, 0xa2, 0x16, 0x99, 0x69, 0x5d, 0x3c, 0xa0,
	0x38, 0x4d, 0xa5, 0x3b, 0x98, 0x4f, 0x21, 0xe4, 0x15, 0x84, 0x1d, 0x66, 0xa8, 0x74, 0x34, 0xfc,
	0x6d, 0x92, 0x84, 0x42, 0xd4, 0x54, 0x66, 0xdd, 0x18, 0xd4, 0x6b, 0xc9, 0x4b, 0x8c, 0xcf, 0x96,
	0x5e, 0x12, 0xb2, 0x59, 0x53, 0x99, 0xdc, 0xa0, 0xfe, 0xcc, 0x4b, 0xa4, 0x39, 0x2c, 0xb2, 0x55,
	0xcd, 0x75, 0x7d, 0x85, 0xbb, 0xe2, 0x01, 0xf5, 0xe1, 0xff, 0xf4, 0xfb, 0x06, 0x20, 0xef, 0x88,
	0x28, 0xc8, 0x0b, 0x78, 0x6a, 0xb6, 0x45, 0x35, 0xd4, 0xdf, 0x29, 0xaa, 0x61, 0x9e, 0xd5, 0x2a,
	0xff, 0xa4, 0xca, 0x92, 0x4b, 0x61, 0xc8, 0x5b, 0x00, 0x83, 0x52, 0xb8, 0x7c, 0x76, 0xe7, 0x41,
	0xbb, 0xf3, 0xe1, 0x1a, 0xd8, 0xc4, 0x4d, 0xde, 0x43, 0x64, 0x6c, 0xe1, 0xa2, 0x2b, 0xbc, 0x2b,
	0xe9, 0x79, 0x7a, 0xdc, 0x0f, 0x3b, 0x7e, 0x45, 0x7f, 0x7b, 0x10, 0xe5, 0xb5, 0xca, 0x18, 0x9a,
	0x4a, 0x49, 0x83, 0xf6, 0x86, 0xc2, 0x16, 0x89, 0xb2, 0x1e, 0x92, 0x8e, 0x07, 0xc0, 0x46, 0x2f,
	0x49, 0xe0, 0x9c, 0xb7, 0x0b, 0xec, 0xb2, 0x2d, 0xd2, 0xa3, 0xb5, 0xb2, 0xde, 0x6d, 0xa1, 0xa2,
	0xef, 0x3f, 0x0e, 0x7a, 0xe8, 0x30, 0x12, 0x36, 0x7a, 0xe9, 0x47, 0x80, 0xec, 0x46, 0xf3, 0x0d,
	0xae, 0xb6, 0x45, 0x75, 0x6a, 0x56, 0xce, 0x8e, 0x3f, 0x65, 0x53, 0xba, 0x6d, 0x5b, 0xbb, 0x53,
	0xf4, 0x0b, 0x84, 0xb9, 0x0d, 0x74, 0xb7, 0x72, 0x2a, 0x98, 0xc0, 0x59, 0x21, 0xef, 0x94, 0x0b,
	0x0d, 0x99, 0xfb, 0x9f, 0x00, 0x83, 0x29, 0xf0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xe4,
	0x56, 0xe6, 0xd3, 0x03, 0x00, 0x00,
}
