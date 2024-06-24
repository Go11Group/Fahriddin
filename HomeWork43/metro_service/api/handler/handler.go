package handler

import "my_module/storage/postgres"

type Handler struct {
	Card        *postgres.CardRepo
	Station     *postgres.StationRepo
	Terminal    *postgres.TerminalRepo
	Transaction *postgres.TransactionRepo
}
