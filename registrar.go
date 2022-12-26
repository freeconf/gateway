package gateway

import (
	"container/list"

	"github.com/freeconf/restconf/device"
	"github.com/freeconf/yang/nodeutil"
)

type Registration struct {
	DeviceId string
	Address  string
	Device   device.Device
}

type Registrar interface {
	RegistrationCount() int
	LookupRegistration(deviceId string) (Registration, bool)
	RegisterDevice(deviceId string, address string)
	OnRegister(l RegisterListener) nodeutil.Subscription
}

type RegisterListener func(Registration)

type LocalRegistrar struct {
	regs      map[string]Registration
	proto     device.ProtocolHandler
	listeners *list.List
}

func NewLocalRegistrar(proto device.ProtocolHandler) *LocalRegistrar {
	return &LocalRegistrar{
		proto:     proto,
		regs:      make(map[string]Registration),
		listeners: list.New(),
	}
}

func (gw *LocalRegistrar) Device(deviceId string) (device.Device, error) {
	return gw.regs[deviceId].Device, nil
}

func InstallRegistrar(reg *LocalRegistrar, dev *device.Local) error {
	if err := dev.Add("fc-gateway", RegistrarNode(reg)); err != nil {
		return err
	}
	return dev.Add("fc-call-home-server", CallHomeServer(reg))
}

func (gw *LocalRegistrar) LookupRegistration(deviceId string) (Registration, bool) {
	found, reg := gw.regs[deviceId]
	return found, reg
}

func (gw *LocalRegistrar) RegisterDevice(deviceId string, address string) error {
	dev, err := gw.proto(address)
	if err != nil {
		return err
	}
	reg := Registration{Address: address, DeviceId: deviceId, Device: dev}
	gw.regs[deviceId] = reg
	gw.updateListeners(reg)
	return nil
}

func (gw *LocalRegistrar) updateListeners(reg Registration) {
	p := gw.listeners.Front()
	for p != nil {
		p.Value.(RegisterListener)(reg)
		p = p.Next()
	}
}

func (gw *LocalRegistrar) RegistrationCount() int {
	return len(gw.regs)
}

func (gw *LocalRegistrar) OnRegister(l RegisterListener) nodeutil.Subscription {
	return nodeutil.NewSubscription(gw.listeners, gw.listeners.PushBack(l))
}
