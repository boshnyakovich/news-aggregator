package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

const habr = "https://habr.com/ru/top/"

type Parser struct {
	log *logger.Logger
}

func NewParser(log *logger.Logger) *Parser {
	return &Parser{
		log: log,
	}
}

func (p *Parser) StartParseHabr(exporter export.Exporter) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{habr},
		ParseFunc: p.parseHabr,
		Exporters: []export.Exporter{exporter},
	}).Start()
}

func (p *Parser) parseHabr(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".post").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Find(".post__title_link").Attr("href")
		if !ok {
			p.log.Error("Cannot parse link")
		}

		authorLink, ok := s.Find(".post__user-info.user-info").Attr("href")
		if !ok {
			p.log.Error("Cannot parse author link")
		}

		g.Exports <- models.HabrNews{
			Author:          s.Find(".user-info__nickname_small").Text(),
			AuthorLink:      authorLink,
			Title:           s.Find(".post__title_link").Text(),
			Preview:         s.Find(".post__text_v1").Text(),
			Views:           s.Find(".post-stats__views-count").Text(),
			PublicationDate: s.Find(".post__time").Text(),
			Link:            link,
		}
	})
}
