package services

import (
	db "payment/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EntityService struct {
	*db.Queries
}

func NewEntityService(dbConn *pgxpool.Conn) *EntityService{
	return &EntityService{
		Queries: db.New(dbConn),
	}
}