-- TABLE: "QuestionGroup"
BEGIN;

CREATE TABLE "QuestionGroup" (
    "Id" SERIAL NOT NULL,
    "TestId" INT NOT NULL,
    "Name" VARCHAR(2000) NOT NULL,
    "Order" INT NOT NULL,
    CONSTRAINT "QuestionGroup_PK" PRIMARY KEY ("Id"),
    CONSTRAINT "QuestionGroup_FK_Test" FOREIGN KEY ("TestId") REFERENCES "Test" ("Id")
) TABLESPACE pg_default;

COMMENT ON TABLE "QuestionGroup" IS 'Contains list of QuestionGroup';

COMMIT;
