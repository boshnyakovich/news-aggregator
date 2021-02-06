package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

const htNews = "https://hi-tech.news/"

type HTParser struct {
	log *logger.Logger
}

func NewHTParser(log *logger.Logger) *HTParser {
	return &HTParser{
		log: log,
	}
}

func (ht *HTParser) Start(exporter export.Exporter) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{htNews},
		ParseFunc: ht.parse,
		Exporters: []export.Exporter{exporter},
	}).Start()
}

func (ht *HTParser) parse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".post-body").Each(func(i int, s *goquery.Selection) {
		var (
			link string
			ok   bool
		)
		if link, ok = s.Find(".post-title-a").Attr("href"); !ok {
			ht.log.Error("cannot parse link")
		}

		g.Exports <- models.HTNews{
			Category: s.Find(".post-media-category").Text(),
			Title:    s.Find(".post-title-a").Text(),
			Preview:  s.Find(".the-excerpt").Text(),
			Link:     link,
		}
	})
}
