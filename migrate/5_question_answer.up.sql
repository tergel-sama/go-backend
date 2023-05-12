-- TABLE: "QuestionAnswer"
BEGIN;

CREATE TABLE "QuestionAnswer" (
    "Id" SERIAL NOT NULL,
    "QuestionId" INT NOT NULL,
    "AnswerText" VARCHAR(2000) NOT NULL,
    "Score" INT NOT NULL,
    CONSTRAINT "QuestionAnswer_PK" PRIMARY KEY ("Id"),
    CONSTRAINT "QuestionAnswer_FK_Question" FOREIGN KEY ("QuestionId") REFERENCES "Question" ("Id")
) TABLESPACE pg_default;

COMMENT ON TABLE "QuestionAnswer" IS 'Contains list of QuestionAnswer';

COMMIT;
