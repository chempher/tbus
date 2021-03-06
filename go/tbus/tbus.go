package tbus

import (
	"fmt"
	"io"
	"time"

	proto "github.com/golang/protobuf/proto"
	prot "github.com/robotalks/tbus/go/tbus/protocol"
)

var (
	// ErrInvalidMethod indicates method index is invalid
	ErrInvalidMethod = fmt.Errorf("invalid method index")
	// ErrInvalidAddr indicates address doesn't map to a device
	ErrInvalidAddr = fmt.Errorf("invalid address")
	// ErrRouteNotSupport indicates the device doesn't support routing
	ErrRouteNotSupport = fmt.Errorf("route not supported")
	// ErrRecvAborted indicates the receiving is cancelled
	ErrRecvAborted = fmt.Errorf("receiving aborted")
	// ErrRecvEnd indicates the receiving is ended
	ErrRecvEnd = io.EOF
	// ErrAddrNotAvail indicates no more address can be allocated
	ErrAddrNotAvail = fmt.Errorf("address not available")
	// ErrNoAssocDevice indicates a logic is not associated with device
	ErrNoAssocDevice = fmt.Errorf("logic not associated with device")
	// ErrInvalidSender indicates sender is unavailable
	ErrInvalidSender = fmt.Errorf("sender not available")
)

// MsgReceiver provides a message chan for read
type MsgReceiver interface {
	MsgChan() <-chan prot.Msg
}

// MsgSender writes a message
type MsgSender interface {
	SendMsg(*prot.Msg) error
}

// MsgRouter is able to route a message
type MsgRouter interface {
	RouteMsg(*prot.Msg) error
}

// BusPort is the device side of the bus
type BusPort interface {
	MsgSender
}

// Bus defines a bus instance
type Bus interface {
	Plug(Device) error
	Unplug(Device) error
}

// Device defines a device instance
type Device interface {
	MsgSender
	DeviceInfo() DeviceInfo
	AttachTo(BusPort, uint8)
	BusPort() BusPort
}

// DeviceLogic implements device functions
type DeviceLogic interface {
	SetDevice(Device)
}

// Master is the bus master
type Master interface {
	Invoke(method uint8, params proto.Message, addrs []uint8) (Invocation, error)
}

// Invocation represents the result of method invocation
type Invocation interface {
	MsgReceiver
	MessageID() uint32
	Ignore()
}

// MsgReader reads a message from a receiver
type MsgReader struct {
	Timeout    time.Duration
	CancelChan <-chan struct{}
}

// SetTimeout sets the timeout value
func (r *MsgReader) SetTimeout(timeout time.Duration) *MsgReader {
	r.Timeout = timeout
	return r
}

// SetCancelChan sets the cancellation chan
func (r *MsgReader) SetCancelChan(ch <-chan struct{}) *MsgReader {
	r.CancelChan = ch
	return r
}

// ReadMsg reads a message
func (r *MsgReader) ReadMsg(recv MsgReceiver) (*prot.Msg, error) {
	var msg prot.Msg
	var ok bool
	if r.Timeout == 0 {
		if r.CancelChan != nil {
			select {
			case <-r.CancelChan:
				return nil, ErrRecvAborted
			case msg, ok = <-recv.MsgChan():
				break
			}
		} else {
			msg, ok = <-recv.MsgChan()
		}
	} else if r.CancelChan != nil {
		select {
		case <-time.After(r.Timeout):
			return nil, nil
		case _, _ = <-r.CancelChan:
			return nil, ErrRecvAborted
		case msg, ok = <-recv.MsgChan():
			break
		}
	}
	if !ok {
		return nil, ErrRecvEnd
	}
	return &msg, nil
}

// ReadReply reads message and decode into reply
func (r *MsgReader) ReadReply(recv MsgReceiver, reply proto.Message) error {
	msg, err := r.ReadMsg(recv)
	if err != nil {
		return err
	}
	if (msg.Body.Flag & prot.BodyError) != 0 {
		replyErr := &Error{}
		if err = proto.Unmarshal(msg.Body.Data, replyErr); err != nil {
			return err
		}
		return replyErr
	}
	if reply != nil {
		err = proto.Unmarshal(msg.Body.Data, reply)
	}
	return err
}
