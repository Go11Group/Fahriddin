package postgres

import (
	"database/sql"
	"my_module/model"
)

type StationRepo struct{
	DB *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo{
	return &StationRepo{DB: db}
}

func (s *StationRepo) Create(station model.Station) error{
	_,err := s.DB.Exec("INSERT INTO Station(name) VALUES($1)",station.Name)
	return err
}

func (s *StationRepo) GetById(id string) (*model.Station,error){
	station := model.Station{}
	err := s.DB.QueryRow("SELECT * FROM Station WHERE = $1",id).Scan(&station.Id,&station.Name)
	if err != nil{
		return nil,err
	}
	return &station,nil
}

// func (s *StationRepo) Update()