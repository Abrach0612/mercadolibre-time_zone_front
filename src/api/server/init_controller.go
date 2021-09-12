package server

import (
	"github.com/mercadolibre/time-zone-front/src/api/controller"
	"github.com/mercadolibre/time-zone-front/src/api/service"
)

type Controllers struct {
	TimeZoneController controller.TimeZoneController
}

func AppendControllers() *Controllers {
	tzController := newTZController()
	return &Controllers{
		TimeZoneController: *tzController,
	}
}

func newTZController() *controller.TimeZoneController {
	return &controller.TimeZoneController{
		TimeZonesServiceInterface: newtimeZonesServiceInterface(),
	}
}
func newtimeZonesServiceInterface() *service.TimeZonesService {
	return &service.TimeZonesService{}
}
