package model

import (
	"strings"
	"time"

	"github.com/mercadolibre/time-zone-front/src/api/util"
)

type EtsInterface interface {
	GetTimeEts(component DataBody)
	SetShipmeRoute(route ComponentRouter, gmt string)
}

type Ets struct {
	DateETS_UTC  time.Time
	RouteTime    *TimeZones
	ShipmetRoute *TimeZones
}

type ComponentRouter struct {
	Date string
	Hour string
}

type ShipmetRoute struct {
	ShipmetRoute *TimeZones
}

type TimeZones struct {
	DateGMT string
	DateUTC time.Time
}

type DataBody struct {
	Year    string
	Month   string
	Day     string
	Hour    string
	HourMin string
	TimeGMT string
}

func (data *DataBody) SetDataBody(year string, month string, day string, hour string, min string, gmt string) {
	data.Year = year
	data.Month = month
	data.Day = day
	data.Hour = hour
	data.HourMin = min
	data.TimeGMT = gmt
}

func (ets *Ets) GetTimeEts(component DataBody) {
	ets.DateETS_UTC = GetTimeUTCDefault(component)
	ets.RouteTime = new(TimeZones)
	ets.RouteTime.SetTimeZone(component)
	ets.ShipmetRoute = new(TimeZones)
	ets.ShipmetRoute.SetTimeZone(component)
}

func (shipment *ShipmetRoute) SetShipmeRoute(route ComponentRouter, gmt string) {
	shipment.ShipmetRoute.DateUTC = GetTimeRoute(route, gmt)
	shipment.ShipmetRoute.DateGMT = strings.Replace(GetTimeRouteDefault(route).Format(time.RFC3339), ":00Z", "GMT"+gmt, -1)
}

func (this *TimeZones) SetTimeZone(component DataBody) {
	this.DateUTC = GetTimeUTCDefault(component)
	this.DateGMT = strings.Replace(GetTimeGMTDefault(component).Format(time.RFC3339), ":00Z", "GMT"+component.TimeGMT, -1)
}

func GetTimeUTCDefault(body DataBody) time.Time {
	t := time.Date(util.GetAtoI(body.Year), time.Month(util.GetAtoI(body.Month)), util.GetAtoI(body.Day), util.GetAtoI(body.Hour), util.GetAtoI(body.HourMin), 0, 0, time.UTC)
	return util.GetTimeGMT(t, body.TimeGMT)
}

func GetTimeGMTDefault(body DataBody) time.Time {
	t := time.Date(util.GetAtoI(body.Year), time.Month(util.GetAtoI(body.Month)), util.GetAtoI(body.Day), util.GetAtoI(body.Hour), util.GetAtoI(body.HourMin), 0, 0, time.UTC)
	return t
}

func GetTimeRoute(component ComponentRouter, gmt string) time.Time {
	dateSplit := strings.Split(component.Date, "-")
	hourSplit := strings.Split(component.Hour, ":")
	t := time.Date(util.GetAtoI(dateSplit[0]), time.Month(util.GetAtoI(dateSplit[1])), util.GetAtoI(dateSplit[2]), util.GetAtoI(hourSplit[0]), util.GetAtoI(hourSplit[1]), 0, 0, time.UTC)
	return util.GetTimeGMT(t, gmt)
}
func GetTimeRouteDefault(component ComponentRouter) time.Time {
	dateSplit := strings.Split(component.Date, "-")
	hourSplit := strings.Split(component.Hour, ":")
	t := time.Date(util.GetAtoI(dateSplit[0]), time.Month(util.GetAtoI(dateSplit[1])), util.GetAtoI(dateSplit[2]), util.GetAtoI(hourSplit[0]), util.GetAtoI(hourSplit[1]), 0, 0, time.UTC)
	return t
}
