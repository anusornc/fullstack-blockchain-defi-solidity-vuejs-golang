package services

import (
	"github.com/anusornc/fullstackblockchain/consts"
	"github.com/anusornc/fullstackblockchain/utils"
)

type Deployer struct {
	cfg consts.IConfig
}

func NewDeployer(cfg consts.IConfig) *Deployer {
	return &Deployer{
		cfg: cfg,
	}
}

func (svc *Deployer) Deploy() error {
	utils.LogM(string(svc.cfg.Network()))
	utils.LogM("Deploy")
	return nil
}
