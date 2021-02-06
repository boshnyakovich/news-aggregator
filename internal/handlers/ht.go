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

	decorateResponse(ctx, statusCode, nil, "")

}

func (h *Handlers) GetHTNewsNews(ctx *fasthttp.RequestCtx) {
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
