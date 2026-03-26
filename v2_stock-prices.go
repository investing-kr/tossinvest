package tossinvest

import (
	"context"
	"fmt"
	"net/http"
)

type V2StockPrices service

type TicksResponseItem struct {
	Time             string  `json:"time"`
	Code             string  `json:"code"`
	Price            float64 `json:"price"`
	PriceKrw         float64 `json:"priceKrw"`
	Base             float64 `json:"base"`
	BaseKrw          float64 `json:"baseKrw"`
	Volume           float64 `json:"volume"`
	TradeType        string  `json:"tradeType"`
	CumulativeVolume float64 `json:"cumulativeVolume"`
}

type ticksResponse struct {
	Result []*TicksResponseItem `json:"result"`
}

func (c *V2StockPrices) Ticks(ctx context.Context, stockCode string, count int) ([]*TicksResponseItem, error) {
	url := fmt.Sprintf("%v/api/v2/stock-prices/%s/ticks?count=%d", c.BaseURL, stockCode, count)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	respBody := &ticksResponse{}
	err = c.getJson(httpReq, &respBody)
	return respBody.Result, err
}
