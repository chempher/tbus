// Code generated by protoc-gen-go.
// source: tbus/bus.proto
// DO NOT EDIT!

/*
Package tbus is a generated protocol buffer package.

It is generated from these files:
	tbus/bus.proto
	tbus/error.proto
	tbus/led.proto
	tbus/motor.proto
	tbus/servo.proto

It has these top-level messages:
	DeviceInfo
	BusEnumeration
	Error
	LEDPowerState
	MotorDriveState
	MotorBrakeState
	ServoPosition
*/
package tbus

import prot "github.com/robotalks/tbus/go/tbus/protocol"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceInfo struct {
	Address  uint32            `protobuf:"varint,1,opt,name=address" json:"address,omitempty"`
	ClassId  uint32            `protobuf:"varint,2,opt,name=class_id,json=classId" json:"class_id,omitempty"`
	DeviceId uint32            `protobuf:"varint,3,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	Labels   map[string]string `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DeviceInfo) Reset()                    { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string            { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()               {}
func (*DeviceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeviceInfo) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type BusEnumeration struct {
	Devices []*DeviceInfo `protobuf:"bytes,1,rep,name=devices" json:"devices,omitempty"`
}

func (m *BusEnumeration) Reset()                    { *m = BusEnumeration{} }
func (m *BusEnumeration) String() string            { return proto.CompactTextString(m) }
func (*BusEnumeration) ProtoMessage()               {}
func (*BusEnumeration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BusEnumeration) GetDevices() []*DeviceInfo {
	if m != nil {
		return m.Devices
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceInfo)(nil), "tbus.DeviceInfo")
	proto.RegisterType((*BusEnumeration)(nil), "tbus.BusEnumeration")
}

func init() { proto.RegisterFile("tbus/bus.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x90, 0xcd, 0x6a, 0xb3, 0x40,
	0x14, 0x86, 0x31, 0xfa, 0x25, 0xf1, 0x84, 0x2f, 0x84, 0x21, 0x14, 0x63, 0xba, 0x08, 0xae, 0x42,
	0x17, 0x23, 0xa4, 0x5d, 0xb4, 0xa5, 0x50, 0x08, 0x75, 0x21, 0x64, 0xe5, 0x0d, 0x14, 0x75, 0x26,
	0x41, 0xaa, 0x8e, 0x38, 0x4e, 0xc0, 0x5d, 0x6f, 0xaa, 0xd7, 0xd1, 0x5b, 0x2a, 0x73, 0x26, 0xd2,
	0x9f, 0xcd, 0x30, 0xef, 0x79, 0xcf, 0xcf, 0x73, 0x0e, 0xcc, 0xbb, 0x4c, 0xc9, 0x30, 0x53, 0x92,
	0x36, 0xad, 0xe8, 0x04, 0x71, 0xb4, 0xf6, 0xd7, 0x27, 0x21, 0x4e, 0x25, 0x0f, 0x31, 0x96, 0xa9,
	0x63, 0xc8, 0xab, 0xa6, 0xeb, 0x4d, 0x8a, 0xbf, 0xc2, 0x92, 0x5c, 0x54, 0x95, 0xa8, 0x43, 0xd1,
	0x74, 0x85, 0xa8, 0x2f, 0xd5, 0xc1, 0xa7, 0x05, 0xf0, 0xc2, 0xcf, 0x45, 0xce, 0xe3, 0xfa, 0x28,
	0x88, 0x07, 0x93, 0x94, 0xb1, 0x96, 0x4b, 0xe9, 0x59, 0x1b, 0x6b, 0xfb, 0x3f, 0x19, 0x24, 0x59,
	0xc1, 0x34, 0x2f, 0x53, 0x29, 0x5f, 0x0b, 0xe6, 0x8d, 0x8c, 0x85, 0x3a, 0x66, 0x64, 0x0d, 0x2e,
	0xc3, 0x16, 0xda, 0xb3, 0xd1, 0x9b, 0x9a, 0x40, 0xcc, 0xc8, 0x1d, 0x8c, 0xcb, 0x34, 0xe3, 0xa5,
	0xf4, 0x9c, 0x8d, 0xbd, 0x9d, 0xed, 0xae, 0xa9, 0x86, 0xa1, 0xdf, 0x33, 0xe9, 0x01, 0xed, 0xa8,
	0xee, 0xda, 0x3e, 0xb9, 0xe4, 0xfa, 0x0f, 0x30, 0xfb, 0x11, 0x26, 0x0b, 0xb0, 0xdf, 0x78, 0x8f,
	0x48, 0x6e, 0xa2, 0xbf, 0x64, 0x09, 0xff, 0xce, 0x69, 0xa9, 0x38, 0xb2, 0xb8, 0x89, 0x11, 0x8f,
	0xa3, 0x7b, 0x2b, 0x78, 0x82, 0xf9, 0x5e, 0xc9, 0xa8, 0x56, 0x15, 0x6f, 0x53, 0xbd, 0x2a, 0xb9,
	0x81, 0x89, 0xc1, 0xd1, 0x4b, 0x69, 0x86, 0xc5, 0x5f, 0x86, 0x64, 0x48, 0xd8, 0x1d, 0xc0, 0xde,
	0x2b, 0x49, 0x9e, 0xc1, 0x1d, 0x3a, 0x70, 0x72, 0x45, 0xcd, 0x71, 0xe9, 0x70, 0x5c, 0x1a, 0xe9,
	0xe3, 0xfa, 0x4b, 0xd3, 0xe6, 0xf7, 0xb4, 0xc0, 0x79, 0xff, 0xf0, 0x2c, 0x1f, 0xdf, 0x6c, 0x8c,
	0x15, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0xbd, 0xda, 0x7b, 0xb4, 0x01, 0x00, 0x00,
}

//
// GENERTED FROM tbus/bus.proto, DO NOT EDIT
//

// BusClassID is the class ID of Bus
const BusClassID uint32 = 0x0001

// BusLogic defines the logic interface
type BusLogic interface {
    DeviceLogic
    MsgRouter
    Enumerate() (*BusEnumeration, error)
}

// BusDev is the device
type BusDev struct {
    DeviceBase
    Logic BusLogic
}

// NewBusDev creates a new device
func NewBusDev(logic BusLogic) *BusDev {
    d := &BusDev{Logic: logic}
    d.Info.ClassId = BusClassID
    logic.SetDevice(d)
    return d
}

// SendMsg implements Device
func (d *BusDev) SendMsg(msg *prot.Msg) (err error) {
    if msg.Head.NeedRoute() {
        return d.Logic.(MsgRouter).RouteMsg(msg)
    }
    var reply proto.Message
    switch msg.Body.Flag {
    case 1: // Enumerate
        reply, err = d.Logic.Enumerate()
    default:
        err = ErrInvalidMethod
    }
    return d.Reply(msg.Head.MsgID, reply, err)
}

// SetDeviceID sets device id
func (d *BusDev) SetDeviceID(id uint32) *BusDev {
    d.Info.DeviceId = id
    return d
}

// BusCtl is the device controller
type BusCtl struct {
    Controller
}

// NewBusCtl creates controller for Bus
func NewBusCtl(master Master) *BusCtl {
    c := &BusCtl{}
    c.Master = master
    return c
}

// SetAddress sets routing address for target device
func (c *BusCtl) SetAddress(addrs []uint8) *BusCtl {
    c.Address = addrs
    return c
}

// Enumerate wraps class Bus
func (c *BusCtl) Enumerate() (*BusEnumeration, error) {
    reply := &BusEnumeration{}
    err := c.Invoke(1, nil, reply)
    return reply, err
}

