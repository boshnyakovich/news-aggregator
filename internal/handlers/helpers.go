package handlers

import (
	"encoding/json"
	"github.com/boshnyakovich/news-aggregator/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/valyala/fasthttp"
	"log"
)

type AlivenessHandler struct {
	db  *sqlx.DB
	log *logger.Logger
}

func NewAlivenessHandler(db *sqlx.DB, log *logger.Logger) *AlivenessHandler {
	return &AlivenessHandler{
		db:  db,
		log: log,
	}
}

func (ah *AlivenessHandler) Alive(ctx *fasthttp.RequestCtx) {
	if err := ah.db.Ping(); err != nil {
		ctx.Response.SetStatusCode(500)
	}
	decorateResponse(ctx, 200, "Alive!", "")
}

type response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

var responseInfo struct {
	Message string `json:"message"`
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

func getLimitOffset(ctx *fasthttp.RequestCtx) (uint64, uint64) {
	return uint64(ctx.QueryArgs().GetUintOrZero("limit")), uint64(ctx.QueryArgs().GetUintOrZero("offset"))
}

func getTitle(ctx *fasthttp.RequestCtx) string {
	return string(ctx.QueryArgs().Peek("title"))
}
