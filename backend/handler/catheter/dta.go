package catheter

import "time"

type CreateCatheterRequest struct {
	Start	time.Time	`json:"startedAt" binding:"required"`
	End		*time.Time	`json:"endedAt"`
}
