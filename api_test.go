package gateway

import (
	"testing"

	"github.com/freeconf/restconf/callhome"
	"github.com/freeconf/restconf/device"
	"github.com/freeconf/yang/fc"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/source"
)

func newTestApp(t *testing.T) (*LocalRegistrar, *device.Local) {
	fc.DebugLog(true)
	ypath := source.Dir("./yang")
	local := device.New(ypath)
	proto := func(address string) (device.Device, error) {
		return local, nil
	}
	app := NewLocalRegistrar(proto)
	InstallRegistrar(app, local)
	return app, local
}

func TestGatewayApi(t *testing.T) {
	app, dev := newTestApp(t)
	app.RegisterDevice("x", "z")

	api, err := dev.Browser("fc-gateway")
	fc.AssertEqual(t, nil, err)
	actual, err := nodeutil.WriteJSON(api.Root())
	fc.AssertEqual(t, nil, err)
	fc.AssertEqual(t, `{"fc-gateway:registration":[{"deviceId":"x","address":"z"}]}`, actual)
}

func TestCallHomeApi(t *testing.T) {
	gw, dev := newTestApp(t)

	// gateway to receive call-home requests
	fc.AssertEqual(t, nil, InstallRegistrar(gw, dev))
	serverCount := 0
	gw.OnRegister(func(r Registration) {
		serverCount++
	})

	fc.AssertEqual(t, nil, dev.Add("fc-call-home-server", CallHomeServer(gw)))

	// using the client driver to make a calll home request to gateway
	caller := callhome.New(gw.proto)
	options := caller.Options()
	options.DeviceId = "x"
	options.Address = "gw"
	options.LocalAddress = "device"

	clientCount := 0
	caller.OnRegister(func(d device.Device, update callhome.RegisterUpdate) {
		clientCount++
	})

	fc.AssertEqual(t, nil, caller.ApplyOptions(options))
	fc.AssertEqual(t, 1, clientCount)
	fc.AssertEqual(t, 1, gw.RegistrationCount())

	fc.AssertEqual(t, 1, serverCount)
}
