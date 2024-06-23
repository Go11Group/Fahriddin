package model

type FilterStation struct{
	Name string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type FilterCard struct{
	Number string	`json:"number"`
	UserId string	`json:"user_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type FilterTerminal struct{
	StationId string `json:"station_id"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type FilterTransaction struct{
	CardId string	`json:"card_id"`
	Amount float64	`json:"amount"`
	TerminalId string	`json:"terminal_id"`
	TransactionType string	`json:"transaction_type"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}