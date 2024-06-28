package service

import (
	"context"
	pb "my_module/genproto/transport"
	"my_module/storage/postgres"
	"time"
)

type TransportService struct{
	pb.UnimplementedTransportServiceServer
	TDB postgres.TransportRepo
}

func NewTransportService(tdb postgres.TransportRepo) *TransportService{
	return &TransportService{TDB: tdb}
}

func (t *TransportService) GetBusSchedule(ctx context.Context,number *pb.Number)(*pb.Buss,error){
	bus,err := t.TDB.GetBusSchedule(number)
	if err != nil{
		return nil,err
	}
	time.Sleep(2*time.Second)
	return bus,nil
}

func (t *TransportService) TrackBusLocation(ctx context.Context, number *pb.Number)(*pb.Location, error){
	location, err := t.TDB.TrackBusLocation(number)
	if err != nil{
		return nil,err
	}
	time.Sleep(2 * time.Second)
	return location, err
}

func (t *TransportService) ReportTrafficJam(ctx context.Context, location *pb.Location)(*pb.Status, error){
	status, err := t.TDB.ReportTrafficJam(location)
	if err != nil{
		return nil,err
	}
	time.Sleep(2 * time.Second)
	return status, err
}
