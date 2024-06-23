package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/model"
	"strings"
)

type StationRepo struct {
	DB *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{DB: db}
}

func (s *StationRepo) Create(station model.Station) error {
	tr, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	_, err = s.DB.Exec("INSERT INTO Station(name) VALUES($1)", station.Name)
	return err
}

func (s *StationRepo) GetById(id string) (model.Station, error) {
	tr, err := s.DB.Begin()
	if err != nil {
		return model.Station{}, err
	}
	defer tr.Commit()

	station := model.Station{}
	err = s.DB.QueryRow("SELECT * FROM Station WHERE id = $1", id).Scan(&station.Id, &station.Name)
	if err != nil {
		return model.Station{}, err
	}
	return station, nil
}

type UpdateStation struct {
	Id   *string `json:"id"`
	Name *string `json:"name"`
}

func (s *StationRepo) Update(updateFilter UpdateStation) error {
	tr, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT id
	  FROM Station
	  WHERE id = $1
	`
	err = s.DB.QueryRow(query, *updateFilter.Id).Err()

	if err != nil {
		return err
	}

	query = `
	  UPDATE Station SET 
	`

	if updateFilter.Name != nil {
		params = append(params, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, *updateFilter.Name)
	}

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, *updateFilter.Id)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d ", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = s.DB.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

func (s *StationRepo) Delete(id string) error {
	tr, err := s.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	_, err = s.DB.Exec("DELETE FROM Station WHERE id = $1", id)

	return err
}

func (s *StationRepo) GetAll(fStation model.FilterStation) ([]model.Station, error) {
	tr, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT id,name FROM Station WHERE true "

	if fStation.Name != "" {
		params["name"] = fStation.Name
		filter += "AND name = :name "
	}

	if fStation.Limit > 0 {
		params["limit"] = fStation.Limit
		filter += " limit :limit "
	}

	if fStation.Offset > 0 {
		params["offset"] = fStation.Offset
		filter += " offset :offset "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var stations []model.Station
	for rows.Next() {
		var station model.Station

		err = rows.Scan(&station.Id, &station.Name)

		if err != nil {
			return nil, err
		}

		stations = append(stations, station)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stations, err
}
