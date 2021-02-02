package service

import (
	"context"
	"github.com/boshnyakovich/news-aggregator/internal/models/domain"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"log"
)

type FontankaExporter struct {
	repo *repository.Repo
}

func NewFontankaExporter(repo *repository.Repo) *FontankaExporter {
	return &FontankaExporter{
		repo: repo,
	}
}

func (fe *FontankaExporter) Export(exports chan interface{}) {
	const op = "services.export"
	for data := range exports {
		news, ok := data.(domain.FontankaNews)
		if !ok {

		}
		err := fe.repo.InsertFontankaNews(context.Background(), news)
		if err != nil {
			log.Println(op, err)
		}
	}
}
