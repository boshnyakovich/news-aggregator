-- +goose Up
DROP TABLE fontanka_news;

-- +goose Down
CREATE TABLE fontanka_news
(
    "id"              uuid      NOT NULL,
    "title"           text,
    "publicationDate" text,
    "link"            text,
    "created_at"      timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY ("id")
);