package postgres

import (
	pb "my_module/genproto/weather"
	"database/sql"
)

type WeatherConditionsRepo struct{
	DB *sql.DB
}

func NewWeatherConditionsRepo(db *sql.DB)*WeatherConditionsRepo{
	return &WeatherConditionsRepo{DB: db}
}

func (w *WeatherConditionsRepo) Create(pb.){

}