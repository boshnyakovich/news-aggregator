-- +goose Up
CREATE TABLE ht_news
(
    "id"               uuid      NOT NULL,
    "category"         text,
    "title"            text,
    "preview"          text,
    "publication_date" text,
    "link"             text,
    "created_at"       timestamp NOT NULL DEFAULT now(),
    PRIMARY KEY ("id")
);

-- +goose Down
DROP TABLE ht_news;
