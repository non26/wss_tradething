package main

import (
	margetstreamhandler "wsstradething/wss/bn/market_stream/marget_stream_handler"
	marketstreamws "wsstradething/wss/bn/market_stream/marget_stream_ws"
	"wsstradething/wssconfig"

	"github.com/labstack/echo/v4"
)

func main() {
	app_config, err := wssconfig.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	echo_app := echo.New()
	mkws := marketstreamws.NewMarketStreamWs()
	mkhandler := margetstreamhandler.NewKlineHandler(mkws)

	market_stream := echo_app.Group(app_config.BinanceFutureServiceName.MarketStream)
	market_stream.GET("/kline/:symbol/:interval", mkhandler.Handler)

	echo_app.Logger.Fatal(echo_app.Start(app_config.Http.Port))

}
