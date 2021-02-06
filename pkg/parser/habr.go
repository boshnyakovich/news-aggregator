package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	"github.com/pkg/errors"
)

type HabrParser struct {
	log *logger.Logger
}

func NewHabrParser(log *logger.Logger) *HabrParser {
	return &HabrParser{
		log: log,
	}
}

func (hp *HabrParser) Start(exporter export.Exporter, habrCriteria models.HabrCriteria) error {
	url, err := hp.MakeUrl(habrCriteria)
	if err != nil {
		return err
	}

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{url},
		ParseFunc: hp.parse,
		Exporters: []export.Exporter{exporter},
	}).Start()

	return nil
}

func (hp *HabrParser) MakeUrl(habrCriteria models.HabrCriteria) (string, error) {
	if habrCriteria.Articles {
		if habrCriteria.All && !habrCriteria.Best {
			switch habrCriteria.Rating {
			case uint64(models.TEN):
				return "https://habr.com/ru/all/top10/", nil
			case uint64(models.TWENTY_FIVE):
				return "https://habr.com/ru/all/top25/", nil
			case uint64(models.FIFTY):
				return "https://habr.com/ru/all/top50/", nil
			case uint64(models.HUNDRED):
				return "https://habr.com/ru/all/top100/", nil
			default:
				return "https://habr.com/ru/all/all/", nil
			}
		} else if !habrCriteria.All && habrCriteria.Best {
			switch habrCriteria.Period {
			case string(models.DAY):
				return "https://habr.com/ru/top/", nil
			case string(models.WEEK):
				return "https://habr.com/ru/top/weekly/", nil
			case string(models.MONTH):
				return "https://habr.com/ru/top/monthly/", nil
			case string(models.YEAR):
				return "https://habr.com/ru/top/yearly/", nil
			default:
				return "https://habr.com/ru/top/", nil
			}
		}
	} else {
		return "https://habr.com/ru/news/", nil
	}

	hp.log.Error("incorrect criteria, try again")
	return "", errors.New("incorrect criteria, try again")
}

func (hp *HabrParser) parse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find(".post").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Find(".post__title_link").Attr("href")
		if !ok {
			hp.log.Error("cannot parse link")
		}

		authorLink, ok := s.Find(".post__user-info.user-info").Attr("href")
		if !ok {
			hp.log.Error("cannot parse author link")
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
