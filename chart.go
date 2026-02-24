package tossinvest

import (
	"context"
	"fmt"
	"net/http"
)

type V1CChart service

type Candle struct {
	Dt     string  `json:"dt"`
	Base   float64 `json:"base"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int     `json:"volume"`
	Amount int64   `json:"amount"`
}

type CandlesResponse struct {
	Code         string   `json:"code"`
	NextDateTime string   `json:"nextDateTime"`
	ExchangeRate float64  `json:"exchangeRate"`
	Candles      []Candle `json:"candles"`
}

type candlesResponse struct {
	Result CandlesResponse `json:"result"`
}

func (c *V1CChart) Candles(ctx context.Context, market string, code string, interval string, count int, useAdjustedRate bool) (*CandlesResponse, error) {
	url := fmt.Sprintf("https://wts-info-api-dc3.tossinvest.com/api/v1/c-chart/%s/%s/%s?count=%d&useAdjustedRate=%t", market, code, interval, count, useAdjustedRate)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	respBody := &candlesResponse{}
	err = c.getJson(httpReq, &respBody)
	return &respBody.Result, err
}
