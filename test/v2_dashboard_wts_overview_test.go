package test

import (
	"log"
	"testing"

	"github.com/investing-kr/tossinvest"
)

func TestDashboardWtsOverviewRanking(t *testing.T) {
	resp, err := wts.V2DashboardWtsOverview.Ranking(ctx, &tossinvest.DashboardWtsOverviewRankingRequest{
		ID:      "biggest_total_amount",
		Filters: []string{},
		Tag:     "us",
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, p := range resp.Products {
		changeRate := (p.Price.Close - p.Price.Base) / p.Price.Base * 100
		log.Printf("#%d %s (%s) amount=%.0f %+.2f%%", p.Rank, p.Name, p.ProductCode, p.Price.TossSecuritiesAmount, changeRate)
	}
}
