package tossinvest

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/xtdlib/rat"
)

type X service

func roundPrice(price any) float64 {
	ratPrice := rat.Rat(price)
	if ratPrice.LessOrEqual("1") {
		ratPrice = ratPrice.SetPrecision(4)
	} else {
		ratPrice = ratPrice.SetPrecision(2)
	}
	p, err := strconv.ParseFloat(ratPrice.DecimalString(), 64)
	if err != nil {
		panic(err)
	}
	return p
}

func (c *X) Buy(stockCode string, price any, quantity float64) (*OrderCreateDirectResponse, error) {
	orderReq := &OrderCreateDirectRequest{}
	orderReq.TradeType = "buy"
	orderReq.Quantity = quantity

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stockinfo, err := c.V2StockInfos.CodeOrSymbol(ctx, stockCode)
	if err != nil {
		return nil, err
	}
	orderReq.StockCode = stockinfo.Code

	orderReq.Market = stockinfo.Market.Code
	orderReq.CurrencyMode = stockinfo.Currency

	switch stockinfo.Currency {
	case "USD":
		orderReq.Price = roundPrice(price)
	case "KRW":
		orderReq.Price = float64(rat.Rat(price).Int())
	default:
		return nil, fmt.Errorf("unsupported currency: %s", stockinfo.Currency)
	}

	return c.V2WtsTradingOrder.CreateDirect(ctx, orderReq)
}

func (c *X) Sell(stockCode string, price any, quantity float64) (*OrderCreateDirectResponse, error) {
	orderReq := &OrderCreateDirectRequest{}
	orderReq.TradeType = "sell"
	orderReq.Quantity = quantity

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stockinfo, err := c.V2StockInfos.CodeOrSymbol(ctx, stockCode)
	if err != nil {
		return nil, err
	}
	orderReq.StockCode = stockinfo.Code

	orderReq.Market = stockinfo.Market.Code
	orderReq.CurrencyMode = stockinfo.Currency

	switch stockinfo.Currency {
	case "USD":
		orderReq.Price = roundPrice(price)
	case "KRW":
		orderReq.Price = float64(rat.Rat(price).Int())
	default:
		return nil, fmt.Errorf("unsupported currency: %s", stockinfo.Currency)
	}

	return c.V2WtsTradingOrder.CreateDirect(ctx, orderReq)
}

func (c *X) SellProduct(productCode string, price any, quantity float64) (*OrderCreateDirectResponse, error) {
	orderReq := &OrderCreateDirectRequest{}
	orderReq.TradeType = "sell"
	orderReq.Price = roundPrice(price)
	orderReq.Quantity = quantity
	orderReq.StockCode = productCode

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return c.V2WtsTradingOrder.CreateDirect(ctx, orderReq)
}
