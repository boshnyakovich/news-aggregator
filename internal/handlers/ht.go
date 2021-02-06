package handlers

import (
	"fmt"
	"github.com/boshnyakovich/news-aggregator/internal/exporters"
	"github.com/boshnyakovich/news-aggregator/internal/repository"
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
	statusCode := 200

	ht.parser.Start(exporters.NewHTNewsExporter(ht.repo))

	responseInfo.Message = "Hi-tech's site data was parsed and saved"
	decorateResponse(ctx, statusCode, responseInfo, "")
}

func (ht *HTHandlers) Get(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	limit, offset := getLimitOffset(ctx)

	news, err := ht.repo.GetHTNews(ctx, limit, offset)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting ht news from storage")

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
		statusCode, errorMessage = 500, fmt.Sprintf("error getting ht news from storage")

		ht.log.Errorf("error search ht news from storage by title: %s: %s", title, err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}
