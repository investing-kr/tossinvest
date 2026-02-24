package test

import (
	"log"
	"testing"
)

func TestCandles(t *testing.T) {
	resp, err := wts.V1CChart.Candles(ctx, "us-s", "NAS0250224006", "day:1", 1, true)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("Code: %s ExchangeRate: %.2f", resp.Code, resp.ExchangeRate)
	for i, c := range resp.Candles {
		log.Printf("[%d] %s O:%.2f H:%.2f L:%.2f C:%.2f V:%d", i, c.Dt, c.Open, c.High, c.Low, c.Close, c.Volume)
	}
}
