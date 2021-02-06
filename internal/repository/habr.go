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

const habrTableName = "habr_news"

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

func (r *Repo) GetHabrNews(ctx context.Context, limit uint64, offset uint64) (result []dao.HabrNews, err error) {
	const op = "repositories.get_habr_news"

	var habrRepo dao.HabrNews

	columns := habrRepo.Columns()

	var (
		sql  string
		args []interface{}
	)

	sqlBuilder := squirrel.
		Select(columns...).
		From(habrTableName).
		OrderBy("created_at DESC").
		PlaceholderFormat(squirrel.Dollar)

	if limit != 0 && offset != 0 {
		sqlBuilder = sqlBuilder.Limit(limit).Offset(offset)
	}

	sql, args, err = sqlBuilder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	for rows.Next() {
		hn := dao.HabrNews{}
		if err = rows.Scan(
			&hn.ID,
			&hn.Author,
			&hn.AuthorLink,
			&hn.Title,
			&hn.Preview,
			&hn.Views,
			&hn.PublicationDate,
			&hn.Link,
			&hn.CreatedAt,
		); err != nil {
			r.log.Errorf("error getting habr news list from db: %s", err)
		}
		result = append(result, hn)
	}

	return result, nil
}
