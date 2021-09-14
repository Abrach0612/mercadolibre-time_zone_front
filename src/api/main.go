package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mercadolibre/time-zone-front/src/api/model"
	"github.com/mercadolibre/time-zone-front/src/api/server"
	"github.com/mercadolibre/time-zone-front/src/api/util"
)

type InitMain struct {
	Controller *server.Controllers
}

var GMTppl string

func main() {
	controllerMain := InitMain{
		Controller: server.AppendControllers(),
	}
	loc, _ := time.LoadLocation("America/Mexico_City")
	t := time.Now().In(loc)
	fmt.Printf("\nHora actual Mexico DF %s GMT %s\n\n\n", t.Format(time.Kitchen), t.Format(time.RFC3339)[19:22])
	dataBody := mapperModelCTP(t.Format(time.RFC3339))
	controllerMain.runETS(dataBody, true)
}

func (main *InitMain) runETS(component model.DataBody, isDefault bool) {
	main.Controller.TimeZoneController.GetTimeByETS(component, isDefault)
	main.runSHP()
}

func (main *InitMain) runSHP() {
	var valueGMT string
	fmt.Println("## Se Produce Cambio de Time Zone.")
	for {
		fmt.Println("Ingresar GMT. Ejm: +1 / -1:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			valueGMT = strings.TrimSpace(scanner.Text())
		}
		if len(valueGMT) == 2 && strings.Contains(valueGMT, "+") || strings.Contains(valueGMT, "-") {
			rout, _ := main.Controller.TimeZoneController.ReadRoute()
			main.Controller.TimeZoneController.SetTimeByShipment(rout, valueGMT)
			break
		} else {
			fmt.Printf("\n## Se envio un mal Formato. Ejm: +1 / -1")
			util.PrintPressEnter()
			continue
		}
	}
	fmt.Println("\n## NUEVO CICLO...")
	main.Controller.TimeZoneController.NewRunETS(GMTppl)
	main.runSHP()
}

func mapperModelCTP(time string) model.DataBody {
	hourSplit, dateSplit, gmt := util.GetSplitDatenadHourETS(time)
	GMTppl = gmt
	dataBody := new(model.DataBody)
	dataBody.SetDataBody(dateSplit[0], dateSplit[1], dateSplit[2], hourSplit[0], hourSplit[1], gmt)
	return *dataBody
}
