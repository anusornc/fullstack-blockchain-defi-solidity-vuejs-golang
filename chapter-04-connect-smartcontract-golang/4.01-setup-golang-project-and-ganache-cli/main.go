package main
import "fmt"
import (
	"github.com/anusornc/fullstackblockchain/consts"
	"github.com/anusornc/fullstackblockchain/services"
)

func main() {
	cfg := consts.NewConfig()
	serviceID := cfg.ServiceID()
	fmt.Printf("Starting serviceID = %s	", serviceID)	
	switch serviceID {
	case "reader":
		services.NewReader(cfg).Read()
	case "deployer":
		services.NewDeployer(cfg).Deploy()
	}
}
