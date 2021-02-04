package service

import (
	"context"
	"github.com/boshnyakovich/news-aggregator/internal/models/domain"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/boshnyakovich/news-aggregator/pkg/parser"
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
	const op = "services.insert_habr_news"

	s.parser.StartParseHabr(NewHabrExporter(s.repo))

	return nil

}

func (s *Service) GetHabrNews(ctx context.Context) ([]domain.HabrNews, error) {
	const op = "services.get_habr_news"

	return nil, nil
}

func (s *Service) InsertHTNews(ctx context.Context) error {
	const op = "services.insert_ht_news"

	s.parser.StartParseHTNews(NewHTNewsExporter(s.repo))

	return nil
}

func (s *Service) GetHTNews(ctx context.Context) ([]domain.HTNews, error) {
	const op = "services.get_ht_news"

	return nil, nil
}
