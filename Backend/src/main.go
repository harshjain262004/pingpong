package main

import (
	"github.com/harshjain262004/pingpong/Backend/src/utils"
)

func main() {
	cfg, err := utils.GetConfig()
	if err != nil {
		panic(err)
	}
	logger := utils.GetLogger(cfg)
	logger.Info("Logger initialized for " + cfg.Dgn + " Environment")
	logger.Info("Server Identity: " + cfg.ServerIdentity)
}
