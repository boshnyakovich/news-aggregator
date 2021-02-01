package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/boshnyakovich/news-aggregator/internal/service"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
	"log"
)

type Handlers struct {
	service *service.Service
	db      *sqlx.DB
	log     *logger.Logger
}

func NewHandlers(service *service.Service, db *sqlx.DB, log *logger.Logger) *Handlers {
	return &Handlers{
		service: service,
		db:      db,
		log:     log,
	}
}

func (h *Handlers) InsertHabrNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	err := h.service.InsertHabrNews(ctx)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting top habr news from storage")

		h.log.Errorf("error getting top habr news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, nil, "")

}

func (h *Handlers) GetHabrNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	news, err := h.service.GetHabrNews(ctx)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting top habr news from storage")

		h.log.Errorf("error getting top habr news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}

func (h *Handlers) InsertFontankaNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	err := h.service.InsertFontankaNews(ctx)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error insert fontanka news from storage")

		h.log.Errorf("error insert fontanka news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, nil, "")

}

func (h *Handlers) GetFontankaNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	news, err := h.service.GetFontankaNews(ctx)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting fontanka news from storage")

		h.log.Errorf("error getting fontanka news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")

}

func (h *Handlers) LivenessHandler(ctx *fasthttp.RequestCtx) {
	if err := h.db.Ping(); err != nil {
		ctx.Response.SetStatusCode(500)
	}
	decorateResponse(ctx, 200, "Alive!", "")
}

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func decorateResponse(ctx *fasthttp.RequestCtx, statusCode int, data interface{}, errorMessage string) {
	var resp response

	if errorMessage == "" {
		resp.Data = data
	} else {
		resp.Data = map[int]int{}
		resp.Error = errorMessage
	}

	body, err := json.Marshal(resp)
	if err != nil {
		log.Println("error marshaling response", err)
	}

	ctx.Response.SetBody(body)
	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.SetContentType("application/json")
}
