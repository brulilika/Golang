package banco

import (
	"database/sql"

	/*
		O _ antes da importação serve para mantê-la e permitir sua
		utilização durante o tempo de execução
	*/
	_ "github.com/lib/pq"
)

func DatabaseConnection() *sql.DB {
	connection := "user=postgres dbname=sua_tabela password=sua_senha host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
