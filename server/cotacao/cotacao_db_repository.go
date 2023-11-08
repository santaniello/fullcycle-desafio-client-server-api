package cotacao

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const timeoutDB = 10 * time.Millisecond

type CotacaoRepository struct {
	db *sql.DB
}

func NewCotacaoRepository(databasePath string) (*CotacaoRepository, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}
	return &CotacaoRepository{db: db}, nil
}

func (cr *CotacaoRepository) Save(bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDB)
	defer cancel()
	cr.createTableCotacaoIfNotExists()
	statement, err := cr.db.PrepareContext(ctx, "INSERT INTO cotacao (bid) VALUES (?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = statement.Exec(bid)
	return err
}

func (cr *CotacaoRepository) createTableCotacaoIfNotExists() error {
	sqlStmt := "CREATE TABLE IF NOT EXISTS cotacao (id INTEGER PRIMARY KEY AUTOINCREMENT, bid TEXT NOT NULL)"
	_, err := cr.db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return err
	}
	return nil
}

func (cr *CotacaoRepository) Close() error {
	return cr.db.Close()
}
