package postgres

import (
	"database/sql"
	"fmt"
	pkg "my_module/ReplaceQueryParams"
	"my_module/model"
	"strings"
)

type TransactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{DB: db}
}

func (t *TransactionRepo) Create(transaction model.Transaction) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = t.DB.Exec(`INSER INTO Transaction(card_id,amount,terminal_id,transaction_type) 
	VALUES($1,$2,$3,$4)`, transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType)
	return err
}

func (t *TransactionRepo) GetById(id string) (model.Transaction, error) {
	tr, err := t.DB.Begin()
	if err != nil {
		return model.Transaction{}, err
	}
	defer tr.Commit()

	transaction := model.Transaction{}
	err = t.DB.QueryRow("SELECT * FROM Transaction WHERE id = $1", id).Scan(&transaction.TerminalId, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)

	return transaction, err
}

type UpdateTransaction struct {
	Id              *string  `json:"id"`
	CardId          *string  `json:"card_id"`
	Amount          *float64 `json:"amount"`
	TerminalId      *string  `json:"terminal_id"`
	TransactionType *string  `json:"transaction_type"`
}

func (t *TransactionRepo) Update(updateFilter UpdateTransaction) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	var params []string
	var args []interface{}

	query := `
	  SELECT id
	  FROM Transaction
	  WHERE id = $1
	`
	err = t.DB.QueryRow(query, *updateFilter.Id).Err()

	if err != nil {
		return err
	}

	query = `
	  UPDATE Transaction SET 
	`

	if updateFilter.CardId != nil {
		params = append(params, fmt.Sprintf("card_id = $%d", len(args)+1))
		args = append(args, *updateFilter.CardId)
	}

	if updateFilter.Amount != nil {
		params = append(params, fmt.Sprintf("amount = $%d", len(args)+1))
		args = append(args, *updateFilter.Amount)
	}

	if updateFilter.TerminalId != nil {
		params = append(params, fmt.Sprintf("terminal_id = $%d", len(args)+1))
		args = append(args, *updateFilter.TerminalId)
	}

	if updateFilter.TransactionType != nil {
		params = append(params, fmt.Sprintf("transaction_type = $%d", len(args)+1))
		args = append(args, *updateFilter.TransactionType)
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

func (t *TransactionRepo) Delete(id string) error {
	tr, err := t.DB.Begin()
	if err != nil {
		return err
	}
	defer tr.Commit()

	_, err = t.DB.Exec("DELETE FROM Transaction WHERE id = $1", id)
	return err
}

func (t *TransactionRepo) GetAll(fTransaction model.FilterTransaction) ([]model.Transaction, error) {
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

	query := "SELECT id,card_id,amount,terminal_id,transaction_type FROM Transaction WHERE true "

	if fTransaction.CardId != "" {
		params["card_id"] = fTransaction.CardId
		filter += "AND card_id = :card_id "
	}

	if fTransaction.Amount != 0 {
		params["amount"] = fTransaction.Amount
		filter += "AND amount = :amount"
	}

	if fTransaction.TerminalId != "" {
		params["terminal_id"] = fTransaction.TerminalId
		filter += "AND terminal_id = :terminal_id"
	}

	if fTransaction.TransactionType != "" {
		params["transaction_type"] = fTransaction.TransactionType
		filter += "AND transaction_type = :transaction_type"
	}


	if fTransaction.Limit > 0 {
		params["limit"] = fTransaction.Limit
		filter += " limit :limit "
	}

	if fTransaction.Offset > 0 {
		params["offset"] = fTransaction.Offset
		filter += " offset :offset "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := t.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var transactions []model.Transaction
	for rows.Next() {
		var transaction model.Transaction

		err = rows.Scan(&transaction.Id,&transaction.CardId,&transaction.Amount,&transaction.TerminalId,&transaction.TransactionType)

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, err
}
