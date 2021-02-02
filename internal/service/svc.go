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

func (s *Service) InsertFontankaNews(ctx context.Context) error {
	const op = "services.insert_fontanka_news"

	s.parser.StartParseFontanka(NewFontankaExporter(s.repo))

	return nil
}

func (s *Service) GetFontankaNews(ctx context.Context) ([]domain.FontankaNews, error) {
	const op = "services.get_fontanka_news"

	return nil, nil
}
