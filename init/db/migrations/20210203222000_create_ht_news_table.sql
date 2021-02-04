-- +goose Up
CREATE TABLE habr_news
(
    "id"               uuid      NOT NULL,
    "author"           text,
    "author_link"      text,
    "title"            text,
    "preview"          text,
    "views"            text,
    "publication_date" text,
    "link"             text,
    "created_at"       timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY ("id")
);

-- +goose Down
DROP TABLE habr_news;
