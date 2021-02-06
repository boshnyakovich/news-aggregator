package handlers

import (
	"fmt"
	"github.com/boshnyakovich/news-aggregator/internal/exporters"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/boshnyakovich/news-aggregator/pkg/parser"
	"github.com/valyala/fasthttp"
)

type HabrHandlers struct {
	repo   *repository.Repo
	parser *parser.HabrParser
	log    *logger.Logger
}

func NewHabrHandlers(repo *repository.Repo, parser *parser.HabrParser, log *logger.Logger) *HabrHandlers {
	return &HabrHandlers{
		repo:   repo,
		parser: parser,
		log:    log,
	}
}

func (hh *HabrHandlers) Insert(ctx *fasthttp.RequestCtx) {
	statusCode := 200

	hh.parser.Start(exporters.NewHabrExporter(hh.repo))

	responseInfo.Message = "Habr's site data was parsed and saved"
	decorateResponse(ctx, statusCode, responseInfo, "")
}

func (hh *HabrHandlers) Get(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	limit, offset := getLimitOffset(ctx)

	news, err := hh.repo.GetHabrNews(ctx, limit, offset)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting habr news from storage")

		hh.log.Errorf("error getting habr news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}

func (hh *HabrHandlers) Search(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	title := getTitle(ctx)

	news, err := hh.repo.SearchHabrNews(ctx, title)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error search habr news from storage by title")

		hh.log.Errorf("error search habr news from storage by title: %s: %s", title, err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}
