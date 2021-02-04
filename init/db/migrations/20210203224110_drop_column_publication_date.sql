-- +goose Up
ALTER TABLE "ht_news" DROP COLUMN "publication_date";

-- +goose Down
ALTER TABLE "ht_news" ADD COLUMN "publication_date" text;

