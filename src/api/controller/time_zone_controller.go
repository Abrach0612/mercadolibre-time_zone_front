package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mercadolibre/time-zone-front/src/api/model"
	"github.com/mercadolibre/time-zone-front/src/api/util"
)

type timeZonesServiceInterface interface {
	CreateOutputRouterTime(componentETS model.Ets)
	ReadJsonOutRoute() (model.ComponentRouter, error)
}

type TimeZoneController struct {
	TimeZonesServiceInterface timeZonesServiceInterface
}

func (controller *TimeZoneController) GetTimeByETS(component model.DataBody, isDefault bool) {
	dataEts := new(model.Ets)
	dataEts.GetTimeEts(component, isDefault)
	if isDefault {
		controller.NewRunETS(component.TimeGMT)
	} else {
		controller.TimeZonesServiceInterface.CreateOutputRouterTime(*dataEts)
	}
}

func (controller *TimeZoneController) NewRunETS(gmt string) {
	var inputNew string
	fmt.Println("\n## Convertir a Fecha GMT del site.")
	for {
		fmt.Println("Ingresar Fecha GMT del site. Ejm: 2021-09-14 20:00 ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			inputNew = strings.TrimSpace(scanner.Text())
		}
		if len(inputNew) != 16 {
			fmt.Printf("\n## Se envio un mal Formato. Ejm: 2021-09-14 20:00")
			util.PrintPressEnter()
			continue
		} else {
			break
		}
	}
	dataBody := newMapperModelCTP(inputNew, gmt)
	controller.GetTimeByETS(dataBody, false)
}

func newMapperModelCTP(inputNew string, gmt string) model.DataBody {
	hourSplit, dateSplit := getSplitDatenadHourETS(inputNew)
	dataBody := new(model.DataBody)
	dataBody.SetDataBody(dateSplit[0], dateSplit[1], dateSplit[2], hourSplit[0], hourSplit[1], gmt)
	return *dataBody
}

func getSplitDatenadHourETS(inputNew string) ([]string, []string) {
	dateSplit := strings.Split(inputNew[:10], "-")
	hourSplit := strings.Split(inputNew[11:16], ":")
	return hourSplit, dateSplit
}

func (controller *TimeZoneController) SetTimeByShipment(route model.ComponentRouter, gmt string) {
	shipmentRout := model.ShipmetRoute{
		ShipmetRoute: &model.TimeZones{},
	}
	shipmentRout.SetShipmeRoute(route, strings.TrimSpace(gmt))
	fmt.Println("## Paso 2:  Obtener CPT de shipment Route")
	fmt.Println("\nLeyendo la fecha y la Hora desde Route Time:")
	fmt.Printf(route.Date + "\b " + route.Hour)
	util.PrintPressEnter()
	fmt.Println("*** Shipmet Route ***")
	fmt.Printf("Fecha Shipment GMT: %s", shipmentRout.ShipmetRoute.DateGMT)
	util.PrintPressEnter()
	fmt.Printf("Fecha Shipment UTC: %v ", shipmentRout.ShipmetRoute.DateUTC.Format(time.RFC3339))
	util.PrintPressEnter()

}
func (controller *TimeZoneController) ReadRoute() (model.ComponentRouter, error) {
	return controller.TimeZonesServiceInterface.ReadJsonOutRoute()
}
