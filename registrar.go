package gateway

import (
	"container/list"

	"github.com/freeconf/restconf/device"
	"github.com/freeconf/yang/nodeutil"
)

type Registration struct {
	DeviceId string
	Address  string
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
	listeners *list.List
}

func NewLocalRegistrar() *LocalRegistrar {
	return &LocalRegistrar{
		regs:      make(map[string]Registration),
		listeners: list.New(),
	}
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

func (gw *LocalRegistrar) RegisterDevice(deviceId string, address string) {
	reg := Registration{Address: address, DeviceId: deviceId}
	gw.regs[deviceId] = reg
	gw.updateListeners(reg)
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
