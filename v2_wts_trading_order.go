package tossinvest

import (
	"cmp"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"strconv"
	"time"
)

type OrderStatus string

const (
	OrderStatusCompleted OrderStatus = "체결완료"
	OrderStatusCanceled  OrderStatus = "취소"
	OrderStatusFailed    OrderStatus = "실패"
	OrderStatusReserved  OrderStatus = "예약"
	OrderStatusPartial   OrderStatus = "부분체결"
)

type V2WtsTradingOrder service
type V1TradingOrdersHistories service

type OrderCreateDirectResponse struct {
	Message    string `json:"message"`
	OrderDate  string `json:"orderDate"`
	OrderNo    int    `json:"orderNo"`
	IsReserved bool   `json:"isReserved"`
}

type orderCreateDirectResponse struct {
	Result OrderCreateDirectResponse `json:"result"`
}

type OrderCreateDirectRequest struct {
	StockCode              string    `json:"stockCode"`
	Market                 string    `json:"market"`
	CurrencyMode           string    `json:"currencyMode"`
	TradeType              TradeType `json:"tradeType"`
	Price                  float64   `json:"price"`
	Quantity               float64   `json:"quantity"`
	OrderAmount            int       `json:"orderAmount"`
	OrderPriceType         string    `json:"orderPriceType"`
	AgreedOver100Million   bool      `json:"agreedOver100Million"`
	MarginTrading          bool      `json:"marginTrading"`
	Max                    bool      `json:"max"`
	AllowAutoExchange      bool      `json:"allowAutoExchange"`
	IsReservationOrder     bool      `json:"isReservationOrder"`
	OpenPriceSinglePriceYn bool      `json:"openPriceSinglePriceYn"`
}

func (c *V2WtsTradingOrder) CreateDirect(ctx context.Context, orderReq *OrderCreateDirectRequest) (*OrderCreateDirectResponse, error) {
	if orderReq.OrderPriceType == "" {
		orderReq.OrderPriceType = "00"
	}

	// orderReq.IsReservationOrder = true

	if orderReq.TradeType == "sell" {
		slog.Debug("SELL", "stockCode", orderReq.StockCode, "price", orderReq.Price, "quantity", orderReq.Quantity)
	} else if orderReq.TradeType == "buy" {
		slog.Debug("BUY", "stockCode", orderReq.StockCode, "price", orderReq.Price, "quantity", orderReq.Quantity)
	} else {
		return nil, fmt.Errorf("unknown trade type: %s", orderReq.TradeType)
	}

	stockinfo, err := c.V2StockInfos.CodeOrSymbol(ctx, orderReq.StockCode)
	if err != nil {
		return nil, err
	}
	orderReq.StockCode = stockinfo.Code

	orderReq.Market = stockinfo.Market.Code
	orderReq.CurrencyMode = stockinfo.Currency

	respBody := &orderCreateDirectResponse{}
	req, err := c.newRequest(ctx, http.MethodPost, "/api/v2/wts/trading/order/create/direct", orderReq)
	if err != nil {
		return nil, err
	}

	err = c.postJson(req, &respBody)
	return &respBody.Result, err
}

func (c *V2WtsTradingOrder) Buy(ctx context.Context, orderReq *OrderCreateDirectRequest) (*OrderCreateDirectResponse, error) {
	orderReq.TradeType = "buy"
	return c.CreateDirect(ctx, orderReq)
}

func (c *V2WtsTradingOrder) Sell(ctx context.Context, orderReq *OrderCreateDirectRequest) (*OrderCreateDirectResponse, error) {
	orderReq.TradeType = "sell"
	return c.CreateDirect(ctx, orderReq)
}

type OrderPendingResponseResponseBody struct {
	StockCode            string      `json:"stockCode"`
	StockName            string      `json:"stockName"`
	Symbol               string      `json:"symbol"`
	LogoImageUrl         string      `json:"logoImageUrl"`
	OrderedAt            string      `json:"orderedAt"`
	LastExecutedAt       string      `json:"lastExecutedAt"`
	OrderedDate          string      `json:"orderedDate"`
	DisplayDate          string      `json:"displayDate"`
	OrderNo              int         `json:"orderNo"`
	OrderId              string      `json:"orderId"`
	TradeType            string      `json:"tradeType"`
	OrderPrice           int         `json:"orderPrice"`
	OrderUsdPrice        float64     `json:"orderUsdPrice"`
	OrderPriceType       string      `json:"orderPriceType"`
	OrderPriceTypeCode   string      `json:"orderPriceTypeCode"`
	OrderType            string      `json:"orderType"`
	ExecutionPrice       int         `json:"executionPrice"`
	ExecutionUsdPrice    float64     `json:"executionUsdPrice"`
	Quantity             int         `json:"quantity"`
	PendingQuantity      int         `json:"pendingQuantity"`
	OrderAmount          int         `json:"orderAmount"`
	UsdOrderAmount       float64     `json:"usdOrderAmount"`
	IsCanceled           bool        `json:"isCanceled"`
	IsFractionalOrder    bool        `json:"isFractionalOrder"`
	IsAfterMarketOrder   bool        `json:"isAfterMarketOrder"`
	Status               OrderStatus `json:"status"`
	CancelInProgress     bool        `json:"cancelInProgress"`
	CorrectSupport       bool        `json:"correctSupport"`
	CorrectionInProgress bool        `json:"correctionInProgress"`
	ForcedLiquidation    bool        `json:"forcedLiquidation"`
	OrderReasonType      string      `json:"orderReasonType"`
	Version              string      `json:"version"`
}

type OrderPendingResponse struct {
	PagingParam struct {
		Number    int    `json:"number"`
		Size      int    `json:"size"`
		Key       string `json:"key"`
		StockCode string `json:"stockCode"`
		Type      string `json:"type"`
	} `json:"pagingParam"`
	Body     []OrderPendingResponseResponseBody `json:"body"`
	LastPage bool                               `json:"lastPage"`
	Range    any                                `json:"range"`
}

type orderPendingResponse struct {
	Result OrderPendingResponse `json:"result"`
}

type OrderPendingRequest struct {
	StockCode      string
	Number         int
	Size           int
	MarketDivision string
}

func (c *V1TradingOrdersHistories) Pending(ctx context.Context, req *OrderPendingRequest) (*OrderPendingResponse, error) {
	respBody := &orderPendingResponse{}
	httpReq, err := c.newRequest(ctx, http.MethodGet, "/api/v1/trading/orders/histories/PENDING", nil)
	if err != nil {
		return nil, err
	}

	number := req.Number
	if number == 0 {
		number = 1
	}

	size := req.Size
	if size == 0 {
		size = 30
	}

	q := httpReq.URL.Query()
	q.Add("stockCode", req.StockCode)
	q.Add("number", strconv.Itoa(number))
	q.Add("size", strconv.Itoa(size))
	q.Add("marketDivision", req.MarketDivision)
	httpReq.URL.RawQuery = q.Encode()

	err = c.postJson(httpReq, &respBody)
	return &respBody.Result, err
}

type TradeType string

func (t TradeType) IsSell() bool {
	switch t {
	case TradeTypeSell:
		return true
	case TradeTypeBuy:
		return false
	default:
		panic("unknown trade type: " + string(t))
	}
}

func (t TradeType) IsBuy() bool {
	switch t {
	case TradeTypeSell:
		return false
	case TradeTypeBuy:
		return true
	default:
		panic("unknown trade type: " + string(t))
	}
}

const (
	TradeTypeBuy  TradeType = "buy"
	TradeTypeSell TradeType = "sell"
)

type OrderCompletedResponseResponseBody struct {
	StockCode string `json:"stockCode"`
	// StockName            string         `json:"stockName"`
	// Symbol               string         `json:"symbol"`
	// LogoImageUrl         any         `json:"logoImageUrl"`
	OrderedAt      string `json:"orderedAt"`
	LastExecutedAt string `json:"lastExecutedAt"`
	OrderedDate    string `json:"orderedDate"`
	DisplayDate    string `json:"displayDate"`
	OrderNo        int    `json:"orderNo"`
	// OrderId              string      `json:"orderId"`
	TradeType            TradeType   `json:"tradeType"`
	OrderPrice           float64     `json:"orderPrice"`
	OrderUsdPrice        float64     `json:"orderUsdPrice"`
	OrderPriceType       string      `json:"orderPriceType"`
	OrderPriceTypeCode   string      `json:"orderPriceTypeCode"`
	OrderType            string      `json:"orderType"`
	ExecutionPrice       float64     `json:"executionPrice"`
	ExecutionUsdPrice    float64     `json:"executionUsdPrice"`
	Quantity             float64     `json:"quantity"`
	PendingQuantity      float64     `json:"pendingQuantity"`
	OrderAmount          int         `json:"orderAmount"`
	UsdOrderAmount       float64     `json:"usdOrderAmount"`
	IsCanceled           bool        `json:"isCanceled"`
	IsFractionalOrder    bool        `json:"isFractionalOrder"`
	IsAfterMarketOrder   bool        `json:"isAfterMarketOrder"`
	Status               OrderStatus `json:"status"`
	CancelInProgress     bool        `json:"cancelInProgress"`
	CorrectSupport       bool        `json:"correctSupport"`
	CorrectionInProgress bool        `json:"correctionInProgress"`
	ForcedLiquidation    bool        `json:"forcedLiquidation"`
	OrderReasonType      any         `json:"orderReasonType"`
	Version              string      `json:"version"`
}

type OrderCompletedResponse struct {
	PagingParam struct {
		Number int    `json:"number"`
		Size   int    `json:"size"`
		Key    string `json:"key"`
	} `json:"pagingParam"`
	Body     []OrderCompletedResponseResponseBody `json:"body"`
	LastPage bool                                 `json:"lastPage"`
	Range    any                                  `json:"range"`
}

type orderCompletedResponse struct {
	Result OrderCompletedResponse `json:"result"`
}

type OrderCompletedRequest struct {
	StockCode      string
	Number         int
	Size           int
	Key            string
	MarketDivision string
}

func (c *V1TradingOrdersHistories) CompletedIn12h(ctx context.Context, req *OrderCompletedRequest) (*OrderCompletedResponse, error) {
	var allBody []OrderCompletedResponseResponseBody
	req.Number = 0

	for {
		result, err := c.Completed(ctx, req)
		if err != nil {
			return nil, err
		}

		req.Key = result.PagingParam.Key
		// log.Println(req.Key)
		// log.Println(result.Body[0])

		validBody := make([]OrderCompletedResponseResponseBody, 0, len(result.Body))
		_ = validBody
		stop := false
		_ = stop

		for _, body := range result.Body {
			// 2026-02-13T13:16:57.881+09:00
			lastExecuted, err := time.Parse(time.RFC3339, body.LastExecutedAt)
			if err != nil {
				return nil, err
			}
			if lastExecuted.After(time.Now().Add(-12 * time.Hour)) {
				validBody = append(validBody, body)
			} else {
				stop = true
			}
		}

		// log.Println("ByDateCompletedAll - fetched page:", result.PagingParam.Number, "items:", len(result.Body))
		allBody = append(allBody, validBody...)

		// order allBody by LastExecutedAt descending
		slices.SortFunc(allBody, func(a, b OrderCompletedResponseResponseBody) int {
			// return -cmp.Compare(a.LastExecutedAt, b.LastExecutedAt)
			return -cmp.Compare(a.LastExecutedAt, b.LastExecutedAt)
		})

		if stop || result.LastPage {
			result.Body = allBody
			return result, nil
		}
		// time.Sleep(time.Second)
	}
}

func (c *V1TradingOrdersHistories) CompletedAll(ctx context.Context, req *OrderCompletedRequest) (*OrderCompletedResponse, error) {
	var allBody []OrderCompletedResponseResponseBody
	req.Number = 0

	for {
		result, err := c.Completed(ctx, req)
		if err != nil {
			return nil, err
		}

		req.Key = result.PagingParam.Key
		// log.Println(req.Key)
		// log.Println(result.Body[0])

		// log.Println("ByDateCompletedAll - fetched page:", result.PagingParam.Number, "items:", len(result.Body))
		allBody = append(allBody, result.Body...)

		// order allBody by LastExecutedAt descending
		slices.SortFunc(allBody, func(a, b OrderCompletedResponseResponseBody) int {
			// return -cmp.Compare(a.LastExecutedAt, b.LastExecutedAt)
			return -cmp.Compare(a.LastExecutedAt, b.LastExecutedAt)
		})

		if result.LastPage {
			result.Body = allBody
			return result, nil
		}
		// time.Sleep(time.Second)
	}
}

func (c *V1TradingOrdersHistories) Completed(ctx context.Context, req *OrderCompletedRequest) (*OrderCompletedResponse, error) {
	respBody := &orderCompletedResponse{}
	httpReq, err := c.newRequest(ctx, http.MethodGet, "/api/v1/trading/orders/histories/COMPLETED", nil)
	if err != nil {
		return nil, err
	}

	number := req.Number
	if number == 0 {
		number = 1
	}

	size := req.Size
	if size == 0 {
		size = 100
	}

	q := httpReq.URL.Query()
	q.Add("stockCode", req.StockCode)
	q.Add("number", strconv.Itoa(number))
	q.Add("size", strconv.Itoa(size))
	q.Add("marketDivision", req.MarketDivision)

	if req.Key != "" {
		q.Add("key", req.Key)
	}

	httpReq.URL.RawQuery = q.Encode()

	err = c.getJson(httpReq, &respBody)

	body := respBody.Result.Body
	// sort body array
	slices.SortFunc(body, func(a, b OrderCompletedResponseResponseBody) int {
		// return -cmp.Compare(a.LastExecutedAt, b.LastExecutedAt)
		return -cmp.Compare(OrderID(a.OrderedDate, a.OrderNo), OrderID(b.OrderedDate, b.OrderNo))
	})

	// // slice filter if a.LastExecutedAt is empty string
	// body = slices.DeleteFunc(body, func(a OrderCompletedResponseResultBody) bool {
	// 	return a.LastExecutedAt == ""
	// })

	respBody.Result.Body = body

	return &respBody.Result, err
}

type cancelPrepareRequest struct {
	IsAfterMarketOrder bool    `json:"isAfterMarketOrder"`
	Quantity           float64 `json:"quantity"`
	StockCode          string  `json:"stockCode"`
	TradeType          string  `json:"tradeType"`
	WithOrderKey       bool    `json:"withOrderKey"`
	IsReservationOrder bool    `json:"isReservationOrder"`
}

type CancelPrepareResponse struct {
	DelayCancelExchange bool   `json:"delayCancelExchange"`
	From                any    `json:"from"`
	To                  any    `json:"to"`
	OrderKey            string `json:"orderKey"`
	AuthRequired        struct {
		Required    bool `json:"required"`
		SimpleTrade bool `json:"simpleTrade"`
		Verifier    any  `json:"verifier"`
	} `json:"authRequired"`
}

type cancelPrepareResponse struct {
	Result CancelPrepareResponse `json:"result"`
}

func (c *V2WtsTradingOrder) CancelPrepare(ctx context.Context, orderDate string, orderNo int, req *cancelPrepareRequest) (*CancelPrepareResponse, error) {
	respBody := &cancelPrepareResponse{}
	path := "/api/v2/wts/trading/order/cancel/prepare/" + orderDate + "/" + strconv.Itoa(orderNo)

	req.WithOrderKey = true

	httpReq, err := c.newRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}

	err = c.postJson(httpReq, &respBody)
	return &respBody.Result, err
}

type OrderCancelRequest struct {
	OrderDate          string  `json:"-"`
	OrderNo            int     `json:"-"`
	IsAfterMarketOrder bool    `json:"isAfterMarketOrder"`
	Quantity           float64 `json:"quantity"`
	StockCode          string  `json:"stockCode"`
	TradeType          string  `json:"tradeType"`
	IsReservationOrder bool    `json:"isReservationOrder"`
}

type OrderCancelResponse struct {
	Value int `json:"result"`
}

type orderCancelResponse struct {
	Result int `json:"result"`
}

func (c *V2WtsTradingOrder) Cancel(ctx context.Context, req *OrderCancelRequest) (*OrderCancelResponse, error) {
	if req.OrderDate == "" {
		req.OrderDate = time.Now().Format("20060102")
	}

	prepareReq := &cancelPrepareRequest{
		IsAfterMarketOrder: req.IsAfterMarketOrder,
		Quantity:           req.Quantity,
		StockCode:          req.StockCode,
		TradeType:          req.TradeType,
		IsReservationOrder: req.IsReservationOrder,
	}

	prepareResp, err := c.CancelPrepare(ctx, req.OrderDate, req.OrderNo, prepareReq)
	if err != nil {
		return nil, err
	}

	respBody := &orderCancelResponse{}
	path := "/api/v2/wts/trading/order/cancel/" + req.OrderDate + "/" + strconv.Itoa(req.OrderNo)
	httpReq, err := c.newRequest(ctx, http.MethodPost, path, req)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("x-order-key", prepareResp.OrderKey)
	err = c.postJson(httpReq, &respBody)
	return &OrderCancelResponse{Value: respBody.Result}, err
}

type OrderAllPendingRequest struct {
	Number int
	Size   int
}

type OrderAllPendingResponse struct {
	Result []OrderPendingResponseResponseBody `json:"result"`
}

func (c *V1TradingOrdersHistories) AllPending(ctx context.Context, req *OrderAllPendingRequest) ([]OrderPendingResponseResponseBody, error) {
	respBody := &OrderAllPendingResponse{}
	httpReq, err := c.newRequest(ctx, http.MethodGet, "/api/v1/trading/orders/histories/all/pending", nil)
	if err != nil {
		return nil, err
	}

	number := req.Number
	if number == 0 {
		number = 1
	}

	size := req.Size
	if size == 0 {
		size = 30
	}

	q := httpReq.URL.Query()
	q.Add("number", strconv.Itoa(number))
	q.Add("size", strconv.Itoa(size))
	httpReq.URL.RawQuery = q.Encode()

	err = c.getJson(httpReq, &respBody)
	return respBody.Result, err
}

func OrderID(orderDate string, orderNo int) string {
	return orderDate + fmt.Sprintf(":%08d", orderNo)
}
