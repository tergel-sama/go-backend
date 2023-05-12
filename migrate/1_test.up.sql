-- TABLE: "Test"
BEGIN;

CREATE TABLE "Test" (
    "Id" SERIAL NOT NULL,
    "Name" VARCHAR(200) NOT NULL,
    "Desc" VARCHAR(2000) NOT NULL,
    "Img" VARCHAR(200) NOT NULL,
    "Minute" INT NOT NULL,
    "AgeCls" VARCHAR(10) NOT NULL,
    "BeforeDesc" VARCHAR(2000) NOT NULL,
    "AfterDesc" VARCHAR(2000) NOT NULL,
    "ExampleReport" VARCHAR(200) NOT NULL,
    "IsActive" BOOLEAN NOT NULL,
    CONSTRAINT "Test_PK" PRIMARY KEY ("Id"),
    CONSTRAINT "Test_AgeCls_CHECK_Type" CHECK ("AgeCls" IN ('child', 'adult','teenager'))
) TABLESPACE pg_default;

COMMENT ON TABLE "Test" IS 'Contains list of test';

COMMIT;
