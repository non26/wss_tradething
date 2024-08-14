package marketstreammodel

import (
	"errors"
	"strings"
)

type KlineRequest struct {
	Interval string
	Symbol   string
}

func NewKlineInterval() map[string]bool {
	m := map[string]bool{
		"1m":  true,
		"3m":  true,
		"5m":  true,
		"15m": true,
		"30m": true,
		"1h":  true,
		"2h":  true,
		"4h":  true,
		"6h":  true,
		"8h":  true,
		"12h": true,
		"1d":  true,
		"3d":  true,
		"1w":  true,
		"1M":  true,
	}
	return m
}

func (k *KlineRequest) Validate() error {
	if k.Symbol == "" {
		return errors.New("validate symbol error")
	}
	k.Symbol = strings.ToLower(k.Symbol)
	if k.Interval == "" {
		return errors.New("validate interval error")
	}
	kline := NewKlineInterval()
	_, ok := kline[k.Interval]
	if !ok {
		return errors.New("validate interval error")
	}
	return nil
}
