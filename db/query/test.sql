-- name: GetTestById :one
SELECT
    *
FROM
    "Test"
WHERE
    "Id" = sqlc.arg('Id') :: INT
LIMIT
    1;

-- name: CreateTest :one
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
        sqlc.arg('Name') :: VARCHAR(200),
        sqlc.arg('Desc') :: VARCHAR(2000),
        sqlc.arg('Img') :: VARCHAR(200),
        sqlc.arg('Minute') :: INT,
        sqlc.arg('AgeCls') :: VARCHAR(10),
        sqlc.arg('BeforeDesc') :: VARCHAR(2000),
        sqlc.arg('AfterDesc') :: VARCHAR(2000),
        sqlc.arg('ExampleReport') :: VARCHAR(200),
        sqlc.arg('IsActive') :: BOOLEAN
    ) RETURNING *;
