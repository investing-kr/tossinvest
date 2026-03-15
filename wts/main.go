package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/investing-kr/tossinvest"
	"github.com/spf13/cobra"
	"github.com/xtdlib/chttp"
	"github.com/xtdlib/rat"
)

var (
	wts    *tossinvest.Client
	ctx    = context.Background()
	output string
)

func main() {
	root := &cobra.Command{
		Use:   "wts",
		Short: "Toss Invest WTS CLI",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cdpaddr := os.Getenv("CHTTP_CDP_ADDR")
			if cdpaddr == "" {
				cdpaddr = "ws://localhost:9222"
			}
			wts = tossinvest.NewClient(chttp.NewClient(cdpaddr))
		},
	}

	root.PersistentFlags().StringVarP(&output, "output", "o", "", "output format (json)")
	root.AddCommand(buyCmd(), sellCmd(), hotCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func printJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(v)
}

func buyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "buy <symbol> <price> <quantity>",
		Short: "Buy a stock",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			order(tossinvest.TradeTypeBuy, args)
		},
	}
}

func sellCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sell <symbol> <price> <quantity>",
		Short: "Sell a stock",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			order(tossinvest.TradeTypeSell, args)
		},
	}
}

func order(side tossinvest.TradeType, args []string) {
	symbol := args[0]
	price, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid price: %s\n", args[1])
		os.Exit(1)
	}
	quantity, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid quantity: %s\n", args[2])
		os.Exit(1)
	}

	out, err := wts.V2WtsTradingOrder.CreateDirect(ctx, &tossinvest.OrderCreateDirectRequest{
		StockCode:          symbol,
		TradeType:          side,
		Quantity:           rat.Rat(quantity).Float64(),
		Price:              rat.Rat(price).Float64(),
		IsReservationOrder: false,
	})
	if err != nil {
		panic(err)
	}

	if output == "json" {
		printJSON(out)
		return
	}
	log.Printf("%#v", out)
}

func hotCmd() *cobra.Command {
	var tag string
	cmd := &cobra.Command{
		Use:   "hot",
		Short: "Show trading volume ranking",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := wts.V2DashboardWtsOverview.Ranking(ctx, &tossinvest.DashboardWtsOverviewRankingRequest{
				ID:      "biggest_total_amount",
				Filters: []string{},
				Tag:     tag,
			})
			if err != nil {
				panic(err)
			}
			if output == "json" {
				printJSON(resp)
				return
			}
			for _, p := range resp.Products {
				changeRate := (p.Price.Close - p.Price.Base) / p.Price.Base * 100
				fmt.Printf("%4s  %-16s  %s  %12.0f  %+.2f%%\n",
					fmt.Sprintf("#%d", p.Rank),
					p.ProductCode,
					padRight(p.Name, 30),
					p.Price.TossSecuritiesAmount,
					changeRate,
				)
			}
		},
	}
	cmd.Flags().StringVarP(&tag, "tag", "t", "us", "filter tag (all, us, kr)")
	return cmd
}

func displayWidth(s string) int {
	w := 0
	for _, r := range s {
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r) ||
			(r >= 0xFF01 && r <= 0xFF60) || (r >= 0xFFE0 && r <= 0xFFE6) {
			w += 2
		} else {
			w++
		}
	}
	return w
}

func padRight(s string, width int) string {
	pad := width - displayWidth(s)
	if pad <= 0 {
		return s
	}
	return s + strings.Repeat(" ", pad)
}
