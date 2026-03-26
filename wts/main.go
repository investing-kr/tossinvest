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

	"github.com/davecgh/go-spew/spew"
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
		Short: "wts.tossinvest.com cli",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cdpaddr := os.Getenv("CHTTP_CDP_ADDR")
			if cdpaddr == "" {
				cdpaddr = "ws://localhost:9222"
			}
			wts = tossinvest.NewClient(chttp.NewClient(cdpaddr))
		},
	}

	root.PersistentFlags().StringVarP(&output, "output", "o", "", "output format (json)")
	root.AddCommand(buyCmd(), sellCmd(), hotCmd(), candlesCmd(), bookCmd(), tickCmd(), infoCmd())

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
	var reserve bool
	cmd := &cobra.Command{
		Use:   "buy <symbol> <price> <quantity>",
		Short: "Buy a stock",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			order(tossinvest.TradeTypeBuy, args, reserve)
		},
	}
	cmd.Flags().BoolVar(&reserve, "reserve", false, "place as reservation order")
	return cmd
}

func infoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "info <symbol>",
		Short: "Show stock information",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			code := resolveCode(args[0])
			spew.Dump(code)
		},
	}
}

func sellCmd() *cobra.Command {
	var reserve bool
	cmd := &cobra.Command{
		Use:   "sell <symbol> <price> <quantity>",
		Short: "Sell a stock",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			order(tossinvest.TradeTypeSell, args, reserve)
		},
	}
	cmd.Flags().BoolVar(&reserve, "reserve", false, "place as reservation order")
	return cmd
}

func order(side tossinvest.TradeType, args []string, reserve bool) {
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
		IsReservationOrder: reserve,
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

func resolveCode(ticker string) *tossinvest.StockInfoResponse {
	info, err := wts.V2StockInfos.CodeOrSymbol(ctx, ticker)
	if err != nil {
		panic(err)
	}
	return info
}

func stockMarket(info *tossinvest.StockInfoResponse) string {
	if info.Currency == "KRW" {
		return "kr-s"
	}
	return "us-s"
}

func candlesCmd() *cobra.Command {
	var (
		interval string
		count    int
	)
	cmd := &cobra.Command{
		Use:   "candles <ticker>",
		Short: "Show candle chart data",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if !strings.Contains(interval, ":") {
				interval = interval + ":1"
			}
			info := resolveCode(args[0])
			resp, err := wts.V1CChart.Candles(ctx, stockMarket(info), info.Code, interval, count, true)
			if err != nil {
				panic(err)
			}
			if output == "json" {
				printJSON(resp)
				return
			}
			fmt.Printf("%s:%s:%s:%s\n", info.Symbol, info.Name, info.Code, interval)
			for _, c := range resp.Candles {
				change := (c.Close - c.Open) / c.Open * 100
				fmt.Printf("%s  O:%-10.2f  H:%-10.2f  L:%-10.2f  C:%-10.2f  V:%-10d  %+.2f%%\n",
					c.Dt, c.Open, c.High, c.Low, c.Close, c.Volume, change)
			}
		},
	}
	cmd.Flags().StringVarP(&interval, "interval", "i", "min:1", "interval (day:1, week:1, month:1, min:1, min:5)")
	cmd.Flags().IntVarP(&count, "count", "c", 20, "number of candles")
	return cmd
}

func bookCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "book <ticker>",
		Short: "Show orderbook",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			info := resolveCode(args[0])
			resp, err := wts.V3StockPrices.Quotes(ctx, info.Code)
			if err != nil {
				panic(err)
			}
			if output == "json" {
				printJSON(resp)
				return
			}

			fmt.Printf("%s:%s:%s  Close: %.2f\n", info.Symbol, info.Name, info.Code, resp.Close)
			fmt.Println()

			// find max volume for bar scaling
			maxVol := 0
			for _, v := range resp.SellVolumes {
				if v > maxVol {
					maxVol = v
				}
			}
			for _, v := range resp.BuyVolumes {
				if v > maxVol {
					maxVol = v
				}
			}

			barWidth := 30
			bar := func(vol int) string {
				if maxVol == 0 {
					return ""
				}
				n := vol * barWidth / maxVol
				return strings.Repeat("█", n)
			}

			// sell side (ask) - highest to lowest
			fmt.Printf("  %-10s  %8s  %s\n", "ASK", "VOL", "")
			for i := len(resp.SellPrices) - 1; i >= 0; i-- {
				fmt.Printf("  %-10.2f  %8d  %s\n", resp.SellPrices[i], resp.SellVolumes[i], bar(resp.SellVolumes[i]))
			}

			// spread line
			fmt.Printf("  %s\n", strings.Repeat("─", 50))

			// buy side (bid) - highest to lowest
			fmt.Printf("  %-10s  %8s  %s\n", "BID", "VOL", "")
			for i := 0; i < len(resp.BuyPrices); i++ {
				fmt.Printf("  %-10.2f  %8d  %s\n", resp.BuyPrices[i], resp.BuyVolumes[i], bar(resp.BuyVolumes[i]))
			}

			fmt.Println()
			fmt.Printf("  Total Ask: %d  Total Bid: %d\n", resp.SellVolume, resp.BuyVolume)
		},
	}
}

func tickCmd() *cobra.Command {
	var count int
	cmd := &cobra.Command{
		Use:   "tick <ticker>",
		Short: "Show recent ticks",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			info := resolveCode(args[0])
			resp, err := wts.V2StockPrices.Ticks(ctx, info.Code, count)
			if err != nil {
				panic(err)
			}
			if output == "json" {
				printJSON(resp)
				return
			}
			fmt.Printf("%s:%s:%s\n", info.Symbol, info.Name, info.Code)
			for _, t := range resp {
				change := (t.Price - t.Base) / t.Base * 100
				fmt.Printf("%s  %10.2f  %8.0f  %-4s  %+.2f%%\n",
					t.Time, t.Price, t.Volume, t.TradeType, change)
			}
		},
	}
	cmd.Flags().IntVarP(&count, "count", "c", 20, "number of ticks")
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
