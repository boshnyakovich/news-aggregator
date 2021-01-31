package service

import (
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/boshnyakovich/news-aggregator/pkg/parser"
)

type Service struct {
	repo *repository.Repo
	log  *logger.Logger
}

const habr = "habr"
const fourPDA = "fourPDA"

	func NewService(repo *repository.Repo, log *logger.Logger) *Service {
	return &Service{
		repo: repo,
		log:  log,
	}
}

func (s *Service) GetTopHabrNews() ([]models.HabrNews, error) {
	parser.StartParse(habr)

	return nil, nil
}

func (s *Service) GetFourPDANewsByDate() ([]models.HabrNews, error) {
	parser.StartParse(fourPDA)

	return nil, nil
}