package model

type FilterUser struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Age    int    `json:"Age"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
