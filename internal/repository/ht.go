package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/internal/repository/models"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"time"
)

const htTableName = "ht_news"

func (r *Repo) InsertHTNews(ctx context.Context, news models.HTNews) error {
	const op = "repositories.InsertHTNews"
	var ht repository.HTNews
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

func (r *Repo) GetHTNews(ctx context.Context, limit uint64, offset uint64) (result []repository.HTNews, err error) {
	const op = "repositories.GetHTNews"

	var htRepo repository.HTNews

	columns := htRepo.Columns()

	var (
		sql  string
		args []interface{}
	)

	sqlBuilder := squirrel.
		Select(columns...).
		From(htTableName).
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
		hn := repository.HTNews{}
		if err = rows.Scan(
			&hn.ID,
			&hn.Category,
			&hn.Title,
			&hn.Preview,
			&hn.Link,
			&hn.CreatedAt,
		); err != nil {
			r.log.Errorf("error getting ht news list from db: %s", err)
		}
		result = append(result, hn)
	}

	return result, nil
}

func (r *Repo) SearchHTNews(ctx context.Context, title string) (result []repository.HTNews, err error) {
	const op = "repositories.SearchHTNews"

	sql := "SELECT * FROM " + htTableName + " WHERE title similar to $1"

	rows, err := r.db.QueryContext(ctx, sql, "%"+title+"%")
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	for rows.Next() {
		hn := repository.HTNews{}
		if err = rows.Scan(
			&hn.ID,
			&hn.Category,
			&hn.Title,
			&hn.Preview,
			&hn.Link,
			&hn.CreatedAt,
		); err != nil {
			r.log.Errorf("error getting ht news by title from db: %s", err)
		}
		result = append(result, hn)
	}

	return result, nil
}
