package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mercadolibre/time-zone-front/src/api/model"
	"github.com/mercadolibre/time-zone-front/src/api/util"
)

type TimeZonesServiceInterface interface {
	CreateOutputRouterTime(componentETS model.Ets)
	ReadJsonOutRoute() (model.ComponentRouter, error)
	ReadJsonInputETS() (model.Component, error)
	ReadJsonInputShipment() (model.JsonShipment, error)
	WriteJsonInputETS(string)
	WriteJsonInputShipment(string)
}
type dataTypeBody struct {
	model.DataBody
}
type TimeZonesService struct {
	TimeZonesServiceInterface TimeZonesServiceInterface
}
type JsonETS struct {
	Date string
}
type JsonSHP struct {
	Gmt string
}

func (this *TimeZonesService) CreateOutputRouterTime(componentETS model.Ets) {
	b, err := json.MarshalIndent(componentETS, "", " ")
	check(err)
	error := ioutil.WriteFile("./resource/"+model.Output_FILE_Route, b, 0644)
	fmt.Printf("\n\nMarshal output Point 1(ETS) %s\n", string(b))
	check(error)
}

func (this *TimeZonesService) ReadJsonInputETS() (model.DataBody, error) {
	var datJson model.Component

	jsonFile, err := os.Open("./resource/" + model.Input_FILE_Ets)
	if check(err) {
		return model.DataBody{}, err
	}

	fmt.Println("Successfully Opened Json:")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &datJson)
	hourSplit, dateSplit, gmt := util.GetSplitDatenadHourETS(datJson.Date)
	dataBody := new(dataTypeBody)
	dataBody.SetDataBody(dateSplit[0], dateSplit[1], dateSplit[2], hourSplit[0], hourSplit[1], gmt)
	return dataBody.DataBody, err
}

func (this *TimeZonesService) WriteJsonInputETS(data string) {
	body := new(JsonETS)
	body.Date = strings.Replace(data, "ets ", "", -1)
	writeJsonGeneral(body, model.Input_FILE_Ets)

}
func (this *TimeZonesService) WriteJsonInputShipment(data string) {
	body := new(JsonSHP)
	body.Gmt = data
	writeJsonGeneral(body, model.Input_FILE_Shiping)
}

func writeJsonGeneral(bodyData interface{}, nameFile string) {
	b, err := json.MarshalIndent(bodyData, "", " ")
	check(err)
	error := ioutil.WriteFile("./resource/"+nameFile, b, 0644)
	fmt.Printf("\n\nMarshal output %s\n", string(b))
	check(error)
}
func (this *TimeZonesService) ReadJsonOutRoute() (model.ComponentRouter, error) {
	var jsonRoute model.JsonRoute

	jsonFile, err := os.Open("./resource/" + model.Output_FILE_Route)
	if check(err) {
		return model.ComponentRouter{}, err
	}

	fmt.Println("Successfully Opened Json Route:")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jsonRoute)

	componentRouter := new(model.ComponentRouter)
	componentRouter.Date = jsonRoute.RouteTime.DateGMT[0:10]
	componentRouter.Hour = jsonRoute.RouteTime.DateGMT[11:16]
	fmt.Println("Date Router: " + componentRouter.Date)
	fmt.Println("Hour Router: " + componentRouter.Hour)

	return *componentRouter, err

}

func (this *TimeZonesService) ReadJsonInputShipment() (model.JsonShipment, error) {
	var gmt model.JsonShipment

	jsonFile, err := os.Open("./resource/" + model.Input_FILE_Shiping)
	if check(err) {
		return model.JsonShipment{}, err
	}

	fmt.Println("Successfully Opened Json Shipment")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &gmt)

	return gmt, err

}

func check(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	}
	return false
}
