// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: test.sql

package db

import (
	"context"
)

const createTest = `-- name: CreateTest :one
INSERT INTO
    "Test" (
        "Name",
        "Desc",
        "Img",
        "Minute",
        "AgeCls",
        "BeforeDesc",
        "AfterDesc",
        "ExampleReport",
        "IsActive"
    )
VALUES
    (
        $1 :: VARCHAR(200),
        $2 :: VARCHAR(2000),
        $3 :: VARCHAR(200),
        $4 :: INT,
        $5 :: VARCHAR(10),
        $6 :: VARCHAR(2000),
        $7 :: VARCHAR(2000),
        $8 :: VARCHAR(200),
        $9 :: BOOLEAN
    ) RETURNING "Id", "Name", "Desc", "Img", "Minute", "AgeCls", "BeforeDesc", "AfterDesc", "ExampleReport", "IsActive"
`

type CreateTestParams struct {
	Name          string
	Desc          string
	Img           string
	Minute        int32
	AgeCls        string
	BeforeDesc    string
	AfterDesc     string
	ExampleReport string
	IsActive      bool
}

func (q *Queries) CreateTest(ctx context.Context, arg CreateTestParams) (Test, error) {
	row := q.db.QueryRowContext(ctx, createTest,
		arg.Name,
		arg.Desc,
		arg.Img,
		arg.Minute,
		arg.AgeCls,
		arg.BeforeDesc,
		arg.AfterDesc,
		arg.ExampleReport,
		arg.IsActive,
	)
	var i Test
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Desc,
		&i.Img,
		&i.Minute,
		&i.AgeCls,
		&i.BeforeDesc,
		&i.AfterDesc,
		&i.ExampleReport,
		&i.IsActive,
	)
	return i, err
}

const getTestById = `-- name: GetTestById :one
SELECT
    "Id", "Name", "Desc", "Img", "Minute", "AgeCls", "BeforeDesc", "AfterDesc", "ExampleReport", "IsActive"
FROM
    "Test"
WHERE
    "Id" = $1 :: INT
LIMIT
    1
`

func (q *Queries) GetTestById(ctx context.Context, id int32) (Test, error) {
	row := q.db.QueryRowContext(ctx, getTestById, id)
	var i Test
	err := row.Scan(
		&i.Id,
		&i.Name,
		&i.Desc,
		&i.Img,
		&i.Minute,
		&i.AgeCls,
		&i.BeforeDesc,
		&i.AfterDesc,
		&i.ExampleReport,
		&i.IsActive,
	)
	return i, err
}
