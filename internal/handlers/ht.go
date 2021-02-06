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

type HTHandlers struct {
	repo   *repository.Repo
	parser *parser.HTParser
	log    *logger.Logger
}

func NewHTHandlers(repo *repository.Repo, parser *parser.HTParser, log *logger.Logger) *HTHandlers {
	return &HTHandlers{
		repo:   repo,
		parser: parser,
		log:    log,
	}
}

func (ht *HTHandlers) Insert(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	body := ctx.Request.Body()
	var htCriteria models.HTCriteria
	if err := json.Unmarshal(body, &htCriteria); err != nil {
		statusCode, errorMessage = 400, "incorrect body"

		ht.log.Errorf("error unmarshaling ht-news criteria: %s", err)

		fasthttpserver.NewResponseBuilder(ctx).
			SetStatusCode(statusCode).
			SetError(errorMessage).
			Send()

		return
	}

	if htCriteria.Category != string(models.SMARTPHONES) && htCriteria.Category != string(models.MEDICINE) && htCriteria.Category != string(models.OTHER)  && htCriteria.Category != "" {
		statusCode, errorMessage = 400, "incorrect criteria category, try again"

		fasthttpserver.NewResponseBuilder(ctx).
			SetStatusCode(statusCode).
			SetError(errorMessage).
			Send()

		return
	}

	if err := ht.parser.Start(exporters.NewHTNewsExporter(ht.repo, htCriteria.Category), htCriteria.Page); err != nil {
		statusCode, errorMessage = 400, "incorrect criteria, try again"

		fasthttpserver.NewResponseBuilder(ctx).
			SetStatusCode(statusCode).
			SetError(errorMessage).
			Send()

		return
	}

	responseInfo.Message = "Hi-tech's site data was parsed and saved"
	decorateResponse(ctx, statusCode, responseInfo, "")
}

func (ht *HTHandlers) Get(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	limit, offset := getLimitOffset(ctx)

	news, err := ht.repo.GetHTNews(ctx, limit, offset)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting ht-news from storage")

		ht.log.Errorf("error getting ht news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")

}

func (ht *HTHandlers) Search(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	title := getTitle(ctx)

	news, err := ht.repo.SearchHTNews(ctx, title)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting ht-news from storage")

		ht.log.Errorf("error search ht news from storage by title: %s: %s", title, err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}
