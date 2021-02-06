package exporters

import (
	"context"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"log"
)

type HTNewsExporter struct {
	repo *repository.Repo
}

func NewHTNewsExporter(repo *repository.Repo) *HTNewsExporter {
	return &HTNewsExporter{
		repo: repo,
	}
}

func (ht *HTNewsExporter) Export(exports chan interface{}) {
	const op = "services.export"
	for data := range exports {
		news, ok := data.(models.HTNews)
		if !ok {

		}
		err := ht.repo.InsertHTNews(context.Background(), news)
		if err != nil {
			log.Println(op, err)
		}
	}
}
