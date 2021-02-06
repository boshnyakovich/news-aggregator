package exporters

import (
	"context"
	"github.com/boshnyakovich/news-aggregator/internal/models/domain"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"log"
)

type HabrExporter struct {
	repo *repository.Repo
}

func NewHabrExporter(repo *repository.Repo) *HabrExporter {
	return &HabrExporter{
		repo: repo,
	}
}

func (he *HabrExporter) Export(exports chan interface{}) {
	const op = "services.export"
	for data := range exports {
		news, ok := data.(domain.HabrNews)
		if !ok {

		}
		err := he.repo.InsertHabrNews(context.Background(), news)
		if err != nil {
			log.Println(op, err)
		}
	}
}
