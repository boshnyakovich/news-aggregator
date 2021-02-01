package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
	habrTableName     = "habr_news"
	fontankaTableName = "fontanka_news"
)

func (r *Repo) InsertFontankaNews(ctx context.Context, news []models.FontankaNews) error {
	const op = "repositories.insert_fontanka_news"

	for _, n := range news {
		var fn FontankaNews
		fn.Title = n.Title
		fn.PublicationDate = n.PublicationDate
		fn.Link = n.Link
		fn.CreatedAt = time.Now()

		columns, values := fn.InsertColumns(), fn.Values()

		sql, args, err := squirrel.
			Insert(fontankaTableName).
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
	}

	return nil
}

func (r *Repo) GetFontankaNews() ([]FontankaNews, error) {
	const op = "repositories.get_fontanka_news"

	return nil, nil
}
