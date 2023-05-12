-- name: CreateQuestionAnswer :one
INSERT INTO
    "QuestionAnswer" (
        "QuestionId",
        "AnswerText",
        "Score"
    )
VALUES
    (
        sqlc.arg('QuestionId') :: INT,
        sqlc.arg('AnswerText') :: VARCHAR(2000),
        sqlc.arg('Score') :: INT
    ) RETURNING *;

-- name: GetQuestionAnswerById :one
SELECT
    *
FROM
    "QuestionAnswer"
WHERE
    "Id" = sqlc.arg('Id') :: INT
LIMIT
    1;

