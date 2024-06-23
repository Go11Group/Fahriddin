package model

type FilterStation struct{
	Name string `json:"name"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}