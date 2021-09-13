package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mercadolibre/time-zone-front/src/api/model"
	"github.com/mercadolibre/time-zone-front/src/api/server"
	"github.com/mercadolibre/time-zone-front/src/api/util"
)

type InitMain struct {
	Controller *server.Controllers
}

func main() {
	var input string

	controllerMain := InitMain{
		Controller: server.AppendControllers(),
	}
	fmt.Printf("\n### SELECT AN ACTION ###\n")
	fmt.Printf("\nSelect 1) ETS\n")
	fmt.Println("Select 2) Shiping-Route")
	fmt.Printf("Select 3) Set values by console\n")
	fmt.Scanln(&input)
	op, error := validateOption(input)
	if error == nil {

		switch op {
		case model.Input_FILE_Ets:
			ro, er := controllerMain.Controller.TimeZoneController.ReadETS()
			if er == nil {
				controllerMain.runSelect1(ro)
			} else {
				fmt.Println(er)
			}
		case model.Input_FILE_Shiping:

			dat, er := controllerMain.Controller.TimeZoneController.ReadShipments()
			controllerMain.checkSHP(dat, er)
		case model.None:
			var input2 string
			fmt.Printf("\n###Example of allowed actions: ###\n")
			fmt.Printf("\n*Format for ETS command: \"ets 2021-09-12T12:00GMT+1\"\n")
			fmt.Printf("*Format for Shipment command: \"shp -2\"\n")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				input2 = strings.TrimSpace(scanner.Text())
			}

			if len(input2) == 25 || len(input2) == 26 {
				input2 := strings.Replace(input2, "shp ", "", -1)
				dat, e := controllerMain.validateOption3ETS(input2)
				if e != nil {
					fmt.Println(e)
				} else {
					controllerMain.Controller.TimeZoneController.WriteETS(input2)
					controllerMain.runSelect1(dat)
				}
			} else if len(input2) == 6 || len(input2) == 7 {
				controllerMain.Controller.TimeZoneController.WriteSHP(input2)
				dat := new(model.JsonShipment)
				dat.Gmt = input2
				controllerMain.runSelect2(*dat)
			} else {
				fmt.Println("Error malformed command")
			}
		}
	} else {
		fmt.Println(error)
	}

}

func (main *InitMain) runSelect2(newGMT model.JsonShipment) {

	rout, _ := main.Controller.TimeZoneController.ReadRoute()
	main.Controller.TimeZoneController.SetTimeByShipment(rout, strings.Replace(newGMT.Gmt, "shp", "", -1))

}

func (main *InitMain) runSelect1(component model.DataBody) {
	main.Controller.TimeZoneController.GetTimeByETS(component)
}

func (main *InitMain) validateOption3ETS(input string) (model.DataBody, error) {
	option := input[:3]

	if strings.ToLower(option) == model.OptionEts {
		hourSplit, dateSplit, gmt := util.GetSplitDatenadHourETS(strings.Replace(input, "ets ", "", -1))
		dataBody := new(model.DataBody)
		dataBody.SetDataBody(dateSplit[0], dateSplit[1], dateSplit[2], hourSplit[0], hourSplit[1], gmt)
		return *dataBody, nil
	} else {
		return model.DataBody{}, fmt.Errorf("Error malformed data  expected ets command  ")
	}

}
func isEts(value string) error {
	if strings.ToLower(value) == model.OptionShipment {
		return nil
	} else if strings.ToLower(value) == model.OptionEts {
		return nil
	} else {
		return fmt.Errorf("Error malformed data")

	}

}

func (main *InitMain) checkSHP(dat model.JsonShipment, er error) {
	if er == nil {
		main.runSelect2(dat)
	} else {
		fmt.Println(er)
	}

}
func validateOption(typeFile string) (string, error) {
	var nameFile string
	var error error
	if typeFile == "1" {
		nameFile = model.Input_FILE_Ets
	} else if typeFile == "2" {
		nameFile = model.Input_FILE_Shiping
	} else if typeFile == "3" {
		nameFile = model.None
	} else {
		return nameFile, fmt.Errorf("No option Validated")
	}
	return nameFile, error
}
