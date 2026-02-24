package test

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestGetProductCode(t *testing.T) {
	tests := []struct {
		ticker       string
		expectedCode string
	}{
		{"AAPL", "US19801212001"},
		{"RGTZ", "NAS0251009002"},
		{"MSFT", "US19860313001"},
		{"GOOGL", "US20040819002"},
		{"AMZN", "US19970515001"},
		{"A005930", "A005930"},
	}

	for _, tt := range tests {
		ticker := tt.ticker
		expectedCode := tt.expectedCode

		ctx := context.Background()
		t.Run(ticker, func(t *testing.T) {
			stockinfo, err := wts.V2StockInfos.CodeOrSymbol(ctx, ticker)
			if err != nil {
				t.Fatal(err)
			}
			if stockinfo.Code != expectedCode {
				t.Fatalf("expected %s, got %s", expectedCode, stockinfo.Code)
			}
		})
	}
}

func TestV2StockInfos_CodeOrSymbol(t *testing.T) {
	ctx := context.Background()

	{
		stockinfo, err := wts.V2StockInfos.CodeOrSymbol(ctx, "A005930")
		if err != nil {
			t.Fatal(err)
		}
		spew.Dump(stockinfo)
		// (*tossinvest.StockInfo)(0xc0000eca80)({
		//  Code: (string) (len=7) "A005930",
		//  GUID: (string) (len=12) "KR7005930003",
		//  Symbol: (string) (len=6) "005930",
		//  ISINCode: (string) (len=12) "KR7005930003",
		//  Status: (string) (len=1) "N",
		//  Name: (string) (len=12) "삼성전자",
		//  EnglishName: (string) (len=11) "SamsungElec",
		//  DetailName: (string) (len=12) "삼성전자",
		//  Market: (tossinvest.StockInfoMarket) {
		//   Code: (string) (len=3) "KSP",
		//   DisplayName: (string) (len=9) "코스피"
		//  },
		//  Group: (tossinvest.StockInfoGroup) {
		//   Code: (string) (len=2) "ST",
		//   DisplayName: (string) (len=6) "주권"
		//  },
		//  CompanyCode: (string) (len=6) "005930",
		//  CompanyName: (string) (len=12) "삼성전자",
		//  LogoImageUrl: (string) (len=67) "https://static.toss.im/png-icons/securities/icn-sec-fill-005930.png",
		//  Currency: (string) (len=3) "KRW",
		//  TradingSuspended: (bool) false,
		//  KrxTradingSuspended: (bool) false,
		//  NxtTradingSuspended: (bool) false,
		//  CommonShare: (bool) true,
		//  Spac: (bool) false,
		//  SpacMergerDate: (*string)(<nil>),
		//  LeverageFactor: (float64) 0,
		//  Clearance: (bool) false,
		//  RiskLevel: (string) (len=1) "1",
		//  PurchasePrerequisite: (string) (len=1) "0",
		//  SharesOutstanding: (int64) 5919637922,
		//  PrevListDate: (*string)(<nil>),
		//  ListDate: (*string)(0xc00031c1a0)((len=10) "1975-06-11"),
		//  DelistDate: (*string)(<nil>),
		//  OfferingPrice: (*float64)(<nil>),
		//  WarrantsCode: (*string)(<nil>),
		//  WarrantsGroupCode: (*string)(<nil>),
		//  EtfTaxCode: (*string)(<nil>),
		//  DaytimePriceSupported: (bool) false,
		//  OptionSupported: (bool) false,
		//  OptionPennyPilotPriceSupported: (bool) false,
		//  OptionOvertimeSupported: (bool) false,
		//  OptionInstrument: (*string)(<nil>),
		//  DerivativeEtp: (bool) false,
		//  PoolingStock: (bool) false,
		//  NxtSupported: (bool) true,
		//  UserTradingSuspended: (bool) false,
		//  LimitOnCompetitiveTradingVolume: (bool) false,
		//  NxtOpenDate: (*string)(0xc00031c1b0)((len=10) "2025-03-24"),
		//  NxtOpenDateRecent: (*string)(0xc00031c1c0)((len=10) "2025-03-24"),
		//  DerivativeEtf: (bool) false
		// })
	}

	{
		stockinfo, err := wts.V2StockInfos.CodeOrSymbol(ctx, "AAPL")
		if err != nil {
			t.Fatal(err)
		}
		spew.Dump(stockinfo)
		// (*tossinvest.StockInfo)(0xc00028efc0)({
		//  Code: (string) (len=13) "US19801212001",
		//  GUID: (string) (len=13) "US19801212001",
		//  Symbol: (string) (len=4) "AAPL",
		//  ISINCode: (string) (len=12) "US0378331005",
		//  Status: (string) (len=1) "N",
		//  Name: (string) (len=6) "애플",
		//  EnglishName: (string) (len=5) "Apple",
		//  DetailName: (string) (len=6) "애플",
		//  Market: (tossinvest.StockInfoMarket) {
		//   Code: (string) (len=3) "NSQ",
		//   DisplayName: (string) (len=6) "NASDAQ"
		//  },
		//  Group: (tossinvest.StockInfoGroup) {
		//   Code: (string) (len=2) "ST",
		//   DisplayName: (string) (len=6) "주권"
		//  },
		//  CompanyCode: (string) (len=12) "NAS000C7F-E0",
		//  CompanyName: (string) (len=6) "애플",
		//  LogoImageUrl: (string) (len=73) "https://static.toss.im/png-icons/securities/icn-sec-fill-NAS000C7F-E0.png",
		//  Currency: (string) (len=3) "USD",
		//  TradingSuspended: (bool) false,
		//  KrxTradingSuspended: (bool) false,
		//  NxtTradingSuspended: (bool) false,
		//  CommonShare: (bool) true,
		//  Spac: (bool) false,
		//  SpacMergerDate: (*string)(<nil>),
		//  LeverageFactor: (float64) 0,
		//  Clearance: (bool) false,
		//  RiskLevel: (string) (len=1) "1",
		//  PurchasePrerequisite: (string) (len=1) "0",
		//  SharesOutstanding: (int64) 14702703000,
		//  PrevListDate: (*string)(<nil>),
		//  ListDate: (*string)(0xc000487660)((len=10) "1980-12-12"),
		//  DelistDate: (*string)(<nil>),
		//  OfferingPrice: (*float64)(<nil>),
		//  WarrantsCode: (*string)(<nil>),
		//  WarrantsGroupCode: (*string)(<nil>),
		//  EtfTaxCode: (*string)(<nil>),
		//  DaytimePriceSupported: (bool) true,
		//  OptionSupported: (bool) true,
		//  OptionPennyPilotPriceSupported: (bool) true,
		//  OptionOvertimeSupported: (bool) false,
		//  OptionInstrument: (*string)(<nil>),
		//  DerivativeEtp: (bool) false,
		//  PoolingStock: (bool) false,
		//  NxtSupported: (bool) false,
		//  UserTradingSuspended: (bool) false,
		//  LimitOnCompetitiveTradingVolume: (bool) false,
		//  NxtOpenDate: (*string)(<nil>),
		//  NxtOpenDateRecent: (*string)(<nil>),
		//  DerivativeEtf: (bool) false
		// })
	}
}
