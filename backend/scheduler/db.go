package scheduler

import (
	"context"
	"dia-manager-backend/config"
	"log"

	"github.com/robfig/cron/v3"
)

func StartScheduler() {
    c := cron.New()

    _, err := c.AddFunc("@every 1h", func() {
        log.Println("Running cron job...")
        runSQLTask()
    })

    if err != nil {
        log.Fatalln("Error adding cron job: " + err.Error())
    }

    c.Start()
}

// Deletes all expired tokens because they are invalid no matter if they are in the database or not
func runSQLTask() {
    _, err := config.DB.Exec(context.Background(), `DELETE FROM invalid_tokens WHERE expires < NOW()`)

    if err != nil {
        log.Println("SQL task failed: " + err.Error())
    } else {
        log.Println("SQL task executed successfully. Deleted expired sessions")
    }
}
