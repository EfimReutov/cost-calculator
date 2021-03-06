package app

import (
	"cost-calculator/config"
	"cost-calculator/server/grpc"
	"cost-calculator/server/rest"
	"cost-calculator/store/postgres"
	"errors"
	"fmt"
)

// Run function configures and starts a postgres connection depending on the type of server
func Run() error {
	cfg, err := config.LoadCfg()
	if err != nil {
		return err
	}

	pg, err := postgres.NewDB(cfg)
	if err != nil {
		return err
	}
	defer pg.Close()

	if cfg.ServerType == config.ServerTypeGRPC {
		return grpc.Run(pg, fmt.Sprintf(":%d", cfg.ServicePort))
	} else if cfg.ServerType == config.ServerTypeREST {
		return rest.Run(pg, cfg)
	} else {
		return errors.New(fmt.Sprintf("invalid server type: %s", cfg.ServerType))
	}
}
