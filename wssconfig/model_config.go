package wssconfig

type AppConfig struct {
	Http                     Http                     `mapstructure:"http"`
	BinanceFutureSecret      BinanceFutureSecret      `mapstructure:"binance-future-secret"`
	BinanceFutureUrl         BinanceFutureUrl         `mapstructure:"binance-future-url"`
	MarketStream             MarketStream             `mapstructure:"market-stream"`
	BinanceFutureServiceName BinanceFutureServiceName `mapstructure:"binance-future-service-name"`
}

type Http struct {
	Port string `mapstructure:"port"`
}

type BinanceFutureSecret struct {
	ApiKey    string `mapstructure:"api-key"`
	SecretKey string `mapstructure:"secret-key"`
}

type BinanceFutureUrl struct {
	BaseUrl string `mapstructure:"base-url"`
}

type MarketStream struct {
	KlineStream       string `mapstructure:"kline-stream"`
	MarketPriceStream string `mapstructure:"market-price-stream"`
}

type BinanceFutureServiceName struct {
	MarketStream string `mapstructure:"market-stream"`
}
