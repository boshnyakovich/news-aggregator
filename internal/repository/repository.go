package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/boshnyakovich/news-aggregator/internal/models/dao"
	"github.com/boshnyakovich/news-aggregator/internal/models/domain"
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

func (r *Repo) InsertFontankaNews(ctx context.Context, news []domain.FontankaNews) error {
	const op = "repositories.insert_fontanka_news"

	for _, n := range news {
		var fn dao.FontankaNews
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

func (r *Repo) GetFontankaNews() ([]dao.FontankaNews, error) {
	const op = "repositories.get_fontanka_news"

	return nil, nil
}

func (r *Repo) InsertHabrNews(ctx context.Context, news []domain.HabrNews) error {
	const op = "repositories.insert_habr_news"

	for _, n := range news {
		var hn dao.HabrNews
		hn.Author              = n.Author
		hn.AuthorLink          = n.AuthorLink
		hn.Title               = n.Title
		hn.Preview             = n.Preview
		hn.Views               = n.Views
		hn.PublicationDate     = n.PublicationDate
		hn.Link                = n.Link
		hn.CreatedAt           = time.Now()


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
	}

	return nil
}

func (r *Repo) GetHabrNews() ([]dao.FontankaNews, error) {
	const op = "repositories.get_habr_news"

	return nil, nil
}