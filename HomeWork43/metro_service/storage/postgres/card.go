package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/model"
	"strings"
)

type CardRepo struct {
	DB *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{DB: db}
}

func (c *CardRepo) Create(card model.Card) error {
	tr, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()
	
	_, err = c.DB.Exec("INSER INTO Card(number,user_id) VALUES($1,$2)", card.Number, card.UserId)
	return err
}

func (c *CardRepo) GetById(id string) (model.Card, error) {
	tr, err := c.DB.Begin()
	if err != nil {
		return model.Card{}, err
	}
	defer tr.Commit()

	card := model.Card{}
	err = c.DB.QueryRow("SELECT * FROM Card WHERE id = $1", id).Scan(&card.Id, &card.Number, &card.UserId)

	return card, err
}

type UpdateCard struct {
	Id     *string `json:"id"`
	Number *string `json:"number"`
	UserId *string `json:"user_id"`
}

func (c *CardRepo) Update(updateFilter UpdateCard) error {
	tr, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT id
	  FROM Card
	  WHERE id = $1
	`
	err = c.DB.QueryRow(query, *updateFilter.Id).Err()

	if err != nil {
		return err
	}

	query = `
	  UPDATE Card SET 
	`

	if updateFilter.Number != nil {
		params = append(params, fmt.Sprintf("number = $%d", len(args)+1))
		args = append(args, *updateFilter.Number)
	}

	if updateFilter.UserId != nil{
		params = append(params, fmt.Sprintf("user_id = $%d",len(args)+1))
		args = append(args, *updateFilter.UserId)
	}

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, *updateFilter.Id)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d ", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err = c.DB.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

func (c *CardRepo) Delete(id string)error{
	tr, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_,err = c.DB.Exec("DELETE FROM Card WHERE id = $1",id)
	return err
}

func (c *CardRepo) GetAll(fCard model.FilterCard) ([]model.Card,error){
	tr, err := c.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT id,number,user_id FROM Card WHERE true "

	if fCard.Number != "" {
		params["number"] = fCard.Number
		filter += "AND number = :number "
	}
	
	if fCard.UserId != "" {
		params["user_id"] = fCard.UserId
		filter += "AND user_id = :user_id"
	}

	if fCard.Limit > 0 {
		params["limit"] = fCard.Limit
		filter += " limit :limit "
	}

	if fCard.Offset > 0 {
		params["offset"] = fCard.Offset
		filter += " offset :offset "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := c.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var cards []model.Card
	for rows.Next() {
		var card model.Card

		err = rows.Scan(&card.Id,&card.Number,&card.UserId)

		if err != nil {
			return nil, err
		}

		cards = append(cards, card)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cards, err
}
