package tossinvest

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type V2TradingMyOrders service

type MyOrdersByDateCompletedRequest struct {
	Market    string
	RangeFrom string
	RangeTo   string
	Size      int
	Number    int
	Key       string
}

type MyOrdersByDateCompletedResponseBody struct {
	OrderedAt     string    `json:"orderedAt"`
	OrderNo       int       `json:"orderNo"`
	OrderId       any       `json:"orderId"`
	RootOrderNo   int       `json:"rootOrderNo"`
	StockCode     string    `json:"stockCode"`
	StockName     string    `json:"stockName"`
	Symbol        string    `json:"symbol"`
	TradeType     TradeType `json:"tradeType"`
	Option        any       `json:"option"`
	Bond          any       `json:"bond"`
	OrderQuantity float64   `json:"orderQuantity"`
	OrderPrice    struct {
		Krw float64 `json:"krw"`
		Usd float64 `json:"usd"`
	} `json:"orderPrice"`
	OrderAmount struct {
		Krw int     `json:"krw"`
		Usd float64 `json:"usd"`
	} `json:"orderAmount"`
	OrderPriceType        string  `json:"orderPriceType"`
	OrderQuantityType     string  `json:"orderQuantityType"`
	Fractional            bool    `json:"fractional"`
	ExecutedQuantity      float64 `json:"executedQuantity"`
	AverageExecutionPrice struct {
		Krw int     `json:"krw"`
		Usd float64 `json:"usd"`
	} `json:"averageExecutionPrice"`
	SumExecutedAmount      float64     `json:"sumExecutedAmount"`
	LastExecutedAt         string      `json:"lastExecutedAt"`
	CancelType             any         `json:"cancelType"`
	ShortSelling           bool        `json:"shortSelling"`
	UserOrderDate          string      `json:"userOrderDate"`
	AfterMarketOrder       bool        `json:"afterMarketOrder"`
	OrderTransactionType   string      `json:"orderTransactionType"`
	OrderTransactionTypeV2 string      `json:"orderTransactionTypeV2"`
	Status                 OrderStatus `json:"status"`
	MarginTrading          bool        `json:"marginTrading"`
	ForcedLiquidation      bool        `json:"forcedLiquidation"`
	OrderReasonType        any         `json:"orderReasonType"`
	Version                string      `json:"version"`
}

type MyOrdersByDateCompletedResponse struct {
	PagingParam struct {
		Range struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"range"`
		StockCode any    `json:"stockCode"`
		Number    int    `json:"number"`
		Size      int    `json:"size"`
		Key       string `json:"key"`
	} `json:"pagingParam"`
	Body          []MyOrdersByDateCompletedResponseBody `json:"body"`
	LastPage      bool                                  `json:"lastPage"`
	ApplyAllAsset bool                                  `json:"applyAllAsset"`
}

type myOrdersByDateCompletedResponse struct {
	Result MyOrdersByDateCompletedResponse `json:"result"`
}

func (c *V2TradingMyOrders) ByDateCompleted(ctx context.Context, req *MyOrdersByDateCompletedRequest) (*MyOrdersByDateCompletedResponse, error) {
	respBody := &myOrdersByDateCompletedResponse{}
	path := "/api/v2/trading/my-orders/markets/" + req.Market + "/by-date/completed"
	httpReq, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	q := httpReq.URL.Query()
	q.Add("range.from", req.RangeFrom)
	q.Add("range.to", req.RangeTo)
	if req.Size > 0 {
		q.Add("size", strconv.Itoa(req.Size))
	}

	if req.Key != "" {
		q.Add("key", req.Key)
	}

	q.Add("number", "1")
	// if req.Number > 0 {
	// 	q.Add("number", strconv.Itoa(req.Number))
	// }

	q.Add("applyAllAsset", "false")
	httpReq.URL.RawQuery = q.Encode()

	err = c.getJson(httpReq, &respBody)

	return &respBody.Result, err
}

func (c *V2TradingMyOrders) ByDateCompletedAll(ctx context.Context, req *MyOrdersByDateCompletedRequest) (*MyOrdersByDateCompletedResponse, error) {
	var allBody []MyOrdersByDateCompletedResponseBody
	maxDepth := 100

	var err error
	var result *MyOrdersByDateCompletedResponse

	req.Number = 0
	for i := 0; i < maxDepth; i++ {
		result, err = c.ByDateCompleted(ctx, req)
		if err != nil {
			return nil, err
		}

		req.Key = result.PagingParam.Key

		allBody = append(allBody, result.Body...)
		if result.LastPage {
			result.Body = allBody
			return result, nil
		}
	}

	return result, fmt.Errorf("tossinvest: too many pages, maybe more than %d", maxDepth)
}
