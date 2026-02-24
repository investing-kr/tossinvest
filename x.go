package tossinvest

import (
	"context"
	"fmt"
	"log"
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
	// orderReq.IsReservationOrder = true
	// orderReq.Market = "KSP"

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

	log.Println("Buy", stockCode, orderReq.Price, quantity)
	return c.V2WtsTradingOrder.CreateDirect(ctx, orderReq)
}

func (c *X) Sell(stockCode string, price any, quantity float64) (*OrderCreateDirectResponse, error) {
	// pricerat := rat.Rat(price)
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

	log.Println("Sell", stockCode, orderReq.Price, orderReq.Quantity)

	return c.V2WtsTradingOrder.CreateDirect(ctx, orderReq)
}

func (c *X) SellProduct(productCode string, price any, quantity float64) (*OrderCreateDirectResponse, error) {
	// pricerat := rat.Rat(price)
	orderReq := &OrderCreateDirectRequest{}
	orderReq.TradeType = "sell"
	orderReq.Price = roundPrice(price)
	orderReq.Quantity = quantity
	orderReq.StockCode = productCode

	log.Println("Sell", productCode, orderReq.Price, orderReq.Quantity)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return c.V2WtsTradingOrder.CreateDirect(ctx, orderReq)
}

func (c *X) BuyProduct(stockCode string, price any, quantity float64) (*OrderCreateDirectResponse, error) {
	orderReq := &OrderCreateDirectRequest{}
	orderReq.TradeType = "buy"

	orderReq.StockCode = stockCode
	orderReq.Price = roundPrice(price)
	orderReq.Quantity = quantity
	// orderReq.IsReservationOrder = true

	log.Println("Buy", stockCode, orderReq.Price, quantity)
	// pricerat := rat.Rat(price)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return c.V2WtsTradingOrder.CreateDirect(ctx, orderReq)
}

func (c *X) BuyProductDry(stockCode string, price any, quantity float64) (*OrderCreateDirectResponse, error) {
	orderReq := &OrderCreateDirectRequest{}
	orderReq.TradeType = "buy"

	orderReq.StockCode = stockCode
	orderReq.Price = roundPrice(price)
	orderReq.Quantity = quantity
	// orderReq.IsReservationOrder = true

	log.Println("Buy", stockCode, orderReq.Price, quantity)
	// pricerat := rat.Rat(price)
	return nil, nil
}

// func (c *X) SellProductDry(productCode string, price any, quantity float64) (*OrderCreateDirectResult, error) {
// 	// pricerat := rat.Rat(price)
// 	orderReq := &OrderCreateDirectRequest{}
// 	orderReq.TradeType = "sell"
// 	orderReq.Price = roundPrice(price)
// 	orderReq.Quantity = quantity
// 	orderReq.StockCode = productCode
//
// 	log.Println("Sell", productCode, orderReq.Price, orderReq.Quantity)
// 	return nil, nil
// }
