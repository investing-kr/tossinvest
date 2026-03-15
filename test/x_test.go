package test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestXBuy(t *testing.T) {
	resp, err := wts.X.Buy("AAPL", 50, 1)
	if err != nil {
		t.Fatal(err)
	}

	spew.Dump(resp)
	// (*tossinvest.OrderCreateDirectResponse)(0xc00058c930)({
	//  Message: (string) (len=27) "애플 구매 주문 완료",
	//  OrderDate: (string) (len=10) "2026-02-25",
	//  OrderNo: (int) 21,
	//  IsReserved: (bool) false
	// })
}
