package main

import (
	"flag"

	"github.com/freeconf/gateway"
	"github.com/freeconf/restconf"

	"github.com/freeconf/restconf/client"
	"github.com/freeconf/restconf/device"
	"github.com/freeconf/yang/source"

	"github.com/freeconf/yang/fc"
)

// Management Gateway.  Serve management functions to available services.
//
// Then open web browser to
//   http://localhost:8080/
//

var startup = flag.String("startup", "startup.json", "startup configuration file.")
var verbose = flag.Bool("verbose", false, "verbose")
var ypathStr = flag.String("ypath", "yang", "location or locations (separated by ':') of yang files")

func main() {
	flag.Parse()
	fc.DebugLog(*verbose)

	ypath := source.Path(*ypathStr)

	// Even though this is a server component, we still organize things thru a device
	// because this proxy will appear like a "Device" to application management systems
	// "northbound"" representing all the devices that are "southbound".
	d := device.New(ypath)

	// We "wrap" each device with a device that splits CRUD operations
	// to local store AND the original device.  This gives us transparent
	// persistance of device data w/o altering the device API.
	reg := gateway.NewLocalRegistrar(client.ProtocolHandler(ypath))
	chkErr(gateway.InstallRegistrar(reg, d))

	// Add RESTCONF service, if you had other protocols to add/replace
	// you could do that here
	mgmt := restconf.NewServer(d)

	// Let RESTCONF know it's proxy for registered devices
	mgmt.ServeDevices(reg)

	// bootstrap config for all local modules
	chkErr(d.ApplyStartupConfigFile(*startup))

	// Wait for cntrl-c...
	select {}
}

func chkErr(err error) {
	if err != nil {
		panic(err)
	}
}
