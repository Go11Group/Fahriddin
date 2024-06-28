package main

import (
	"fmt"
	t "my_module/genproto/transport"
	w "my_module/genproto/weather"
	"my_module/service"
	"my_module/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		panic(err)
	}

	serviceTransport := service.NewTransportService(postgres.TransportRepo{DB: db})
	serviceWeather := service.NewWeatherService(postgres.WeatherRepo{DB: db})

	server := grpc.NewServer()

	t.RegisterTransportServiceServer(server, serviceTransport)
	w.RegisterWeatherServiceServer(server, serviceWeather)

	err = server.Serve(listener)

	if err != nil {
		panic(err)
	}

	fmt.Printf("server listening at %v", listener.Addr())
}
