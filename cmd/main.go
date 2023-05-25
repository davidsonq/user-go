package main

import (
	"fmt"

	"github.com/davidsonq/user-go/configs"
)

func main() {
	cfg := configs.GetConfig()

	fmt.Print(cfg.APIconfigs.Port)
}
