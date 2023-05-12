package handlers

import (
	"context"
	"database/sql"
	"log"
	"mind-demo-backend/conf"
	"mind-demo-backend/db"

	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	lstd  *log.Logger
	lerr  *log.Logger
	c     *conf.Config
	pgsql *sql.DB
}

var validate = validator.New()

func NewHandlers(lstd *log.Logger, lerr *log.Logger, c *conf.Config, pgsql *sql.DB) *Handlers {
	return &Handlers{lstd: lstd, lerr: lerr, c: c, pgsql: pgsql}
}

func (hd *Handlers) queries(ctx context.Context) (*db.Queries, *sql.DB, error) {
	queries := db.New(hd.pgsql)

	return queries, hd.pgsql, nil
}
