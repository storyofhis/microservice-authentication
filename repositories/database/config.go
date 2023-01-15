package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/storyofhis/auth-service/repositories/model"
)

type DBTransaction struct {
	postgres *sql.Tx
	Builder  squirrel.StatementBuilderType
	ctx      context.Context
}

var config *model.Configuration

// getter function
func Get() *model.Configuration {
	return config
}

func Load() {
	var pathConfig string = "./config.json"
	dataPath, err := ioutil.ReadFile(pathConfig)
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(dataPath, &config); err != nil {
		log.Fatal(err)
	}
}

func OpenConnection(ctx context.Context, readOnly bool) (*DBTransaction, error) {
	var (
		t   = &DBTransaction{}
		db  *sql.DB
		err error
	)
	db, err = sql.Open(Get().Database.Driver, Get().Database.Url)
	if err != nil {
		return t, err
	}
	defer db.Close()

	transaction, err := db.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  readOnly,
	})

	if err != nil {
		return nil, err
	}

	t.ctx = ctx
	t.postgres = transaction
	t.Builder = squirrel.StatementBuilder.
		PlaceholderFormat(squirrel.Dollar).
		RunWith(t.postgres)

	return t, nil
}

func (t *DBTransaction) Commit() (err error) {
	return t.postgres.Commit()
}

func (t *DBTransaction) Rollback() {
	_ = t.postgres.Rollback()
}
