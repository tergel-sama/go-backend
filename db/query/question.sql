-- name: CreateQuestion :one
INSERT INTO
    "Question" (
        "QuestionGroupId",
        "Question"
    )
VALUES
    (
        sqlc.arg('QuestionGroupId') :: INT,
        sqlc.arg('Question') :: VARCHAR(2000)
    ) RETURNING *;

-- name: GetQuestionById :one
SELECT
    *
FROM
    "Question"
WHERE
    "Id" = sqlc.arg('Id') :: INT
LIMIT
    1;


-- name: GetQuestionByTestId :many
SELECT 
  q.* 
FROM 
  "Test" t 
  INNER JOIN "QuestionGroup" qg ON t."Id" = qg."TestId" 
  INNER JOIN "Question" q ON q."QuestionGroupId" = qg."Id" 
WHERE 
  t."Id" = sqlc.arg('TestId') :: INT;
