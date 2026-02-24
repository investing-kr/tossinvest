package tossinvest

import (
	"context"
	"net/http"
	"time"
)

type V1DashboardAsset service

type DashboardAssetSectionsAllRequest struct {
	Types []string `json:"types"`
}

type DashboardAssetSectionsAllResponse struct {
	Sections []struct {
			Data struct {
				PrincipalAmount struct {
					Krw float64 `json:"krw"`
					Usd float64 `json:"usd"`
				} `json:"principalAmount"`
				EvaluatedAmount struct {
					Krw float64 `json:"krw"`
					Usd float64 `json:"usd"`
				} `json:"evaluatedAmount"`
				ProfitLossAmount struct {
					Krw float64 `json:"krw"`
					Usd float64 `json:"usd"`
				} `json:"profitLossAmount"`
				DailyProfitLossAmount struct {
					Krw float64 `json:"krw"`
					Usd float64 `json:"usd"`
				} `json:"dailyProfitLossAmount"`
				ProfitLossRate struct {
					Krw float64 `json:"krw"`
					Usd float64 `json:"usd"`
				} `json:"profitLossRate"`
				DailyProfitLossRate struct {
					Krw float64 `json:"krw"`
					Usd float64 `json:"usd"`
				} `json:"dailyProfitLossRate"`
				Kr struct {
					Items []struct {
						ID                int     `json:"id"`
						Key               string  `json:"key"`
						StockCode         string  `json:"stockCode"`
						StockIsin         string  `json:"stockIsin"`
						StockSymbol       any     `json:"stockSymbol"`
						StockName         string  `json:"stockName"`
						LogoImageURL      string  `json:"logoImageUrl"`
						Quantity          float64 `json:"quantity"`
						TradableQuantity  float64 `json:"tradableQuantity"`
						UnsettledQuantity int     `json:"unsettledQuantity"`
						CurrentPrice      struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"currentPrice"`
						BasePrice struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"basePrice"`
						CloseWithoutAfter struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"closeWithoutAfter"`
						BaseWithoutAfter struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"baseWithoutAfter"`
						PurchasePrice struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"purchasePrice"`
						PurchaseAmount struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"purchaseAmount"`
						EvaluatedAmount struct {
							Krw float64 `json:"krw"`
							Usd any     `json:"usd"`
						} `json:"evaluatedAmount"`
						ProfitLossAmount struct {
							Krw float64 `json:"krw"`
							Usd any     `json:"usd"`
						} `json:"profitLossAmount"`
						DailyProfitLossAmount struct {
							Krw float64 `json:"krw"`
							Usd any     `json:"usd"`
						} `json:"dailyProfitLossAmount"`
						ProfitLossRate struct {
							Krw float64 `json:"krw"`
							Usd any     `json:"usd"`
						} `json:"profitLossRate"`
						DailyProfitLossRate struct {
							Krw float64 `json:"krw"`
							Usd any     `json:"usd"`
						} `json:"dailyProfitLossRate"`
						Commission struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"commission"`
						CommissionRate float64 `json:"commissionRate"`
						BuyCommission  struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"buyCommission"`
						SellCommission struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"sellCommission"`
						Tax struct {
							Krw int `json:"krw"`
							Usd any `json:"usd"`
						} `json:"tax"`
						TaxRate               float64 `json:"taxRate"`
						Delisting             bool    `json:"delisting"`
						Unlisting             bool    `json:"unlisting"`
						NxtSupported          bool    `json:"nxtSupported"`
						DomesticExchange      string  `json:"domesticExchange"`
						MarketCode            string  `json:"marketCode"`
						StockGroupCode        string  `json:"stockGroupCode"`
						StockWarrants         bool    `json:"stockWarrants"`
						StockWarrantsLink     any     `json:"stockWarrantsLink"`
						ShortSellingQuantity  int     `json:"shortSellingQuantity"`
						RightExpectedQuantity int     `json:"rightExpectedQuantity"`
						RightEvaluatedAmount  int     `json:"rightEvaluatedAmount"`
						Notice                struct {
							SplitMerge           bool `json:"splitMerge"`
							EarningsAnnouncement bool `json:"earningsAnnouncement"`
						} `json:"notice"`
						ShareHoldingsType string `json:"shareHoldingsType"`
					} `json:"items"`
					PrincipalAmount struct {
						Krw int `json:"krw"`
						Usd any `json:"usd"`
					} `json:"principalAmount"`
					EvaluatedAmount struct {
						Krw float64 `json:"krw"`
						Usd any     `json:"usd"`
					} `json:"evaluatedAmount"`
					ProfitLossAmount struct {
						Krw float64 `json:"krw"`
						Usd any     `json:"usd"`
					} `json:"profitLossAmount"`
					DailyProfitLossAmount struct {
						Krw float64 `json:"krw"`
						Usd any     `json:"usd"`
					} `json:"dailyProfitLossAmount"`
					ProfitLossRate struct {
						Krw float64 `json:"krw"`
						Usd int     `json:"usd"`
					} `json:"profitLossRate"`
					DailyProfitLossRate struct {
						Krw float64 `json:"krw"`
						Usd int     `json:"usd"`
					} `json:"dailyProfitLossRate"`
					Sorted       bool `json:"sorted"`
					HasDelisting bool `json:"hasDelisting"`
				} `json:"kr"`
				Us struct {
					Items []struct {
						ID                int     `json:"id"`
						Key               string  `json:"key"`
						StockCode         string  `json:"stockCode"`
						StockIsin         string  `json:"stockIsin"`
						StockSymbol       string  `json:"stockSymbol"`
						StockName         string  `json:"stockName"`
						LogoImageURL      string  `json:"logoImageUrl"`
						Quantity          float64 `json:"quantity"`
						TradableQuantity  float64 `json:"tradableQuantity"`
						UnsettledQuantity int     `json:"unsettledQuantity"`
						CurrentPrice      struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"currentPrice"`
						BasePrice struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"basePrice"`
						CloseWithoutAfter struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"closeWithoutAfter"`
						BaseWithoutAfter struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"baseWithoutAfter"`
						PurchasePrice struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"purchasePrice"`
						PurchaseAmount struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"purchaseAmount"`
						EvaluatedAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"evaluatedAmount"`
						ProfitLossAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"profitLossAmount"`
						DailyProfitLossAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"dailyProfitLossAmount"`
						ProfitLossRate struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"profitLossRate"`
						DailyProfitLossRate struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"dailyProfitLossRate"`
						Commission struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"commission"`
						CommissionRate float64 `json:"commissionRate"`
						BuyCommission  struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"buyCommission"`
						SellCommission struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"sellCommission"`
						Tax                   any    `json:"tax"`
						TaxRate               any    `json:"taxRate"`
						Delisting             bool   `json:"delisting"`
						Unlisting             bool   `json:"unlisting"`
						NxtSupported          bool   `json:"nxtSupported"`
						DomesticExchange      string `json:"domesticExchange"`
						MarketCode            string `json:"marketCode"`
						StockGroupCode        string `json:"stockGroupCode"`
						StockWarrants         bool   `json:"stockWarrants"`
						StockWarrantsLink     any    `json:"stockWarrantsLink"`
						ShortSellingQuantity  int    `json:"shortSellingQuantity"`
						RightExpectedQuantity int    `json:"rightExpectedQuantity"`
						RightEvaluatedAmount  int    `json:"rightEvaluatedAmount"`
						Notice                struct {
							SplitMerge           bool `json:"splitMerge"`
							EarningsAnnouncement bool `json:"earningsAnnouncement"`
						} `json:"notice"`
						ShareHoldingsType string `json:"shareHoldingsType"`
					} `json:"items"`
					PrincipalAmount struct {
						Krw int     `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"principalAmount"`
					EvaluatedAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"evaluatedAmount"`
					ProfitLossAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"profitLossAmount"`
					DailyProfitLossAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"dailyProfitLossAmount"`
					ProfitLossRate struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"profitLossRate"`
					DailyProfitLossRate struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"dailyProfitLossRate"`
					Sorted       bool `json:"sorted"`
					HasDelisting bool `json:"hasDelisting"`
				} `json:"us"`
				Option struct {
					Items           []any `json:"items"`
					PrincipalAmount struct {
						Krw int `json:"krw"`
						Usd int `json:"usd"`
					} `json:"principalAmount"`
					EvaluatedAmount struct {
						Krw int `json:"krw"`
						Usd int `json:"usd"`
					} `json:"evaluatedAmount"`
					ProfitLossAmount struct {
						Krw int `json:"krw"`
						Usd int `json:"usd"`
					} `json:"profitLossAmount"`
					DailyProfitLossAmount struct {
						Krw int `json:"krw"`
						Usd int `json:"usd"`
					} `json:"dailyProfitLossAmount"`
					ProfitLossRate struct {
						Krw int `json:"krw"`
						Usd int `json:"usd"`
					} `json:"profitLossRate"`
					DailyProfitLossRate struct {
						Krw int `json:"krw"`
						Usd int `json:"usd"`
					} `json:"dailyProfitLossRate"`
					Sorted       bool `json:"sorted"`
					HasDelisting bool `json:"hasDelisting"`
				} `json:"option"`
				Bond struct {
					Items []struct {
						ID                 int       `json:"id"`
						Key                string    `json:"key"`
						GUID               string    `json:"guid"`
						Isin               string    `json:"isin"`
						Symbol             any       `json:"symbol"`
						DisplayName        string    `json:"displayName"`
						UnderlyingGUID     string    `json:"underlyingGuid"`
						LogoImageURL       string    `json:"logoImageUrl"`
						MaturityDate       string    `json:"maturityDate"`
						MaturityDateTime   time.Time `json:"maturityDateTime"`
						DurationToMaturity string    `json:"durationToMaturity"`
						UnsettledQuantity  float64   `json:"unsettledQuantity"`
						TradableQuantity   float64   `json:"tradableQuantity"`
						CurrentPrice       struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"currentPrice"`
						BasePrice struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"basePrice"`
						PurchasePrice struct {
							Krw int     `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"purchasePrice"`
						ExpectedMaturityProfitLoss struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"expectedMaturityProfitLoss"`
						ExpectedMaturityProfitLossRate struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"expectedMaturityProfitLossRate"`
						ExpectedMaturityCommission struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"expectedMaturityCommission"`
						Commission struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"commission"`
						ExpectedMaturityTax struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"expectedMaturityTax"`
						Tax struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"tax"`
						ExpectedMaturityAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"expectedMaturityAmount"`
						PurchaseAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"purchaseAmount"`
						EvaluatedAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"evaluatedAmount"`
						ProfitLossAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"profitLossAmount"`
						DailyProfitLossAmount struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"dailyProfitLossAmount"`
						ProfitLossRate struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"profitLossRate"`
						DailyProfitLossRate struct {
							Krw float64 `json:"krw"`
							Usd float64 `json:"usd"`
						} `json:"dailyProfitLossRate"`
						Delisting         bool   `json:"delisting"`
						ShareHoldingsType string `json:"shareHoldingsType"`
					} `json:"items"`
					PrincipalAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"principalAmount"`
					EvaluatedAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"evaluatedAmount"`
					ExpectedMaturityProfitLoss struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"expectedMaturityProfitLoss"`
					ExpectedMaturityProfitLossRate struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"expectedMaturityProfitLossRate"`
					ExpectedMaturityCommission struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"expectedMaturityCommission"`
					Commission struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"commission"`
					ExpectedMaturityTax struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"expectedMaturityTax"`
					Tax struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"tax"`
					ExpectedMaturityAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"expectedMaturityAmount"`
					ProfitLossAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"profitLossAmount"`
					DailyProfitLossAmount struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"dailyProfitLossAmount"`
					ProfitLossRate struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"profitLossRate"`
					DailyProfitLossRate struct {
						Krw float64 `json:"krw"`
						Usd float64 `json:"usd"`
					} `json:"dailyProfitLossRate"`
					Sorted       bool `json:"sorted"`
					HasDelisting bool `json:"hasDelisting"`
				} `json:"bond"`
				HiddenStock struct {
					Count  int  `json:"count"`
					All    bool `json:"all"`
					Amount int  `json:"amount"`
				} `json:"hiddenStock"`
				UsePolling bool `json:"usePolling"`
				StockNudge any  `json:"stockNudge"`
				BondNudge  any  `json:"bondNudge"`
			} `json:"data"`
			Type string `json:"type"`
		} `json:"sections"`
}

type dashboardAssetSectionsAllResponse struct {
	Result DashboardAssetSectionsAllResponse `json:"result"`
}

func (c *V1DashboardAsset) SectionsAll(ctx context.Context, req *DashboardAssetSectionsAllRequest) (*DashboardAssetSectionsAllResponse, error) {
	respBody := &dashboardAssetSectionsAllResponse{}
	httpReq, err := c.newRequest(ctx, http.MethodPost, "/api/v1/dashboard/asset/sections/all", req)
	if err != nil {
		return nil, err
	}
	err = c.postJson(httpReq, &respBody)
	return &respBody.Result, err
}
