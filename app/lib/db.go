package lib

import (
	"crypto/tls"
	"fmt"

	"github.com/go-pg/pg"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	fmt.Println(q.FormattedQuery())
}

// InitDB ...
func InitDB(user, pass, datab string) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:      user,
		Password:  pass,
		Database:  datab,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	db.AddQueryHook(dbLogger{})

	return db
}
