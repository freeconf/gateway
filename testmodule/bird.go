package testmodule

import (
	"github.com/freeconf/restconf/device"
	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/node"
	"github.com/freeconf/yang/nodeutil"
	"github.com/freeconf/yang/parser"
	"github.com/freeconf/yang/source"
)

type Bird struct {
	Name     string
	Wingspan int
	Species  *Species
}

type Species struct {
	Name  string
	Class string
}

func BirdDevice(ypath source.Opener, json string) (*device.Local, map[string]*Bird) {
	d := device.New(ypath)
	b, birds := BirdBrowser(ypath, json)
	d.AddBrowser(b)
	if json != "" {
		if err := b.Root().UpsertFrom(nodeutil.ReadJSON(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return d, birds
}

func BirdBrowser(ypath source.Opener, json string) (*node.Browser, map[string]*Bird) {
	data := make(map[string]*Bird)
	b := node.NewBrowser(BirdModule(ypath), BirdNode(data))
	if json != "" {
		if err := b.Root().UpsertFrom(nodeutil.ReadJSON(json)).LastErr; err != nil {
			panic(err)
		}
	}
	return b, data
}

func BirdModule(ypath source.Opener) *meta.Module {
	return parser.RequireModule(ypath, "bird")
}

func BirdNode(birds map[string]*Bird) node.Node {
	return &nodeutil.Basic{
		OnChild: func(r node.ChildRequest) (node.Node, error) {
			switch r.Meta.Ident() {
			case "bird":
				return nodeutil.ReflectList(birds), nil
			}
			return nil, nil
		},
	}
}
