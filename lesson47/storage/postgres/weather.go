package postgres

import (
	"database/sql"
	pb "my_module/genproto/weather"
)

type WeatherRepo struct {
	DB *sql.DB
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{DB: db}
}

func (w *WeatherRepo) GetCurrentWeather(time *pb.Time) (*pb.Weather, error) {
	tr, err := w.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	weather := pb.Weather{}
	err = w.DB.QueryRow("SELECT weather,tempratura,wind,damp FROM WeatherTime WHERE time = $1", time).Scan(
		&weather.Weather, &weather.Tempratura, &weather.Wind, &weather.Damp,
	)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}

func (w *WeatherRepo) GetWeatherForecast(day *pb.Day) (*pb.Weather, error) {
	tr, err := w.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	weather := pb.Weather{}
	err = w.DB.QueryRow("SELECT weather,tempratura,wind,damp FROM WeatherDay WHERE day = $1", day).Scan(
		&weather.Weather, &weather.Tempratura, &weather.Wind, &weather.Damp,
	)
	if err != nil {
		return nil, err
	}
	return &weather, nil
}

func (w *WeatherRepo) ReportWeatherCondition(weather *pb.Weather) (*pb.Status, error) {
	_, err := w.DB.Exec(`INSERT INTO WeatherTime(time, weather, temp, damp, wind) VALUES($1, $2, $3, $4, $5)`,
		weather.Time, weather.Weather, weather.Tempratura, weather.Damp, weather.Wind)
	if err != nil || weather.Time == "" || weather.Weather == "" {
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}
