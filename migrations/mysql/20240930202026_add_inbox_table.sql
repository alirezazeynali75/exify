-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  "inboxs" (
    "event_id" varchar(191) NOT NULL,
    "updated_at" datetime(3) DEFAULT NULL,
    "created_at" datetime(3) DEFAULT NULL,
    PRIMARY KEY ("event_id")
  )
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
