package tbus

import (
	"io"
	"net"
	"sync"

	prot "github.com/evo-bots/tbus/go/tbus/protocol"
	proto "github.com/golang/protobuf/proto"
)

// MsgStreamer read/write msg using stream
type MsgStreamer struct {
	Writer io.Writer
	lock   sync.Mutex
}

// NewMsgStreamer creates a new MsgStreamer
func NewMsgStreamer(writer io.Writer) *MsgStreamer {
	return &MsgStreamer{Writer: writer}
}

// SendMsg implements MsgSender
func (s *MsgStreamer) SendMsg(msg *prot.Msg) (err error) {
	s.lock.Lock()
	_, err = s.Writer.Write(msg.Head.Raw)
	if err == nil {
		_, err = s.Writer.Write(msg.Body.Raw)
	}
	s.lock.Unlock()
	return
}

// DecodeStream decode msgs from stream and pipe to sender
func DecodeStream(reader io.Reader, sender MsgSender) error {
	for {
		msg, err := prot.Decode(reader)
		if err == io.EOF {
			return nil
		}
		if err == nil {
			err = sender.SendMsg(&msg)
		}
		if err != nil {
			return err
		}
	}
}

// StreamDevice sends msg to a writer
type StreamDevice struct {
	MsgStreamer
	Info   DeviceInfo
	Reader io.Reader

	busPort BusPort
	init    bool
	initErr error
}

// NewStreamDevice creates a stream device
func NewStreamDevice(classID uint32, rw io.ReadWriter) *StreamDevice {
	dev := &StreamDevice{Reader: rw, init: true}
	dev.Writer = rw
	dev.Info.ClassId = classID
	return dev
}

// Address returns current address
func (d *StreamDevice) Address() uint8 {
	return uint8(d.Info.Address)
}

// ClassID implements Device
func (d *StreamDevice) ClassID() uint32 {
	return d.Info.ClassId
}

// DeviceID implements Device
func (d *StreamDevice) DeviceID() uint32 {
	return d.Info.DeviceId
}

// BusPort implements Device
func (d *StreamDevice) BusPort() BusPort {
	return d.busPort
}

// SetDeviceID sets device ID
func (d *StreamDevice) SetDeviceID(id uint32) *StreamDevice {
	d.Info.DeviceId = id
	return d
}

// AttachTo implements Device
func (d *StreamDevice) AttachTo(busPort BusPort, addr uint8) {
	d.busPort = busPort
	d.Info.Address = uint32(addr)
	if d.init && busPort != nil {
		// during init send attach info to stream
		encoded, _ := proto.Marshal(&d.Info)
		msg, err := prot.EncodeAsMsg(nil, 0, 0, encoded)
		if err == nil {
			err = d.SendMsg(msg)
		}
		if err != nil {
			d.initErr = err
		}
	}
	d.init = false
}

// Run pipes remote msg to bus port
func (d *StreamDevice) Run() error {
	if d.initErr != nil {
		return d.initErr
	}
	return DecodeStream(d.Reader, d.busPort)
}

// StreamBusPort exposes a device to remote
type StreamBusPort struct {
	MsgStreamer
	Reader io.Reader
	Device Device
}

// NewStreamBusPort creates a stream bus port
func NewStreamBusPort(rw io.ReadWriter, dev Device, addr uint8) *StreamBusPort {
	p := &StreamBusPort{Reader: rw}
	p.Writer = rw
	p.Device = dev
	p.Device.AttachTo(p, addr)
	return p
}

// Run pipes remote msg to device
func (p *StreamBusPort) Run() error {
	return DecodeStream(p.Reader, p.Device)
}

// Dial is abstract remote connector
type Dial func() (io.ReadWriteCloser, error)

// NetBusPort hosts a device over network
type NetBusPort struct {
	Dialer Dial
	Device Device

	conn io.ReadWriteCloser
}

// NewNetBusPort creates a NetBusPort
func NewNetBusPort(dev Device, dialer Dial) *NetBusPort {
	return &NetBusPort{Dialer: dialer, Device: dev}
}

// Conn returns current connection
func (p *NetBusPort) Conn() io.ReadWriteCloser {
	return p.conn
}

// Run connect to remote and host the device
func (p *NetBusPort) Run() error {
	conn, err := p.Dialer()
	if err == nil {
		p.conn = conn
		err = p.runConn()
		p.conn = nil
	}
	return err
}

func (p *NetBusPort) runConn() error {
	defer p.conn.Close()

	// the first message is sending device info for bus attachment
	info := &DeviceInfo{
		Address:  0,
		ClassId:  p.Device.ClassID(),
		DeviceId: p.Device.DeviceID(),
	}
	encoded, err := proto.Marshal(info)
	if err != nil {
		return err
	}
	if _, err = prot.Encode(p.conn, nil, 0, 0, encoded); err != nil {
		return err
	}

	// expect a bus attachment
	msg, err := prot.Decode(p.conn)
	if err != nil {
		return err
	}
	info = &DeviceInfo{}
	if err = proto.Unmarshal(msg.Body.Data, info); err != nil {
		return err
	}

	// do a bus attach
	port := NewStreamBusPort(p.conn, p.Device, uint8(info.Address))
	err = port.Run()
	p.Device.AttachTo(nil, 0)
	return err
}

// NetDeviceHost accepts connections from remote NetBusPort
// and creates StreamDevice for each connection.
type NetDeviceHost struct {
	Listener net.Listener
	acceptCh chan *StreamDevice
}

// NewNetDeviceHost creates a new net device host
func NewNetDeviceHost(listener net.Listener) *NetDeviceHost {
	return &NetDeviceHost{
		Listener: listener,
		acceptCh: make(chan *StreamDevice),
	}
}

// AcceptChan returns chan for accepted device
func (h *NetDeviceHost) AcceptChan() <-chan *StreamDevice {
	return h.acceptCh
}

// Run starts accepting device connections
func (h *NetDeviceHost) Run() error {
	for {
		conn, err := h.Listener.Accept()
		if err != nil {
			return err
		}
		msg, err := prot.Decode(conn)
		if err != nil {
			conn.Close()
		} else {
			info := &DeviceInfo{}
			err = proto.Unmarshal(msg.Body.Data, info)
			if err != nil {
				conn.Close()
			} else {
				device := NewStreamDevice(info.ClassId, conn)
				device.SetDeviceID(info.DeviceId)
				h.acceptCh <- device
			}
		}
	}
}