package margetstreamhandler

import (
	"net/http"
	marketstreamws "wsstradething/wss/bn/market_stream/marget_stream_ws"
	marketstreammodel "wsstradething/wss/bn/market_stream/market_stream_model"

	"github.com/labstack/echo/v4"
)

type klineHandler struct {
	margetStream marketstreamws.IMarketStreamWs
}

func NewKlineHandler(
	margetStream marketstreamws.IMarketStreamWs,
) *klineHandler {
	return &klineHandler{
		margetStream,
	}
}

func (k *klineHandler) GetBody(c echo.Context) *marketstreammodel.KlineRequest {
	s := c.Param("symbol")
	i := c.Param("interval")
	m := marketstreammodel.KlineRequest{
		Symbol:   s,
		Interval: i,
	}
	return &m
}

func (k *klineHandler) Handler(c echo.Context) error {
	println("get body")
	m := k.GetBody(c)
	println("")
	err := k.margetStream.KlineWs(c.Request().Context(), c.Response(), c.Request(), nil, m.Symbol, m.Interval)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			err.Error(),
		)
	}

	return nil
}
