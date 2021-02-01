package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	"regexp"
)

type Parser struct {
	log *logger.Logger
}

func NewParser(log *logger.Logger) *Parser {
	return &Parser{
		log: log,
	}
}

func (p *Parser) StartParse(website string) {
	switch website {
	case "habr":
		{
			geziyor.NewGeziyor(&geziyor.Options{
				StartURLs: []string{"https://habr.com/ru/top/"},
				ParseFunc: p.parseHabr,
				Exporters: []export.Exporter{&export.JSON{}},
			}).Start()
		}
	case "fontanka":
		{
			geziyor.NewGeziyor(&geziyor.Options{
				StartURLs: []string{"https://www.fontanka.ru/"},
				ParseFunc: p.parseFontanka,
				Exporters: []export.Exporter{&export.JSON{}},
			}).Start()
		}
	}
}

func (p *Parser) parseHabr(g *geziyor.Geziyor, r *client.Response) {

}

func (p *Parser) parseFontanka(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".HFa1").Each(func(i int, s *goquery.Selection) {
		var (
			link string
			ok   bool
		)
		if link, ok = s.Find(".HFd-").Attr("href"); !ok {
			p.log.Error("Cannot find link")
		} else {
			match, err := regexp.MatchString("http", link)
			if err == nil && !match {
				link = "https://www.fontanka.ru/" + link
			} else if err != nil {
				p.log.Error("cannot parse link by regexp")
			}
		}

		g.Exports <- map[string]interface{}{
			"title":            s.Find(".HFd-").Text(),
			"publication_date": s.Find(".HFyn").Text(),
			"link":             link,
		}
	})
}
