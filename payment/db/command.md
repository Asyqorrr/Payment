migrate create -ext sql -dir db/migrations -seq init_schema


migrate create -ext sql -dir db/migrations -seq add_column_product

migrate -path db/migrations -database "postgresql://postgres:1234@localhost:5432/fintech?sslmode=disable" -verbose up

sqlc init
sqlc generate