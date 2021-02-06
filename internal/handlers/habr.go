package handlers

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

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

	limit, offset := getLimitOffset(ctx)

	news, err := h.service.GetHabrNews(ctx, limit, offset)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting top habr news from storage")

		h.log.Errorf("error getting top habr news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}

func (h *Handlers) SearchHabrNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	title := getTitle(ctx)

	news, err := h.service.SearchHabrNews(ctx, title)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error search habr news from storage by title")

		h.log.Errorf("error search habr news from storage by title: %s: %s", title, err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}
