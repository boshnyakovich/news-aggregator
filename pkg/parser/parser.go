package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/boshnyakovich/news-aggregator/internal/models/domain"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	"log"
	"time"
)

const (
	habr   = "https://habr.com/ru/top/"
	htNews = "https://hi-tech.news/"
)

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

func (p *Parser) StartParseHTNews(exporter export.Exporter) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{htNews},
		ParseFunc: p.parseHTNews,
		Exporters: []export.Exporter{exporter},
	}).Start()

}

func (p *Parser) parseHabr(g *geziyor.Geziyor, r *client.Response) {
	// start := time.Now()
	// i := 0

	r.HTMLDoc.Find(".post").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Find(".post__title_link").Attr("href")
		if !ok {
			p.log.Error("Cannot parse link")
		}

		authorLink, ok := s.Find(".post__user-info.user-info").Attr("href")
		if !ok {
			p.log.Error("Cannot parse author link")
		}

		g.Exports <- domain.HabrNews{
			Author:          s.Find(".user-info__nickname_small").Text(),
			AuthorLink:      authorLink,
			Title:           s.Find(".post__title_link").Text(),
			Preview:         s.Find(".post__text_v1").Text(),
			Views:           s.Find(".post-stats__views-count").Text(),
			PublicationDate: s.Find(".post__time").Text(),
			Link:            link,
		}
		// log.Println(i+1, time.Since(start).Minutes())
	})
	// log.Println(i, "AFTER:", time.Since(start).Minutes())
}

func (p *Parser) parseHTNews(g *geziyor.Geziyor, r *client.Response) {
	start := time.Now()
	i := 0

	r.HTMLDoc.Find(".post-body").Each(func(i int, s *goquery.Selection) {
		var (
			link string
			ok   bool
		)
		log.Println(1)
		if link, ok = s.Find(".post-title-a").Attr("href"); !ok {
			p.log.Error("Cannot parse link")
		}
		log.Println(2)
		g.Exports <- domain.HTNews{
			Category: s.Find(".post-media-category").Text(),
			Title:    s.Find(".post-title-a").Text(),
			Preview:  s.Find(".the-excerpt").Text(),
			Link:     link,
		}

		log.Println(i+1, time.Since(start).Minutes())
	})

	log.Println(i, "AFTER:", time.Since(start).Minutes())
}
