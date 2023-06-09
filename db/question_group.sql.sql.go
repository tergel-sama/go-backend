// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: question_group.sql.sql

package db

import (
	"context"
)

const createQuestionGroup = `-- name: CreateQuestionGroup :one
INSERT INTO
    "QuestionGroup" (
        "TestId",
        "Name",
        "Order"
    )
VALUES
    (
        $1 :: INT,
        $2 :: VARCHAR(2000),
        $3 :: INT
    ) RETURNING "Id", "TestId", "Name", "Order"
`

type CreateQuestionGroupParams struct {
	TestId int32
	Name   string
	Order  int32
}

func (q *Queries) CreateQuestionGroup(ctx context.Context, arg CreateQuestionGroupParams) (QuestionGroup, error) {
	row := q.db.QueryRowContext(ctx, createQuestionGroup, arg.TestId, arg.Name, arg.Order)
	var i QuestionGroup
	err := row.Scan(
		&i.Id,
		&i.TestId,
		&i.Name,
		&i.Order,
	)
	return i, err
}

const getQuestionGroupById = `-- name: GetQuestionGroupById :one
SELECT
    "Id", "TestId", "Name", "Order"
FROM
    "QuestionGroup"
WHERE
    "Id" = $1 :: INT
LIMIT
    1
`

func (q *Queries) GetQuestionGroupById(ctx context.Context, id int32) (QuestionGroup, error) {
	row := q.db.QueryRowContext(ctx, getQuestionGroupById, id)
	var i QuestionGroup
	err := row.Scan(
		&i.Id,
		&i.TestId,
		&i.Name,
		&i.Order,
	)
	return i, err
}
