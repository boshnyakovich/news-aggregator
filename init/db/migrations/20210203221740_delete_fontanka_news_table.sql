-- +goose Up
CREATE TABLE fontanka_news
(
    "id"              uuid      NOT NULL,
    "title"           text,
    "publicationDate" text,
    "link"            text,
    "created_at"      timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY ("id")
);

-- +goose Down
DROP TABLE fontanka_news;20210202230846_create_fontanka_news_table.sql