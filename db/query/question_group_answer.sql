-- name: CreateQuestionGroupAnswer :one
INSERT INTO
    "QuestionGroupAnswer" (
        "QuestionGroupId",
        "MaxScr",
        "MinScr",
        "AnswerText"
    )
VALUES
    (
        sqlc.arg('QuestionGroupId') :: INT,
        sqlc.arg('MaxScr') :: INT,
        sqlc.arg('MinScr') :: INT,
        sqlc.arg('AnswerText') :: VARCHAR(2000)
    ) RETURNING *;

-- name: GetQuestionGroupAnswerById :one
SELECT
    *
FROM
    "QuestionGroupAnswer"
WHERE
    "Id" = sqlc.arg('Id') :: INT
LIMIT
    1;
