package test

import (
	"log"
	"testing"

	"github.com/investing-kr/tossinvest"
)

func TestCreateDirect(t *testing.T) {
	// out, err := wts.V2WtsTradingOrder.CreateDirect(ctx, &tossinvest.OrderCreateDirectRequest{
	// 	StockCode: .StockCode,
	// 	// Market:                 "usd",
	// 	// CurrencyMode:           "",
	// 	TradeType: order.TradeType,
	// 	Price:     order.OrderPrice.Usd,
	// 	Quantity:  order.OrderQuantity,
	// })
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Print(out)
	// time.Sleep(time.Second)
}

func TestMyOrdersByDateCompleted(t *testing.T) {
	resp, err := wts.V2TradingMyOrders.ByDateCompletedAll(ctx, &tossinvest.MyOrdersByDateCompletedRequest{
		Market:    "us",
		RangeFrom: "2026-02-04",
		RangeTo:   "2026-02-04",
		Size:      20,
		Number:    1,
	})
	_ = resp
	if err != nil {
		t.Fatal(err)
	}

	log.Println("Total Count:", len(resp.Body))

}
