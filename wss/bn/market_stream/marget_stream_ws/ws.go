package marketstreamws

import (
	"context"
	"net/http"
)

type IMarketStreamWs interface {
	KlineWs(
		c context.Context,
		w http.ResponseWriter,
		r *http.Request,
		h http.Header,
		symbol string,
		interval string,
	) error
}

type marketStreamWs struct{}

func NewMarketStreamWs() IMarketStreamWs {
	return &marketStreamWs{}
}
