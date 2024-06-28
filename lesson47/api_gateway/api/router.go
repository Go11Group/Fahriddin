package api

import (
	"my_module/api/handler"
	t "my_module/genproto/transport"
	w "my_module/genproto/weather"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Router(conn *grpc.ClientConn) *gin.Engine{
	
	transport := t.NewTransportServiceClient(conn)
	weather := w.NewWeatherServiceClient(conn)
	
	handler := handler.NewHandler(weather,transport)
	
	router := gin.Default()

	w :=router.Group("/weather")

	w.GET("/current",handler.GetCurrentWeather)
	w.GET("/forecast",handler.GetWeatherForecast)
	w.GET("/condition",handler.ReportWeatherCondition)

	t := router.Group("/transport")

	t.GET("/schedule",handler.GetBusSchedule)
	t.GET("/location",handler.TrackBusLocation)
	t.GET("/trafficjam",handler.ReportTrafficJam)

	return router
}
