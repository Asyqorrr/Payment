package services

import (
	db "payment/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserService struct {
	*db.Queries
}

func NewUserService(dbConn *pgxpool.Conn) *UserService{
	return &UserService{
		Queries: db.New(dbConn),
	}
}