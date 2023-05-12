-- TABLE: "Question"
BEGIN;

CREATE TABLE "Question" (
    "Id" SERIAL NOT NULL,
    "QuestionGroupId" INT NOT NULL,
    "Question" VARCHAR(2000) NOT NULL,
    CONSTRAINT "Question_PK" PRIMARY KEY ("Id"),
    CONSTRAINT "Question_FK_QuestionGroup" FOREIGN KEY ("QuestionGroupId") REFERENCES "QuestionGroup" ("Id")
) TABLESPACE pg_default;

COMMENT ON TABLE "Question" IS 'Contains list of Question';

COMMIT;
