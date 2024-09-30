-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  "deposits" (
    "id" char(36) NOT NULL,
    "tracking_id" varchar(100) NOT NULL,
    "iban" varchar(34) NOT NULL,
    "gateway" varchar(50) NOT NULL,
    "amount" decimal(20, 4) NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    PRIMARY KEY ("id")
  )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
