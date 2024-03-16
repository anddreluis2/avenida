package main

import (
	"fmt"

	"github.com/anddreluis2/avenida/config"
	"github.com/anddreluis2/avenida/router"
)

func main() {
	//Initialize configs

	err := config.InitDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
