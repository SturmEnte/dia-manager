package main

import (
	"dia-manager-backend/env"
	"dia-manager-backend/router"
	"dia-manager-backend/scheduler"
)

func main() {
    environmentVars := env.Load()
    env.ConnectPostgres(environmentVars.DatabaseUri)

    scheduler.StartScheduler()

    r := router.SetupRouter(environmentVars)
    r.Run(":" + environmentVars.ServerPort)
}
