package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/boshnyakovich/news-aggregator/internal/models/domain"
	"github.com/boshnyakovich/news-aggregator/internal/models/dao"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"time"
)

const htTableName   = "ht_news"

func (r *Repo) InsertHTNews(ctx context.Context, news domain.HTNews) error {
	const op = "repositories.insert_ht_news"
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

func (r *Repo) GetHTNews(ctx context.Context, limit uint64, offset uint64) (result []dao.HTNews, err error) {
	const op = "repositories.get_ht_news"

	var htRepo dao.HTNews

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
		hn := dao.HTNews{}
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
