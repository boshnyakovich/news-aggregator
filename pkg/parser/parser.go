package parser

import (
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func StartParse(website string) {
	switch website {
	case "habr":
		{
			geziyor.NewGeziyor(&geziyor.Options{
				StartURLs: []string{"https://habr.com/ru/top/"},
				ParseFunc: parseHabr,
				Exporters: []export.Exporter{&export.JSON{}},
			}).Start()
		}
	case "fourPDA":
		{
			geziyor.NewGeziyor(&geziyor.Options{
				StartURLs: []string{"https://4pda.ru/"},
				ParseFunc: parseFourPDA,
				Exporters: []export.Exporter{&export.JSON{}},
			}).Start()
		}
	}
}

func parseHabr(g *geziyor.Geziyor, r *client.Response) {

}

func parseFourPDA(g *geziyor.Geziyor, r *client.Response) {

}
