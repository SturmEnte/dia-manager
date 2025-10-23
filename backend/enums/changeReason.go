package enums

type ChangeReason int16

const (
	ChangeReasonUnknown               ChangeReason = 0
	ChangeReasonOther                 ChangeReason = 1
	ChangeReasonRegularChangeInterval ChangeReason = 2
	ChangeReasonInflammation          ChangeReason = 3
	ChangeReasonSlowInsulinReaction   ChangeReason = 4
)