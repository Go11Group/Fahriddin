package handler

import (
	w "my_module/genproto/weather"
	t "my_module/genproto/transport"
)
type Handler struct{
	Weather w.WeatherServiceClient
	Transport t.TransportServiceClient
}

func NewHandler(weather w.WeatherServiceClient, transport t.TransportServiceClient)*Handler{
	return &Handler{Weather: weather,Transport: transport}
}
