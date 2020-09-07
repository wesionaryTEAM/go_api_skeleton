package main

import (
	"teamplace-api/infrastructure"
	"teamplace-api/utils"
)

func main() {
	utils.LoadEnv()

	db := infrastructure.SetupModels()

	infrastructure.SetupRoutes(db)
}
