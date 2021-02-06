package services

import (
	"context"
	"github.com/boshnyakovich/news-aggregator/internal/exporters"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/boshnyakovich/news-aggregator/pkg/parser"
	"github.com/pkg/errors"
)

type Service struct {
	repo   *repository.Repo
	parser *parser.Parser
	log    *logger.Logger
}

func NewService(repo *repository.Repo, parser *parser.Parser, log *logger.Logger) *Service {
	return &Service{
		repo:   repo,
		parser: parser,
		log:    log,
	}
}

func (s *Service) InsertHabrNews(ctx context.Context) error {
	const op = "services.InsertHabrNews"

	s.parser.StartParseHabr(exporters.NewHabrExporter(s.repo))

	return nil
}

func (s *Service) GetHabrNews(ctx context.Context, limit uint64, offset uint64) (result []models.HabrNews, err error) {
	const op = "services.GetHabrNews"

	resultRepo, err := s.repo.GetHabrNews(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	for _, rr := range resultRepo {
		r := models.HabrNews{
			ID:              rr.ID,
			Author:          rr.Author,
			AuthorLink:      rr.AuthorLink,
			Title:           rr.Title,
			Preview:         rr.Preview,
			Views:           rr.Views,
			PublicationDate: rr.PublicationDate,
			Link:            rr.Link,
			CreatedAt:       rr.CreatedAt,
		}
		result = append(result, r)
	}

	return result, nil
}

func (s *Service) SearchHabrNews(ctx context.Context, title string) (result []models.HabrNews, err error) {
	const op = "services.SearchHabrNews"

	resultRepo, err := s.repo.SearchHabrNews(ctx, title)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	for _, rr := range resultRepo {
		r := models.HabrNews{
			ID:              rr.ID,
			Author:          rr.Author,
			AuthorLink:      rr.AuthorLink,
			Title:           rr.Title,
			Preview:         rr.Preview,
			Views:           rr.Views,
			PublicationDate: rr.PublicationDate,
			Link:            rr.Link,
			CreatedAt:       rr.CreatedAt,
		}
		result = append(result, r)
	}

	return result, nil
}
