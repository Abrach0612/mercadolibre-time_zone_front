package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mercadolibre/time-zone-front/src/api/model"
)

type TimeZonesServiceInterface interface {
	CreateOutputRouterTime(componentETS model.Ets)
	ReadJsonOutRoute() (model.ComponentRouter, error)
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
	check(error)
}

func (this *TimeZonesService) ReadJsonOutRoute() (model.ComponentRouter, error) {
	var jsonRoute model.JsonRoute

	jsonFile, err := os.Open("./resource/" + model.Output_FILE_Route)
	if check(err) {
		return model.ComponentRouter{}, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &jsonRoute)

	componentRouter := new(model.ComponentRouter)
	componentRouter.Date = jsonRoute.RouteTime.DateGMT[0:10]
	componentRouter.Hour = jsonRoute.RouteTime.DateGMT[11:16]
	return *componentRouter, err

}

func check(e error) bool {
	if e != nil {
		fmt.Println(e)
		return true
	}
	return false
}
