package controller

import (
	"fmt"
	"strings"

	"github.com/mercadolibre/time-zone-front/src/api/model"
)

type timeZonesServiceInterface interface {
	CreateOutputRouterTime(componentETS model.Ets)
	ReadJsonInputETS() (model.DataBody, error)
	ReadJsonOutRoute() (model.ComponentRouter, error)
	ReadJsonInputShipment() (model.JsonShipment, error)
	WriteJsonInputETS(string)
	WriteJsonInputShipment(string)
}

type TimeZoneController struct {
	TimeZonesServiceInterface timeZonesServiceInterface
}

func (controller *TimeZoneController) GetTimeByETS(component model.DataBody) {
	dataEts := new(model.Ets)
	dataEts.GetTimeEts(component)
	controller.TimeZonesServiceInterface.CreateOutputRouterTime(*dataEts)
}

func (controller *TimeZoneController) SetTimeByShipment(route model.ComponentRouter, gmt string) {
	shipmentRout := model.ShipmetRoute{
		ShipmetRoute: &model.TimeZones{},
	}
	fmt.Println("Shipment GMT:", strings.TrimSpace(gmt))
	shipmentRout.SetShipmeRoute(route, strings.TrimSpace(gmt))
	fmt.Println("########### Point 2(Shipment-Route)  ##################")
	fmt.Println("Date Shipment GMT:", shipmentRout.ShipmetRoute.DateGMT)
	fmt.Println("Date Shipment UTC:", shipmentRout.ShipmetRoute.DateUTC)
	fmt.Println("#############################")

}

func (controller *TimeZoneController) WriteETS(data string) {
	controller.TimeZonesServiceInterface.WriteJsonInputETS(data)
}
func (controller *TimeZoneController) WriteSHP(data string) {
	controller.TimeZonesServiceInterface.WriteJsonInputShipment(data)
}

func (controller *TimeZoneController) ReadETS() (model.DataBody, error) {
	return controller.TimeZonesServiceInterface.ReadJsonInputETS()
}
func (controller *TimeZoneController) ReadRoute() (model.ComponentRouter, error) {
	return controller.TimeZonesServiceInterface.ReadJsonOutRoute()
}
func (controller *TimeZoneController) ReadShipments() (model.JsonShipment, error) {
	return controller.TimeZonesServiceInterface.ReadJsonInputShipment()
}
