package test

import (
	"log"
	"testing"

	"github.com/investing-kr/tossinvest"
)

func TestDashboardAssetSectionsAll(t *testing.T) {
	resp, err := wts.V1DashboardAsset.SectionsAll(ctx, &tossinvest.DashboardAssetSectionsAllRequest{
		Types: []string{"OVERVIEW"},
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, section := range resp.Sections {
		for _, item := range section.Data.Kr.Items {
			log.Printf("%v %v\n", item.StockName, item.Quantity)
		}

		log.Println("-------")

		for _, item := range section.Data.Us.Items {
			log.Printf("%v %v\n", item.StockName, item.Quantity)
		}
	}
}
