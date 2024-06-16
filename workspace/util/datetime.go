package util

import (
	"math/rand"
	"time"
)

func RandomDateTime() time.Time {
	year := rand.Intn(2024-1980) + 1980
	month := rand.Intn(12-1) + 1
	day := rand.Intn(30-1) + 1
	if month == 2 && day > 28 {
		month = 28
	}
	hh := rand.Intn(23)
	mm := rand.Intn(59)
	ss := rand.Intn(59)
	return time.Date(year, time.Month(month), day, hh, mm, ss, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
}
