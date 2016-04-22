package install

import (
	"fmt"
//	"io"

	log "github.com/Sirupsen/logrus"
	"github.com/megamsys/libgo/action"
//	"github.com/megamsys/libgo/exec"
//	"strings"
)

type runActionsArgs struct {

  username    string
  password    string
	host        string
	bridgename      string
  phydev       string
	network      string
	netmask      string
	gateway      string
	dnsname1     string
	dnsname2     string
}


var CheckHostInfo = action.Action{
	Name: "CheckHostInfo",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		args := ctx.Params[0].(runActionsArgs)

		log.Debugf("Host  %s ", args.host)
		fmt.Println("=====================actin==================")
		fmt.Println(args.host)
		var m VerticeInstall
		m.GetHostInfo(args.host,args.username,args.password)
		    fmt.Println()
    log.Debugf("Verified [%s] host ", args.host)
    return &args ,nil

	},
	Backward: func(ctx action.BWContext) {

	},
}

var CreateBridgeAction = action.Action{
	Name: "CreateBridge",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		args := ctx.Params[0].(runActionsArgs)

		log.Debugf("Host  %s ", args.host)
		fmt.Println("=====================actin123==================")
		fmt.Println(args.host)
		var m CreateBridge
		m.Bridge(args.bridgename, args.phydev, args.network, args.netmask, args.gateway, args.dnsname1, args.dnsname2, args.host, args.username, args.password)
		    fmt.Println()
    log.Debugf("Verified [%s] host ", args.host)
    return &args ,nil

	},
	Backward: func(ctx action.BWContext) {

	},
}
