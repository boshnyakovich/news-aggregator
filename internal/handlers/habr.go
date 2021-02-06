package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/boshnyakovich/news-aggregator/internal/exporters"
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
	"github.com/boshnyakovich/news-aggregator/pkg/fasthttpserver"
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
	var errorMessage string
	statusCode := 200

	body := ctx.Request.Body()
	var habrCriteria models.HabrCriteria
	if err := json.Unmarshal(body, &habrCriteria); err != nil {
		statusCode, errorMessage = 400, "incorrect body"

		hh.log.Errorf("error unmarshaling habr criteria: %s", err)

		fasthttpserver.NewResponseBuilder(ctx).
			SetStatusCode(statusCode).
			SetError(errorMessage).
			Send()

		return
	}

	if err := hh.parser.Start(exporters.NewHabrExporter(hh.repo), habrCriteria); err != nil {
		statusCode, errorMessage = 400, "incorrect criteria, try again"

		fasthttpserver.NewResponseBuilder(ctx).
			SetStatusCode(statusCode).
			SetError(errorMessage).
			Send()

		return
	}

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
