package tossinvest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type service struct {
	*Client
}

type Client struct {
	HttpClient *http.Client
	BaseURL    string

	V2WtsTradingOrder        *V2WtsTradingOrder
	V1TradingOrdersHistories *V1TradingOrdersHistories
	V1DashboardAsset         *V1DashboardAsset
	V2StockPrices            *V2StockPrices
	V3StockPrices            *V3StockPrices
	V2StockInfos             *V2StockInfos
	V3SearchAll              *V3SearchAll
	V2TradingMyOrders        *V2TradingMyOrders
	V1CChart                 *V1CChart
	X                        *X
}

type ClientOption func(*Client)

func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.BaseURL = baseURL
	}
}

func NewClient(httpClient *http.Client, opts ...ClientOption) *Client {
	c := &Client{
		HttpClient: httpClient,
		BaseURL:    "https://wts-cert-api.tossinvest.com",
	}

	c.V2WtsTradingOrder = &V2WtsTradingOrder{c}
	c.V1TradingOrdersHistories = &V1TradingOrdersHistories{c}
	c.V1DashboardAsset = &V1DashboardAsset{c}
	c.V2StockPrices = &V2StockPrices{c}
	c.V3StockPrices = &V3StockPrices{c}
	c.V2StockInfos = &V2StockInfos{Client: c}
	c.V3SearchAll = &V3SearchAll{c}
	c.V2TradingMyOrders = &V2TradingMyOrders{c}
	c.V1CChart = &V1CChart{c}
	c.X = &X{c}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) doJson(req *http.Request, respBody interface{}) error {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodystring, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("tossinvest: failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		errResp := &ErrorResponse{}
		err = json.Unmarshal(bodystring, &errResp)
		if err != nil {
			return fmt.Errorf("tossinvest: %w: %s", err, bodystring)
		}
		return errResp
	}

	err = json.Unmarshal(bodystring, &respBody)
	if err != nil {
		return fmt.Errorf("tossinvest: %w: %s", err, bodystring)
	}

	return nil
}

func (c *Client) postJson(req *http.Request, respBody interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	return c.doJson(req, respBody)
}

func (c *Client) getJson(req *http.Request, respBody interface{}) error {
	return c.doJson(req, respBody)
}

func (c *Client) newRequest(ctx context.Context, method string, url string, jsonbody interface{}) (*http.Request, error) {
	var reqbody *bytes.Reader
	if jsonbody != nil {
		payloadBytes, err := json.Marshal(jsonbody)
		if err != nil {
			return nil, err
		}
		reqbody = bytes.NewReader(payloadBytes)
	}
	url = c.BaseURL + url

	if ctx == nil {
		ctx = context.Background()
	}

	// it panics if reqbody is nil
	if reqbody != nil {
		return http.NewRequestWithContext(ctx, method, url, reqbody)
	} else {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}
}
