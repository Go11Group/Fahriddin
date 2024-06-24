package handler

import "my_module/storage/postgres"

type Handler struct{
	User *postgres.UsersRepo
}