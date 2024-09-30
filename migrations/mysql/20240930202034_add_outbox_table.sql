-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  "outboxs" (
    "id" bigint unsigned NOT NULL AUTO_INCREMENT,
    "payload" text NOT NULL,
    "topic" varchar(255) NOT NULL,
    "status" varchar(50) NOT NULL,
    "created_at" datetime(3) DEFAULT NULL,
    "updated_at" datetime(3) DEFAULT NULL,
    PRIMARY KEY ("id")
  )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
