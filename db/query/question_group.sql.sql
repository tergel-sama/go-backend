-- name: CreateQuestionGroup :one
INSERT INTO
    "QuestionGroup" (
        "TestId",
        "Name",
        "Order"
    )
VALUES
    (
        sqlc.arg('TestId') :: INT,
        sqlc.arg('Name') :: VARCHAR(2000),
        sqlc.arg('Order') :: INT
    ) RETURNING *;

-- name: GetQuestionGroupById :one
SELECT
    *
FROM
    "QuestionGroup"
WHERE
    "Id" = sqlc.arg('Id') :: INT
LIMIT
    1;
