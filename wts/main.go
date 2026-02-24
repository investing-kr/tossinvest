package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/investing-kr/tossinvest"
	"github.com/xtdlib/chttp"
	"github.com/xtdlib/rat"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Fprintf(os.Stderr, "usage: wts <buy|sell> <symbol> <price> <quantity>\n")
		os.Exit(1)
	}

	side := os.Args[1]
	symbol := os.Args[2]
	price, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid price: %s\n", os.Args[3])
		os.Exit(1)
	}
	quantity, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid quantity: %s\n", os.Args[4])
		os.Exit(1)
	}
	if side != "buy" && side != "sell" {
		fmt.Fprintf(os.Stderr, "side must be 'buy' or 'sell'\n")
		os.Exit(1)
	}

	ctx := context.Background()
	var wts = tossinvest.NewClient(chttp.NewClient(""))
	stockinfo, err := wts.V2StockInfos.CodeOrSymbol(ctx, symbol)
	if err != nil {
		panic(err)
	}

	out, err := wts.V2WtsTradingOrder.CreateDirect(ctx, &tossinvest.OrderCreateDirectRequest{
		StockCode:          stockinfo.Code,
		TradeType:          tossinvest.TradeType(side),
		Quantity:           rat.Rat(quantity).Float64(),
		Price:              rat.Rat(price).Float64(),
		IsReservationOrder: false,
	})
	if err != nil {
		panic(err)
	}

	log.Printf("%#v", out)
}
