package gateway

import (
	"testing"

	"github.com/freeconf/restconf"
	"github.com/freeconf/restconf/device"
	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/source"
)

func TestCallHome(t *testing.T) {
	fc.DebugLog(true)

	registrar := NewLocalRegistrar()
	ypath := source.Dir("./yang")
	regDevice := device.New(ypath)
	if err := regDevice.Add("fc-gateway", RegistrarNode(registrar)); err != nil {
		t.Error(err)
	}
	caller := restconf.NewCallHome(func(string) (device.Device, error) {
		return regDevice, nil
	})
	options := caller.Options()
	options.DeviceId = "x"
	options.Address = "north"
	options.LocalAddress = "south"
	var gotUpdate bool
	caller.OnRegister(func(d device.Device, update restconf.RegisterUpdate) {
		gotUpdate = true
	})
	caller.ApplyOptions(options)
	if !gotUpdate {
		t.Error("no update recieved")
	}
	fc.AssertEqual(t, 1, registrar.RegistrationCount())
}
