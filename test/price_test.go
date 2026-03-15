package test

import (
	"log"
	"testing"
)

func TestTicks(t *testing.T) {
	stockinfo, err := wts.V2StockInfos.CodeOrSymbol(ctx, "AAPL")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := wts.V2StockPrices.Ticks(ctx, stockinfo.Code, 5)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Received %d ticks\n", len(resp))
	for i, tick := range resp {
		log.Printf("[%d] %s %s Price: %.2f Volume: %v TradeType: %s\n",
			i, tick.Time, tick.Code, tick.Price, tick.Volume, tick.TradeType)
	}
}
