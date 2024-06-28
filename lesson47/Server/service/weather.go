package service

import (
	"context"
	"log"
	pb "my_module/genproto/weather"
	"my_module/storage/postgres"
	"time"
)

type WeatherService struct {
	pb.UnimplementedWeatherServiceServer
	WDB postgres.WeatherRepo
}

func NewWeatherService(wdb postgres.WeatherRepo) *WeatherService {
	return &WeatherService{WDB: wdb}
}

func (w *WeatherService) GetCurrentWeather(ctx context.Context, times *pb.Time) (*pb.Weather, error) {
	weather, err := w.WDB.GetCurrentWeather(times)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(2*time.Second)
	return weather, err
}

func (w *WeatherService) GetWeatherForecast(ctx context.Context, day *pb.Day) (*pb.Weather, error) {
	weather, err := w.WDB.GetWeatherForecast(day)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return weather, err
}

func (w *WeatherService) ReportWeatherCondition(ctx context.Context, weather *pb.Weather) (*pb.Status, error) {
	status, err := w.WDB.ReportWeatherCondition(weather)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(2 * time.Second)
	return status, err
}
