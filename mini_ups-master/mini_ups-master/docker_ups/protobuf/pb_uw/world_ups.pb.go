// Code generated by protoc-gen-go. DO NOT EDIT.
// source: world_ups.proto

package pb_uw

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	"net"
	"time"
	psql "database_utils"
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
//-----------------------------------------------------------------------
type Request interface {
	WaitAck()
	Retransmit()
}


//----------------------------------------------------------------------

type UInitTruck struct {
	Id                   *int32   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	X                    *int32   `protobuf:"varint,2,req,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,3,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UInitTruck) Reset()         { *m = UInitTruck{} }
func (m *UInitTruck) String() string { return proto.CompactTextString(m) }
func (*UInitTruck) ProtoMessage()    {}
func (*UInitTruck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{0}
}

func (m *UInitTruck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UInitTruck.Unmarshal(m, b)
}
func (m *UInitTruck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UInitTruck.Marshal(b, m, deterministic)
}
func (m *UInitTruck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UInitTruck.Merge(m, src)
}
func (m *UInitTruck) XXX_Size() int {
	return xxx_messageInfo_UInitTruck.Size(m)
}
func (m *UInitTruck) XXX_DiscardUnknown() {
	xxx_messageInfo_UInitTruck.DiscardUnknown(m)
}

var xxx_messageInfo_UInitTruck proto.InternalMessageInfo

func (m *UInitTruck) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *UInitTruck) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *UInitTruck) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type UConnect struct {
	Worldid              *int64        `protobuf:"varint,1,opt,name=worldid" json:"worldid,omitempty"`
	Trucks               []*UInitTruck `protobuf:"bytes,2,rep,name=trucks" json:"trucks,omitempty"`
	IsAmazon             *bool         `protobuf:"varint,3,req,name=isAmazon" json:"isAmazon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UConnect) Reset()         { *m = UConnect{} }
func (m *UConnect) String() string { return proto.CompactTextString(m) }
func (*UConnect) ProtoMessage()    {}
func (*UConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{1}
}

func (m *UConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UConnect.Unmarshal(m, b)
}
func (m *UConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UConnect.Marshal(b, m, deterministic)
}
func (m *UConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UConnect.Merge(m, src)
}
func (m *UConnect) XXX_Size() int {
	return xxx_messageInfo_UConnect.Size(m)
}
func (m *UConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_UConnect.DiscardUnknown(m)
}

var xxx_messageInfo_UConnect proto.InternalMessageInfo

func (m *UConnect) GetWorldid() int64 {
	if m != nil && m.Worldid != nil {
		return *m.Worldid
	}
	return 0
}

func (m *UConnect) GetTrucks() []*UInitTruck {
	if m != nil {
		return m.Trucks
	}
	return nil
}

func (m *UConnect) GetIsAmazon() bool {
	if m != nil && m.IsAmazon != nil {
		return *m.IsAmazon
	}
	return false
}

type UConnected struct {
	Worldid              *int64   `protobuf:"varint,1,req,name=worldid" json:"worldid,omitempty"`
	Result               *string  `protobuf:"bytes,2,req,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UConnected) Reset()         { *m = UConnected{} }
func (m *UConnected) String() string { return proto.CompactTextString(m) }
func (*UConnected) ProtoMessage()    {}
func (*UConnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{2}
}

func (m *UConnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UConnected.Unmarshal(m, b)
}
func (m *UConnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UConnected.Marshal(b, m, deterministic)
}
func (m *UConnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UConnected.Merge(m, src)
}
func (m *UConnected) XXX_Size() int {
	return xxx_messageInfo_UConnected.Size(m)
}
func (m *UConnected) XXX_DiscardUnknown() {
	xxx_messageInfo_UConnected.DiscardUnknown(m)
}

var xxx_messageInfo_UConnected proto.InternalMessageInfo

func (m *UConnected) GetWorldid() int64 {
	if m != nil && m.Worldid != nil {
		return *m.Worldid
	}
	return 0
}

func (m *UConnected) GetResult() string {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return ""
}

type UGoPickup struct {
	Truckid              *int32   `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	Whid                 *int32   `protobuf:"varint,2,req,name=whid" json:"whid,omitempty"`
	Seqnum               *int64   `protobuf:"varint,3,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}



func (m *UGoPickup) WaitAck(conn *net.Conn)        {
	psql.Insert_request_gopickup(*m.Seqnum,*m.Truckid,*m.Whid)
	go func() {
		for{
			timer := time.NewTimer(15 * time.Minute)
			<-timer.C
			if _, err := psql.Query_request_gopickup(*m.Seqnum); err != nil {
				break
			} else {
				(*m).Retransmit(conn)
			}
		}
	} ()
}

func (m *UGoPickup) Retransmit(conn *net.Conn)  {

	out, err := proto.Marshal(m)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(m)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*conn).Write(out); err != nil {
		fmt.Println("Failf to retransmit UGoPickup")
	}
}

func (m *UGoPickup) Reset()         { *m = UGoPickup{} }
func (m *UGoPickup) String() string { return proto.CompactTextString(m) }
func (*UGoPickup) ProtoMessage()    {}
func (*UGoPickup) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{3}
}

func (m *UGoPickup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UGoPickup.Unmarshal(m, b)
}
func (m *UGoPickup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UGoPickup.Marshal(b, m, deterministic)
}
func (m *UGoPickup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UGoPickup.Merge(m, src)
}
func (m *UGoPickup) XXX_Size() int {
	return xxx_messageInfo_UGoPickup.Size(m)
}
func (m *UGoPickup) XXX_DiscardUnknown() {
	xxx_messageInfo_UGoPickup.DiscardUnknown(m)
}

var xxx_messageInfo_UGoPickup proto.InternalMessageInfo

func (m *UGoPickup) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *UGoPickup) GetWhid() int32 {
	if m != nil && m.Whid != nil {
		return *m.Whid
	}
	return 0
}

func (m *UGoPickup) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UFinished struct {
	Truckid              *int32   `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	X                    *int32   `protobuf:"varint,2,req,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,3,req,name=y" json:"y,omitempty"`
	Status               *string  `protobuf:"bytes,4,req,name=status" json:"status,omitempty"`
	Seqnum               *int64   `protobuf:"varint,5,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UFinished) Reset()         { *m = UFinished{} }
func (m *UFinished) String() string { return proto.CompactTextString(m) }
func (*UFinished) ProtoMessage()    {}
func (*UFinished) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{4}
}

func (m *UFinished) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UFinished.Unmarshal(m, b)
}
func (m *UFinished) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UFinished.Marshal(b, m, deterministic)
}
func (m *UFinished) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UFinished.Merge(m, src)
}
func (m *UFinished) XXX_Size() int {
	return xxx_messageInfo_UFinished.Size(m)
}
func (m *UFinished) XXX_DiscardUnknown() {
	xxx_messageInfo_UFinished.DiscardUnknown(m)
}

var xxx_messageInfo_UFinished proto.InternalMessageInfo

func (m *UFinished) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *UFinished) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *UFinished) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *UFinished) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *UFinished) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UDeliveryMade struct {
	Truckid              *int32   `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	Packageid            *int64   `protobuf:"varint,2,req,name=packageid" json:"packageid,omitempty"`
	Seqnum               *int64   `protobuf:"varint,3,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UDeliveryMade) Reset()         { *m = UDeliveryMade{} }
func (m *UDeliveryMade) String() string { return proto.CompactTextString(m) }
func (*UDeliveryMade) ProtoMessage()    {}
func (*UDeliveryMade) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{5}
}

func (m *UDeliveryMade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UDeliveryMade.Unmarshal(m, b)
}
func (m *UDeliveryMade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UDeliveryMade.Marshal(b, m, deterministic)
}
func (m *UDeliveryMade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UDeliveryMade.Merge(m, src)
}
func (m *UDeliveryMade) XXX_Size() int {
	return xxx_messageInfo_UDeliveryMade.Size(m)
}
func (m *UDeliveryMade) XXX_DiscardUnknown() {
	xxx_messageInfo_UDeliveryMade.DiscardUnknown(m)
}

var xxx_messageInfo_UDeliveryMade proto.InternalMessageInfo

func (m *UDeliveryMade) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *UDeliveryMade) GetPackageid() int64 {
	if m != nil && m.Packageid != nil {
		return *m.Packageid
	}
	return 0
}

func (m *UDeliveryMade) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UDeliveryLocation struct {
	Packageid            *int64   `protobuf:"varint,1,req,name=packageid" json:"packageid,omitempty"`
	X                    *int32   `protobuf:"varint,2,req,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,3,req,name=y" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UDeliveryLocation) Reset()         { *m = UDeliveryLocation{} }
func (m *UDeliveryLocation) String() string { return proto.CompactTextString(m) }
func (*UDeliveryLocation) ProtoMessage()    {}
func (*UDeliveryLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{6}
}

func (m *UDeliveryLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UDeliveryLocation.Unmarshal(m, b)
}
func (m *UDeliveryLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UDeliveryLocation.Marshal(b, m, deterministic)
}
func (m *UDeliveryLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UDeliveryLocation.Merge(m, src)
}
func (m *UDeliveryLocation) XXX_Size() int {
	return xxx_messageInfo_UDeliveryLocation.Size(m)
}
func (m *UDeliveryLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_UDeliveryLocation.DiscardUnknown(m)
}

var xxx_messageInfo_UDeliveryLocation proto.InternalMessageInfo

func (m *UDeliveryLocation) GetPackageid() int64 {
	if m != nil && m.Packageid != nil {
		return *m.Packageid
	}
	return 0
}

func (m *UDeliveryLocation) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *UDeliveryLocation) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type UGoDeliver struct {
	Truckid              *int32               `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	Packages             []*UDeliveryLocation `protobuf:"bytes,2,rep,name=packages" json:"packages,omitempty"`
	Seqnum               *int64               `protobuf:"varint,3,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}



func (m *UGoDeliver) WaitAck(conn *net.Conn)     {
	var pkgs []psql.UDeliveryLocation
	for _, pkg := range (*m).Packages {
		var obj psql.UDeliveryLocation
		obj.PackId = *pkg.Packageid
		obj.X = *pkg.X
		obj.Y = *pkg.Y
		pkgs = append(pkgs,obj)
	}
	psql.Insert_request_godeliver(*m.Seqnum,*m.Truckid, pkgs)
	go func() {
		for{
			timer := time.NewTimer(15 * time.Minute)
			<-timer.C
			if _, err := psql.Query_request_godeliver(*m.Seqnum); err != nil {
				break
			} else {
				(*m).Retransmit(conn)
			}
		}
	} ()
}

func (m *UGoDeliver) Retransmit(conn *net.Conn)  {

	out, err := proto.Marshal(m)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(m)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*conn).Write(out); err != nil {
		fmt.Println("fail to retransmit UGoDeliver")
	}
}


func (m *UGoDeliver) Reset()         { *m = UGoDeliver{} }
func (m *UGoDeliver) String() string { return proto.CompactTextString(m) }
func (*UGoDeliver) ProtoMessage()    {}
func (*UGoDeliver) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{7}
}

func (m *UGoDeliver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UGoDeliver.Unmarshal(m, b)
}
func (m *UGoDeliver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UGoDeliver.Marshal(b, m, deterministic)
}
func (m *UGoDeliver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UGoDeliver.Merge(m, src)
}
func (m *UGoDeliver) XXX_Size() int {
	return xxx_messageInfo_UGoDeliver.Size(m)
}
func (m *UGoDeliver) XXX_DiscardUnknown() {
	xxx_messageInfo_UGoDeliver.DiscardUnknown(m)
}

var xxx_messageInfo_UGoDeliver proto.InternalMessageInfo

func (m *UGoDeliver) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *UGoDeliver) GetPackages() []*UDeliveryLocation {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *UGoDeliver) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UErr struct {
	Err                  *string  `protobuf:"bytes,1,req,name=err" json:"err,omitempty"`
	Originseqnum         *int64   `protobuf:"varint,2,req,name=originseqnum" json:"originseqnum,omitempty"`
	Seqnum               *int64   `protobuf:"varint,3,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UErr) Reset()         { *m = UErr{} }
func (m *UErr) String() string { return proto.CompactTextString(m) }
func (*UErr) ProtoMessage()    {}
func (*UErr) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{8}
}

func (m *UErr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UErr.Unmarshal(m, b)
}
func (m *UErr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UErr.Marshal(b, m, deterministic)
}
func (m *UErr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UErr.Merge(m, src)
}
func (m *UErr) XXX_Size() int {
	return xxx_messageInfo_UErr.Size(m)
}
func (m *UErr) XXX_DiscardUnknown() {
	xxx_messageInfo_UErr.DiscardUnknown(m)
}

var xxx_messageInfo_UErr proto.InternalMessageInfo

func (m *UErr) GetErr() string {
	if m != nil && m.Err != nil {
		return *m.Err
	}
	return ""
}

func (m *UErr) GetOriginseqnum() int64 {
	if m != nil && m.Originseqnum != nil {
		return *m.Originseqnum
	}
	return 0
}

func (m *UErr) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UQuery struct {
	Truckid              *int32   `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	Seqnum               *int64   `protobuf:"varint,2,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}


func (m *UQuery) WaitAck(conn *net.Conn)       {
	psql.Insert_request_query(*m.Seqnum, *m.Truckid)
	go func() {
		for{
			timer := time.NewTimer(15 * time.Minute)
			<-timer.C
			if _, err := psql.Query_request_goquery(*m.Seqnum); err != nil {
				break
			} else {
				(*m).Retransmit(conn)
			}
		}
	} ()
}

func (m *UQuery) Retransmit(conn *net.Conn) {

	out, err := proto.Marshal(m)
	if err != nil {
		fmt.Println("Failed to encode UConnect")
	}
	n := proto.Size(m)
	out = append(proto.EncodeVarint(uint64(n)), out...)
	if _, err = (*conn).Write(out); err != nil {
		fmt.Printf("fail to retransmit UQuery")
	}
}

func (m *UQuery) Reset()         { *m = UQuery{} }
func (m *UQuery) String() string { return proto.CompactTextString(m) }
func (*UQuery) ProtoMessage()    {}
func (*UQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{9}
}

func (m *UQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UQuery.Unmarshal(m, b)
}
func (m *UQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UQuery.Marshal(b, m, deterministic)
}
func (m *UQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UQuery.Merge(m, src)
}
func (m *UQuery) XXX_Size() int {
	return xxx_messageInfo_UQuery.Size(m)
}
func (m *UQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_UQuery.DiscardUnknown(m)
}

var xxx_messageInfo_UQuery proto.InternalMessageInfo

func (m *UQuery) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *UQuery) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UTruck struct {
	Truckid              *int32   `protobuf:"varint,1,req,name=truckid" json:"truckid,omitempty"`
	Status               *string  `protobuf:"bytes,2,req,name=status" json:"status,omitempty"`
	X                    *int32   `protobuf:"varint,3,req,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,4,req,name=y" json:"y,omitempty"`
	Seqnum               *int64   `protobuf:"varint,5,req,name=seqnum" json:"seqnum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UTruck) Reset()         { *m = UTruck{} }
func (m *UTruck) String() string { return proto.CompactTextString(m) }
func (*UTruck) ProtoMessage()    {}
func (*UTruck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{10}
}

func (m *UTruck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UTruck.Unmarshal(m, b)
}
func (m *UTruck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UTruck.Marshal(b, m, deterministic)
}
func (m *UTruck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTruck.Merge(m, src)
}
func (m *UTruck) XXX_Size() int {
	return xxx_messageInfo_UTruck.Size(m)
}
func (m *UTruck) XXX_DiscardUnknown() {
	xxx_messageInfo_UTruck.DiscardUnknown(m)
}

var xxx_messageInfo_UTruck proto.InternalMessageInfo

func (m *UTruck) GetTruckid() int32 {
	if m != nil && m.Truckid != nil {
		return *m.Truckid
	}
	return 0
}

func (m *UTruck) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *UTruck) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *UTruck) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *UTruck) GetSeqnum() int64 {
	if m != nil && m.Seqnum != nil {
		return *m.Seqnum
	}
	return 0
}

type UCommands struct {
	Pickups              []*UGoPickup  `protobuf:"bytes,1,rep,name=pickups" json:"pickups,omitempty"`
	Deliveries           []*UGoDeliver `protobuf:"bytes,2,rep,name=deliveries" json:"deliveries,omitempty"`
	Simspeed             *uint32       `protobuf:"varint,3,opt,name=simspeed" json:"simspeed,omitempty"`
	Disconnect           *bool         `protobuf:"varint,4,opt,name=disconnect" json:"disconnect,omitempty"`
	Queries              []*UQuery     `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	Acks                 []int64       `protobuf:"varint,6,rep,name=acks" json:"acks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UCommands) Reset()         { *m = UCommands{} }
func (m *UCommands) String() string { return proto.CompactTextString(m) }
func (*UCommands) ProtoMessage()    {}
func (*UCommands) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{11}
}

func (m *UCommands) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UCommands.Unmarshal(m, b)
}
func (m *UCommands) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UCommands.Marshal(b, m, deterministic)
}
func (m *UCommands) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UCommands.Merge(m, src)
}
func (m *UCommands) XXX_Size() int {
	return xxx_messageInfo_UCommands.Size(m)
}
func (m *UCommands) XXX_DiscardUnknown() {
	xxx_messageInfo_UCommands.DiscardUnknown(m)
}

var xxx_messageInfo_UCommands proto.InternalMessageInfo

func (m *UCommands) GetPickups() []*UGoPickup {
	if m != nil {
		return m.Pickups
	}
	return nil
}

func (m *UCommands) GetDeliveries() []*UGoDeliver {
	if m != nil {
		return m.Deliveries
	}
	return nil
}

func (m *UCommands) GetSimspeed() uint32 {
	if m != nil && m.Simspeed != nil {
		return *m.Simspeed
	}
	return 0
}

func (m *UCommands) GetDisconnect() bool {
	if m != nil && m.Disconnect != nil {
		return *m.Disconnect
	}
	return false
}

func (m *UCommands) GetQueries() []*UQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *UCommands) GetAcks() []int64 {
	if m != nil {
		return m.Acks
	}
	return nil
}

type UResponses struct {
	Completions          []*UFinished     `protobuf:"bytes,1,rep,name=completions" json:"completions,omitempty"`
	Delivered            []*UDeliveryMade `protobuf:"bytes,2,rep,name=delivered" json:"delivered,omitempty"`
	Finished             *bool            `protobuf:"varint,3,opt,name=finished" json:"finished,omitempty"`
	Acks                 []int64          `protobuf:"varint,4,rep,name=acks" json:"acks,omitempty"`
	Truckstatus          []*UTruck        `protobuf:"bytes,5,rep,name=truckstatus" json:"truckstatus,omitempty"`
	Error                []*UErr          `protobuf:"bytes,6,rep,name=error" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UResponses) Reset()         { *m = UResponses{} }
func (m *UResponses) String() string { return proto.CompactTextString(m) }
func (*UResponses) ProtoMessage()    {}
func (*UResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbb2b5a3bcea3396, []int{12}
}

func (m *UResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UResponses.Unmarshal(m, b)
}
func (m *UResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UResponses.Marshal(b, m, deterministic)
}
func (m *UResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UResponses.Merge(m, src)
}
func (m *UResponses) XXX_Size() int {
	return xxx_messageInfo_UResponses.Size(m)
}
func (m *UResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_UResponses.DiscardUnknown(m)
}

var xxx_messageInfo_UResponses proto.InternalMessageInfo

func (m *UResponses) GetCompletions() []*UFinished {
	if m != nil {
		return m.Completions
	}
	return nil
}

func (m *UResponses) GetDelivered() []*UDeliveryMade {
	if m != nil {
		return m.Delivered
	}
	return nil
}

func (m *UResponses) GetFinished() bool {
	if m != nil && m.Finished != nil {
		return *m.Finished
	}
	return false
}

func (m *UResponses) GetAcks() []int64 {
	if m != nil {
		return m.Acks
	}
	return nil
}

func (m *UResponses) GetTruckstatus() []*UTruck {
	if m != nil {
		return m.Truckstatus
	}
	return nil
}

func (m *UResponses) GetError() []*UErr {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*UInitTruck)(nil), "UInitTruck")
	proto.RegisterType((*UConnect)(nil), "UConnect")
	proto.RegisterType((*UConnected)(nil), "UConnected")
	proto.RegisterType((*UGoPickup)(nil), "UGoPickup")
	proto.RegisterType((*UFinished)(nil), "UFinished")
	proto.RegisterType((*UDeliveryMade)(nil), "UDeliveryMade")
	proto.RegisterType((*UDeliveryLocation)(nil), "UDeliveryLocation")
	proto.RegisterType((*UGoDeliver)(nil), "UGoDeliver")
	proto.RegisterType((*UErr)(nil), "UErr")
	proto.RegisterType((*UQuery)(nil), "UQuery")
	proto.RegisterType((*UTruck)(nil), "UTruck")
	proto.RegisterType((*UCommands)(nil), "UCommands")
	proto.RegisterType((*UResponses)(nil), "UResponses")
}

func init() { proto.RegisterFile("world_ups.proto", fileDescriptor_cbb2b5a3bcea3396) }

var fileDescriptor_cbb2b5a3bcea3396 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0x3e, 0x76, 0x37, 0x99, 0x6d, 0x0b, 0xf8, 0x80, 0x22, 0x40, 0x68, 0x09, 0x1c, 0x82,
	0xa8, 0x72, 0xe8, 0x09, 0x71, 0x40, 0x42, 0x05, 0x2a, 0x24, 0x2a, 0x51, 0xab, 0x39, 0x57, 0x51,
	0x62, 0x5a, 0xab, 0x1b, 0x3b, 0x6b, 0x27, 0xb4, 0xcb, 0xbf, 0xe4, 0x1f, 0xf0, 0x53, 0x90, 0x3f,
	0x92, 0x4d, 0x80, 0xcd, 0xcd, 0xcf, 0x93, 0x79, 0x6f, 0xde, 0x64, 0xc6, 0xf0, 0xe0, 0x8e, 0x8b,
	0x75, 0x79, 0xd5, 0xd6, 0x32, 0xad, 0x05, 0x6f, 0x78, 0xfc, 0x16, 0x20, 0xfb, 0xc2, 0x68, 0x73,
	0x29, 0xda, 0xe2, 0x16, 0x1d, 0x81, 0x4b, 0xcb, 0xc8, 0x59, 0xb9, 0xc9, 0x0c, 0xbb, 0xb4, 0x44,
	0x07, 0xe0, 0xdc, 0x47, 0xae, 0x86, 0xce, 0xbd, 0x42, 0xdb, 0xc8, 0x33, 0x68, 0x1b, 0x13, 0x08,
	0xb2, 0x53, 0xce, 0x18, 0x29, 0x1a, 0x14, 0xc1, 0x42, 0x13, 0xeb, 0x64, 0x27, 0xf1, 0x70, 0x07,
	0xd1, 0x4b, 0x98, 0x37, 0x8a, 0x5a, 0x46, 0xee, 0xca, 0x4b, 0x96, 0x27, 0xcb, 0x74, 0x27, 0x87,
	0x6d, 0x08, 0x3d, 0x81, 0x80, 0xca, 0x0f, 0x55, 0xfe, 0x93, 0x33, 0xcd, 0x1f, 0xe0, 0x1e, 0xc7,
	0xef, 0x01, 0x3a, 0x19, 0x52, 0x8e, 0x85, 0xdc, 0xa1, 0xd0, 0x63, 0x98, 0x0b, 0x22, 0xdb, 0x75,
	0xa3, 0xeb, 0x0d, 0xb1, 0x45, 0xf1, 0x05, 0x84, 0xd9, 0x19, 0xff, 0x46, 0x8b, 0xdb, 0xb6, 0x56,
	0xe9, 0x5a, 0xb2, 0x37, 0xd9, 0x41, 0x84, 0xc0, 0xbf, 0xbb, 0xa1, 0xa5, 0x35, 0xab, 0xcf, 0x8a,
	0x52, 0x92, 0x0d, 0x6b, 0x2b, 0x5d, 0x94, 0x87, 0x2d, 0x8a, 0x37, 0x10, 0x66, 0x9f, 0x29, 0xa3,
	0xf2, 0xc6, 0x54, 0xb4, 0x87, 0x72, 0xa2, 0x79, 0x9a, 0xba, 0xc9, 0x9b, 0x56, 0x46, 0xbe, 0xa9,
	0xd6, 0xa0, 0x81, 0xe4, 0x6c, 0x24, 0x79, 0x05, 0x87, 0xd9, 0x47, 0xb2, 0xa6, 0x3f, 0x88, 0xd8,
	0x9e, 0xe7, 0x25, 0x99, 0x90, 0x7d, 0x06, 0x61, 0x9d, 0x17, 0xb7, 0xf9, 0x35, 0xb1, 0x76, 0x3c,
	0xbc, 0xbb, 0xd8, 0xeb, 0xe9, 0x1c, 0x1e, 0xf5, 0x02, 0x5f, 0x79, 0x91, 0x37, 0x94, 0xb3, 0x31,
	0x95, 0xf3, 0x37, 0xd5, 0xd4, 0x70, 0x30, 0x80, 0xec, 0x8c, 0x5b, 0xc2, 0x89, 0x62, 0x53, 0x08,
	0x2c, 0x61, 0x37, 0x20, 0x28, 0xfd, 0xa7, 0x0e, 0xdc, 0x7f, 0xb3, 0xb7, 0xfc, 0x4b, 0xf0, 0xb3,
	0x4f, 0x42, 0xa0, 0x87, 0xe0, 0x11, 0x21, 0xb4, 0x4a, 0x88, 0xd5, 0x11, 0xc5, 0x70, 0xc0, 0x05,
	0xbd, 0xa6, 0xcc, 0xe6, 0x99, 0x8e, 0x8c, 0xee, 0xf6, 0xb2, 0xbe, 0x83, 0x79, 0x76, 0xd1, 0x12,
	0xb1, 0x9d, 0x70, 0xb0, 0xcb, 0x75, 0x47, 0xb9, 0x0c, 0xe6, 0x99, 0x59, 0xaa, 0xe9, 0x5c, 0x33,
	0x05, 0xee, 0x68, 0x0a, 0x74, 0x67, 0xbd, 0x51, 0x67, 0xfd, 0xe1, 0xe4, 0xfc, 0x6f, 0x42, 0x7e,
	0x39, 0x10, 0x66, 0xa7, 0xbc, 0xaa, 0x72, 0x56, 0x4a, 0xf4, 0x0a, 0x16, 0xb5, 0x1e, 0x79, 0x19,
	0x39, 0xba, 0xad, 0x90, 0xf6, 0x5b, 0x80, 0xbb, 0x10, 0x7a, 0x03, 0x50, 0x9a, 0x5e, 0x53, 0x32,
	0x58, 0xd0, 0xfe, 0xc7, 0xe1, 0x41, 0x58, 0x2d, 0xa9, 0xa4, 0x95, 0xac, 0x09, 0x29, 0x23, 0x6f,
	0xe5, 0x24, 0x87, 0xb8, 0xc7, 0xe8, 0x39, 0x40, 0x49, 0x65, 0x61, 0xd6, 0x34, 0xf2, 0x57, 0x4e,
	0x12, 0xe0, 0xc1, 0x0d, 0x7a, 0x01, 0x8b, 0x4d, 0x6b, 0x54, 0x66, 0x5a, 0x65, 0x91, 0x9a, 0xc6,
	0xe2, 0xee, 0x5e, 0x2d, 0x60, 0xae, 0x9e, 0x89, 0xf9, 0xca, 0x4b, 0x3c, 0xac, 0xcf, 0xf1, 0x6f,
	0x07, 0x20, 0xc3, 0x44, 0xd6, 0x9c, 0x49, 0x22, 0xd1, 0x31, 0x2c, 0x0b, 0x5e, 0xd5, 0x6b, 0xa2,
	0x86, 0x62, 0x60, 0xac, 0xdb, 0x45, 0x3c, 0x0c, 0xa3, 0x63, 0x08, 0x6d, 0xf5, 0xa4, 0xb4, 0xde,
	0x8e, 0xd2, 0xd1, 0x12, 0xe1, 0xdd, 0x07, 0xca, 0xdd, 0x77, 0x4b, 0xa3, 0xdd, 0x05, 0xb8, 0xc7,
	0x7d, 0x69, 0xfe, 0xae, 0x34, 0xf4, 0x1a, 0x96, 0xe6, 0xf1, 0x32, 0xff, 0xaf, 0x77, 0x65, 0x1e,
	0xb6, 0x61, 0x0c, 0x3d, 0x85, 0x19, 0x11, 0x82, 0x0b, 0x6d, 0x6d, 0x79, 0x32, 0x4b, 0xd5, 0xa4,
	0x62, 0x73, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x49, 0x38, 0xd3, 0x1b, 0x91, 0x05, 0x00, 0x00,
}
