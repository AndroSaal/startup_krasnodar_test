package main

import (
	"fmt"

	"github.com/startup_krasnodar_test/src/pkg/config"
)

func main() {
	config := config.MustLoadConfig()
	fmt.Println(config.DBConfig, config.SrvConfig)
}
