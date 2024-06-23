package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/model"
	"strings"
)

type TerminalRepo struct {
	DB *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{DB: db}
}

func (t *TerminalRepo) Create(terminal model.Terminal) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = t.DB.Exec("INSERT INTO Terminal(station_id) VALUES($1)", terminal.StationId)
	return err
}

func (t *TerminalRepo) GetById(id string) (model.Terminal, error) {
	tr, err := t.DB.Begin()
	if err != nil {
		return model.Terminal{}, err
	}
	defer tr.Commit()
	terminal := model.Terminal{}

	err = t.DB.QueryRow("SELECT * FROM Terminal WHERE id = $1", id).Scan(&terminal.Id, &terminal.StationId)

	return terminal, err
}

type UpdateTerminal struct {
	Id        *string `json:"id"`
	StationId *string `json:"station_id"`
}

func (t *TerminalRepo) Update(updateFilter UpdateTerminal) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT id
	  FROM Terminal
	  WHERE id = $1
	`
	err = t.DB.QueryRow(query, *updateFilter.Id).Err()

	if err != nil {
		return err
	}

	query = `
	  UPDATE Terminal SET 
	`

	if updateFilter.StationId != nil {
		params = append(params, fmt.Sprintf("station_id = $%d", len(args)+1))
		args = append(args, *updateFilter.StationId)
	}

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, *updateFilter.Id)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d ", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = t.DB.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

func (t *TerminalRepo) Delete(id string) error{
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_,err = t.DB.Exec("DELETE FROM Terminal WHERE id = $1",id)
	return err
}

func (t * TerminalRepo) GetAll(fTerminal model.FilterTerminal) ([]model.Terminal,error){
	tr, err := t.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT id,station_id FROM Terminal WHERE true "

	if fTerminal.StationId != "" {
		params["station_id"] = fTerminal.StationId
		filter += "AND station_id = :station_id "
	}

	if fTerminal.Limit > 0 {
		params["limit"] = fTerminal.Limit
		filter += " limit :limit "
	}

	if fTerminal.Offset > 0 {
		params["offset"] = fTerminal.Offset
		filter += " offset :offset "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := t.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var terminals []model.Terminal
	for rows.Next() {
		var terminal model.Terminal

		err = rows.Scan(&terminal.Id,&terminal.StationId)

		if err != nil {
			return nil, err
		}

		terminals = append(terminals, terminal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return terminals, err
}