package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	"github.com/pkg/errors"
	"strconv"
)

type HTParser struct {
	log *logger.Logger
}

func NewHTParser(log *logger.Logger) *HTParser {
	return &HTParser{
		log: log,
	}
}

func (ht *HTParser) Start(exporter export.Exporter, page uint64) error {
	url, err := ht.MakeUrl(page)
	if err != nil {
		return err
	}

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: ht.parse,
		Exporters: []export.Exporter{exporter},
	}).Start()

	return nil
}

func (ht *HTParser) MakeUrl(page uint64) (string, error) {
	if page < 90 && page != 0 {
		pageStr := strconv.FormatUint(page, 10)
		return fmt.Sprintf("https://hi-tech.news/page/%s/", pageStr), nil
	}
	ht.log.Error("incorrect criteria page, try again")
	return "", errors.New("incorrect criteria page, try again")
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
