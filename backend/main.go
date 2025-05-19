package main

import (
	"dia-manager-backend/config"
	"dia-manager-backend/router"
)

func main() {
    cfg := config.Load()

    config.ConnectPostgres(cfg.DatabaseUri)

    r := router.SetupRouter(cfg)

    r.Run(":" + cfg.ServerPort)
}
