package services

import (
	"context"
	"github.com/boshnyakovich/news-aggregator/internal/exporters"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/pkg/errors"
)

func (s *Service) InsertHTNews(ctx context.Context) error {
	const op = "services.InsertHTNews"

	s.parser.StartParseHTNews(exporters.NewHTNewsExporter(s.repo))

	return nil
}

func (s *Service) GetHTNews(ctx context.Context, limit uint64, offset uint64) (result []models.HTNews, err error) {
	const op = "services.GetHTNews"

	resultRepo, err := s.repo.GetHTNews(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	for _, rr := range resultRepo {
		r := models.HTNews{
			ID:        rr.ID,
			Category:  rr.Category,
			Title:     rr.Title,
			Preview:   rr.Preview,
			Link:      rr.Link,
			CreatedAt: rr.CreatedAt,
		}
		result = append(result, r)
	}

	return result, nil
}

func (s *Service) SearchHTNews(ctx context.Context, title string) (result []models.HTNews, err error) {
	const op = "services.SearchHabrNews"

	resultRepo, err := s.repo.SearchHTNews(ctx, title)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	for _, rr := range resultRepo {
		r := models.HTNews{
			ID:        rr.ID,
			Category:  rr.Category,
			Title:     rr.Title,
			Preview:   rr.Preview,
			Link:      rr.Link,
			CreatedAt: rr.CreatedAt,
		}
		result = append(result, r)
	}

	return result, nil
}
