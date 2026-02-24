package test

import (
	"context"
	"os"
	"testing"

	"github.com/investing-kr/tossinvest"
	"github.com/xtdlib/chttp"
)

var wts *tossinvest.Client
var ctx = context.Background()

func TestMain(m *testing.M) {
	// export CHTTP_CDP_ADDR=ws://localhost:9222
	wts = tossinvest.NewClient(chttp.NewClient(os.Getenv("CHTTP_CDP_ADDR")))
	code := m.Run()
	os.Exit(code)
}

// func TestBuy(t *testing.T) {
// 	t.SkipNow()
// 	resp, err := wts.V2WtsTradingOrder.CreateDirect(ctx, &tossinvest.BuyRequest{
// 		StockCode:          "AAPL",
// 		Price:              3.05,
// 		Quantity:           1,
// 		IsReservationOrder: false,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	log.Printf("%+v\n", resp)
// }

// func TestPending(t *testing.T) {
// 	resp, err := wts.V1TradingOrdersHistories.Pending(ctx, &tossinvest.OrderPendingRequest{
// 		StockCode:      "NAS0250624005",
// 		MarketDivision: "us",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	for _, order := range resp.Body {
// 		// spew.Dump(order)
// 		fmt.Println(order.OrderedDate, order.OrderNo, order.StockCode, order.StockName, order.TradeType, order.OrderUsdPrice, order.PendingQuantity)
// 	}
//
// 	// log.Printf("%+v\n", resp.Body)
// }
//
// func TestCompleted(t *testing.T) {
// 	resp, err := wts.V1TradingOrdersHistories.Completed(ctx, &tossinvest.OrderCompletedRequest{
// 		StockCode:      "NAS0250624005",
// 		MarketDivision: "us",
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	for _, order := range resp.Body {
// 		if !order.IsCanceled {
// 			fmt.Println(order.LastExecutedAt, order.StockCode, order.TradeType, order.OrderUsdPrice, order.Quantity)
// 		}
// 	}
// }
//
// func TestAllPending(t *testing.T) {
// 	resp, err := wts.V1TradingOrdersHistories.AllPending(ctx, &tossinvest.OrderAllPendingRequest{})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	for _, order := range resp {
// 		fmt.Println(order.OrderedDate, order.OrderNo, order.StockCode, order.StockName, order.TradeType, order.OrderUsdPrice, order.PendingQuantity)
// 	}
// }
//
// func TestOrderCancel(t *testing.T) {
// 	// t.SkipNow()
// 	resp, err := wts.V2WtsTradingOrder.Cancel(ctx, &tossinvest.OrderCancelRequest{
// 		OrderDate:          "2025-10-10",
// 		OrderNo:            1,
// 		IsAfterMarketOrder: false,
// 		Quantity:           200,
// 		StockCode:          "NAS0250624005",
// 		TradeType:          "buy",
// 		IsReservationOrder: false,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	log.Printf("Cancelled order: %d\n", resp.Value)
// }
