package enums

type ChangeReason int16

const (
	ChangeReasonUnknown ChangeReason = 0
	ChangeReasonCreated ChangeReason = 1
	ChangeReasonUpdated ChangeReason = 2
	ChangeReasonDeleted ChangeReason = 3
)