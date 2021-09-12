package model

import (
	"time"
)

type TimeZonesInterface interface {
	SetTimeZone(time time.Time, gmt string)
}

type TimeZondes struct {
	TimeGMT time.Time
	DateUTC time.Time
}
