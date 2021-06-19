package helpers

import (
	"aquarium-lights/internal/models"

	"time"
)

func Bod(t models.CustomTime) time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}
