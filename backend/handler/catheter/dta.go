package catheter

import (
	"dia-manager-backend/enums"
	"time"
)

type CreateCatheterRequest struct {
	Start			time.Time			`json:"startedAt" binding:"required"`
	End				*time.Time			`json:"endedAt"`
	ChangeReason	*enums.ChangeReason	`json:"changeReason"`
}

type UpdateCatheterRequest struct {
	Start			*time.Time			`json:"startedAt"`
	End				*time.Time			`json:"endedAt"`
	ChangeReason	*enums.ChangeReason	`json:"changeReason"`
}