-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  "withdrawals" (
    "id" varchar(191) NOT NULL,
    "tracking_id" longtext,
    "destination" longtext,
    "gateway" longtext,
    "amount" longtext,
    "status" longtext,
    "created_at" datetime(3) DEFAULT NULL,
    "updated_at" datetime(3) DEFAULT NULL,
    PRIMARY KEY ("id")
  )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
