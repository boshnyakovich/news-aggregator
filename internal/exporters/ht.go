package exporters

import (
	"context"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"log"
)

type HTNewsExporter struct {
	repo              *repository.Repo
	categoryInRequest string
}

func NewHTNewsExporter(repo *repository.Repo, categoryInRequest string) *HTNewsExporter {
	return &HTNewsExporter{
		repo:              repo,
		categoryInRequest: categoryInRequest,
	}
}

func (ht *HTNewsExporter) Export(exports chan interface{}) {
	const op = "exporters.ht.export"
	for data := range exports {
		news, ok := data.(models.HTNews)
		if !ok {
			log.Println(op, news)
		}

		if ht.categoryInRequest == "" || news.Category == ht.categoryInRequest {
			err := ht.repo.InsertHTNews(context.Background(), news)
			if err != nil {
				log.Println(op, err)
			}
		}
	}
}
