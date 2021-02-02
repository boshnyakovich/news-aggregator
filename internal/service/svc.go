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

const (
	habr     = "habr"
	fontanka = "fontanka"
)

func (s *Service) InsertHabrNews(ctx context.Context) error {
	const op = "services.insert_habr_news"
	s.parser.StartParse(habr)

	// TODO: get data
	var news []domain.HabrNews
	err := s.repo.InsertHabrNews(ctx, news)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetHabrNews(ctx context.Context) ([]domain.HabrNews, error) {
	const op = "services.get_habr_news"
	s.parser.StartParse(habr)

	return nil, nil
}

func (s *Service) InsertFontankaNews(ctx context.Context) error {
	const op = "services.insert_fontanka_news"

	s.parser.StartParse(fontanka)

	// TODO: get data
	var news []domain.FontankaNews
	err := s.repo.InsertFontankaNews(ctx, news)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetFontankaNews(ctx context.Context) ([]domain.FontankaNews, error) {
	const op = "services.get_fontanka_news"

	s.parser.StartParse(fontanka)

	return nil, nil
}
