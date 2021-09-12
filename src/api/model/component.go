package model

type Component struct {
	Date string `json:"date"`
}

type JsonRoute struct {
	RouteTime struct {
		DateGMT string `json:"DateGMT"`
	}
}

type JsonShipment struct {
	Gmt string `json:"Gmt"`
}
