package handlers

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func (h *Handlers) InsertHTNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	err := h.service.InsertHTNews(ctx)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error insert ht news from storage")

		h.log.Errorf("error insert ht news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	responseInfo.Message = "Hi-tech's site data was parsed and saved"
	decorateResponse(ctx, statusCode, responseInfo, "")
}

func (h *Handlers) GetHTNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	limit, offset := getLimitOffset(ctx)

	news, err := h.service.GetHTNews(ctx, limit, offset)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting ht news from storage")

		h.log.Errorf("error getting ht news from storage: %s", err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")

}

func (h *Handlers) SearchHTNews(ctx *fasthttp.RequestCtx) {
	var errorMessage string
	statusCode := 200

	title := getTitle(ctx)

	news, err := h.service.SearchHTNews(ctx, title)
	if err != nil {
		statusCode, errorMessage = 500, fmt.Sprintf("error getting ht news from storage")

		h.log.Errorf("error search ht news from storage by title: %s: %s", title, err)
		decorateResponse(ctx, statusCode, nil, errorMessage)
		return
	}

	decorateResponse(ctx, statusCode, news, "")
}
