package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/internal/repository/models"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
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

func (r *Repo) InsertHabrNews(ctx context.Context, news models.HabrNews) error {
	const op = "repositories.InsertHabrNews"

	var hn repository.HabrNews
	id, err := uuid.NewV4()
	if err != nil {
		return errors.Wrap(err, op)
	}
	hn.ID = id.String()

	preview := news.Preview
	preview = strings.Replace(preview, "\n", "", -1)
	preview = strings.Replace(preview, "<br>", "", -1)
	hn.Preview = preview

	hn.Author = news.Author
	hn.AuthorLink = news.AuthorLink
	hn.Title = news.Title
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

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return errors.Wrap(err, op)
	}
	defer rows.Close()

	return nil
}

func (r *Repo) GetHabrNews(ctx context.Context, limit uint64, offset uint64) (result []repository.HabrNews, err error) {
	const op = "repositories.GetHabrNews"

	var habrRepo repository.HabrNews

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
		hn := repository.HabrNews{}
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

func (r *Repo) SearchHabrNews(ctx context.Context, title string) (result []repository.HabrNews, err error) {
	const op = "repositories.SearchHabrNews"

	sql := "SELECT * FROM habr_news WHERE title similar to $1"

	rows, err := r.db.QueryContext(ctx, sql, "%"+title+"%")
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	for rows.Next() {
		hn := repository.HabrNews{}
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
			r.log.Errorf("error getting habr news by title from db: %s", err)
		}
		result = append(result, hn)
	}

	return result, nil
}
