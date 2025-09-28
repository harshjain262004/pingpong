package main

import "github.com/harshjain262004/pingpong/Backend/utils"

func main() {
	cfg, err := utils.GetConfig()
	if err != nil {
		panic(err)
	}
	println("Server Identity:", cfg.ServerIdentity)

}
