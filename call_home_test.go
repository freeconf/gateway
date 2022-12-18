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
	ypath := source.Dir("./yang")

	// use a single device to hols all modules
	dev := device.New(ypath)
	proto := func(string) (device.Device, error) {
		return dev, nil
	}

	// gateway to receive call-home requests
	gw := NewLocalRegistrar()
	fc.AssertEqual(t, nil, InstallRegistrar(gw, dev))
	fc.AssertEqual(t, nil, dev.Add("fc-gateway", RegistrarNode(gw)))
	fc.AssertEqual(t, nil, dev.Add("fc-call-home-server", CallHomeServer(gw)))

	// device looking to call home to gateway
	caller := restconf.NewCallHome(proto)
	options := caller.Options()
	options.DeviceId = "x"
	options.Address = "gw"
	options.LocalAddress = "device"

	count := 0
	caller.OnRegister(func(d device.Device, update restconf.RegisterUpdate) {
		count++
	})
	fc.AssertEqual(t, nil, caller.ApplyOptions(options))
	fc.AssertEqual(t, 1, count)
	fc.AssertEqual(t, 1, gw.RegistrationCount())
}
