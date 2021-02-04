package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/boshnyakovich/news-aggregator/internal/models/dao"
	"github.com/boshnyakovich/news-aggregator/internal/models/domain"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"log"
	"time"
)

type Repo struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewRepo(db *sqlx.DB, log *logger.Logger) *Repo {
	return &Repo{
		db:  db,
		log: log,
	}
}

const (
	habrTableName = "habr_news"
	htTableName   = "ht_news"
)

func (r *Repo) InsertHTNews(ctx context.Context, news domain.HTNews) error {
	const op = "repositories.insert_ht_news"

	log.Println(op, news)
	var ht dao.HTNews
	id, err := uuid.NewV4()
	if err != nil {
		return errors.Wrap(err, op)
	}

	ht.ID = id.String()
	ht.Category = news.Category
	ht.Title = news.Title
	ht.Preview = news.Preview
	ht.Link = news.Link
	ht.CreatedAt = time.Now()

	columns, values := ht.InsertColumns(), ht.Values()

	sql, args, err := squirrel.
		Insert(htTableName).
		Columns(columns...).
		Values(values...).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, op)
	}
	_, err = r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (r *Repo) GetHTNews() ([]dao.HTNews, error) {
	const op = "repositories.get_ht_news"

	return nil, nil
}

func (r *Repo) InsertHabrNews(ctx context.Context, news domain.HabrNews) error {
	const op = "repositories.insert_habr_news"

	var hn dao.HabrNews
	id, err := uuid.NewV4()
	if err != nil {
		return errors.Wrap(err, op)
	}

	hn.ID = id.String()
	hn.Author = news.Author
	hn.AuthorLink = news.AuthorLink
	hn.Title = news.Title
	hn.Preview = news.Preview
	hn.Views = news.Views
	hn.PublicationDate = news.PublicationDate
	hn.Link = news.Link
	hn.CreatedAt = time.Now()

	columns, values := hn.InsertColumns(), hn.Values()

	sql, args, err := squirrel.
		Insert(habrTableName).
		Columns(columns...).
		Values(values...).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, op)
	}

	_, err = r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, op)
	}

	return nil
}

func (r *Repo) GetHabrNews() ([]dao.HTNews, error) {
	const op = "repositories.get_habr_news"

	return nil, nil
}
