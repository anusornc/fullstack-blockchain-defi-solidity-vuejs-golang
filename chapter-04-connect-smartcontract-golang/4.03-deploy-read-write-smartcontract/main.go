package main

import (
	"github.com/anusornc/fullstackblockchain/consts"
	"github.com/anusornc/fullstackblockchain/services"
)

func main() {
	cfg := consts.NewConfig()
	serviceID := cfg.ServiceID()
	switch serviceID {
	case "reader":
		services.NewReader(cfg).Read()
	case "writer":
		services.NewWriter(cfg).Write()
	case "deployer":
		services.NewDeployer(cfg).Deploy()
	}
}
