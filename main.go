package main

import (
	"fmt"

	"github.com/iTsLhaj/gator/internal/config"
)

func main() {

	gatorConfig, _ := config.Read()
	gatorConfig.SetUser("Kenzo")

	gatorConfig, _ = config.Read()
	fmt.Printf("Config: %+v\n\t- Database URL: %s\n\t- Username: %s\n",
		gatorConfig,
		gatorConfig.DbUrl,
		gatorConfig.Username,
	)

}
