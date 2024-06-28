package postgres

import (
	"database/sql"
	pb "my_module/genproto/transport"
)

type TransportRepo struct {
	DB *sql.DB
}

func NewTransportRepo(db *sql.DB) *TransportRepo {
	return &TransportRepo{DB: db}
}

func (t *TransportRepo) GetBusSchedule(number *pb.Number)(*pb.Buss,error){
	tr, err := t.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	bus := pb.Buss{}
	err = t.DB.QueryRow("SELECT * FROM Bus where number = $1",number).Scan(
		&bus.Number,&bus.From,&bus.To,&bus.Loc,&bus.TrafficStat,
	)
	if err != nil{
		return nil,err
	}
	return &bus,nil
}

func (t *TransportRepo)TrackBusLocation(number *pb.Number)(*pb.Location,error){
	tr, err := t.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	loc := pb.Location{}
	err = t.DB.QueryRow(`SELECT location FROM Bus Where number = $1`, number).Scan(
		&loc)
	return &loc, err
}

func(t *TransportRepo) ReportTrafficJam(location *pb.Location)(*pb.Status,error){
	_, err := t.DB.Exec(`UPDATE SET Bus traffic_stat = true WHERE location = $1`, location)
	if err != nil{
		return &pb.Status{Status: false}, err
	}
	return &pb.Status{Status: true}, nil
}