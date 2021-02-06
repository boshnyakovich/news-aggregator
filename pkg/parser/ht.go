package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

const htNews = "https://hi-tech.news/"

func (p *Parser) StartParseHTNews(exporter export.Exporter) {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{htNews},
		ParseFunc: p.parseHTNews,
		Exporters: []export.Exporter{exporter},
	}).Start()

}

func (p *Parser) parseHTNews(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".post-body").Each(func(i int, s *goquery.Selection) {
		var (
			link string
			ok   bool
		)
		if link, ok = s.Find(".post-title-a").Attr("href"); !ok {
			p.log.Error("Cannot parse link")
		}

		g.Exports <- models.HTNews{
			Category: s.Find(".post-media-category").Text(),
			Title:    s.Find(".post-title-a").Text(),
			Preview:  s.Find(".the-excerpt").Text(),
			Link:     link,
		}
	})
}
