package tossinvest

import (
	"context"
	"fmt"
	"net/http"
)

type V3StockPrices service

type QuotesResponse struct {
	Close         float64   `json:"close"`
	CloseKrw      int       `json:"closeKrw"`
	SellPrices    []float64 `json:"offerPrices"` // sell price
	SellPricesKrw []int     `json:"offerPricesKrw"`
	SellVolumes   []int     `json:"offerVolumes"`
	SellVolume    int       `json:"offerVolume"`

	BuyPrices    []float64 `json:"bidPrices"` // buy price
	BuyPricesKrw []int     `json:"bidPricesKrw"`
	BuyVolumes   []int     `json:"bidVolumes"`
	BuyVolume    int       `json:"bidVolume"`
}

type quotesResponse struct {
	Result QuotesResponse `json:"result"`
}

func (c *V3StockPrices) Quotes(ctx context.Context, stockCode string) (*QuotesResponse, error) {
	url := fmt.Sprintf("https://wts-info-api-dc3.tossinvest.com/api/v3/stock-prices/%s/quotes", stockCode)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	respBody := &quotesResponse{}
	err = c.getJson(httpReq, &respBody)
	return &respBody.Result, err
}
