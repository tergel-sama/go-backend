package handlers

import (
	"context"
	"database/sql"
	"go-backend/conf"
	"go-backend/db"

	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	c     *conf.Config
	pgsql *sql.DB
}

var validate = validator.New()

func NewHandlers(c *conf.Config, pgsql *sql.DB) *Handlers {
	return &Handlers{c: c, pgsql: pgsql}
}

func (hd *Handlers) queries(ctx context.Context) (*db.Queries, *sql.DB, error) {
	queries := db.New(hd.pgsql)

	return queries, hd.pgsql, nil
}
