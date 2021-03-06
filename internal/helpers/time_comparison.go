package helpers

import "time"

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
