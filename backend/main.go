package main

import (
	"dia-manager-backend/config"
	"dia-manager-backend/router"
	"dia-manager-backend/scheduler"
)

func main() {
    cfg := config.Load()
    config.ConnectPostgres(cfg.DatabaseUri)

    scheduler.StartScheduler()

    r := router.SetupRouter(cfg)
    r.Run(":" + cfg.ServerPort)
}
