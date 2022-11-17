package Repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	//host     = "database" //используем в случае подключения к базе данных из контейнера
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "4650"
	dbname   = "apteca"
	sslmode  = "disable"
)

var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

var Connection *sql.DB

func OpenTable() error {
	var err error
	Connection, err = sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	return nil
}
