package test

import (
	"testing"
	"time"

	"github.com/investing-kr/tossinvest"
)

func TestCompleted(t *testing.T) {
	completed, err := wts.V2TradingMyOrders.ByDateCompleted(ctx, &tossinvest.MyOrdersByDateCompletedRequest{
		Market:    "us",
		RangeFrom: time.Now().Add(-time.Hour * 24).Format("2006-01-02"),
		RangeTo:   time.Now().Format("2006-01-02"),
		Size:      100,
	})
	_ = completed
	if err != nil {
		t.Fatal(err)
	}

	for _, order := range completed.Body {

		t.Log(order.OrderNo, order.Version, "executed", order.LastExecutedAt, order.StockName, order.TradeType)
	}
}
