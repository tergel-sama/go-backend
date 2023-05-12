-- TABLE: "QuestionGroupAnswer"
BEGIN;

CREATE TABLE "QuestionGroupAnswer" (
    "Id" SERIAL NOT NULL,
    "QuestionGroupId" INT NOT NULL,
    "MaxScr" INT NOT NULL,
    "MinScr" INT NOT NULL,
    "AnswerText" VARCHAR(2000) NOT NULL,
    CONSTRAINT "QuestionGroupAnswer_PK" PRIMARY KEY ("Id"),
    CONSTRAINT "QuestionGroupAnswer_FK_QuestionGroup" FOREIGN KEY ("QuestionGroupId") REFERENCES "QuestionGroup" ("Id")
) TABLESPACE pg_default;

COMMENT ON TABLE "QuestionGroupAnswer" IS 'Contains list of QuestionGroupAnswer';

COMMIT;
