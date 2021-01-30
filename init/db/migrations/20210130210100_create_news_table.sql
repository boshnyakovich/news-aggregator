-- +goose Up
CREATE TABLE news
(
    "id" uuid NOT NULL,
    "title" text,
    "link" text,
    PRIMARY KEY ("id")
);

-- +goose Down
DROP TABLE news;