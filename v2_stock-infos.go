package tossinvest

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type V2StockInfos struct {
	*Client
	cache sync.Map
}

type StockInfoMarket struct {
	Code        string `json:"code"`
	DisplayName string `json:"displayName"`
}

type StockInfoGroup struct {
	Code        string `json:"code"`
	DisplayName string `json:"displayName"`
}

type StockInfoResponse struct {
	Code                            string          `json:"code"`
	GUID                            string          `json:"guid"`
	Symbol                          string          `json:"symbol"`
	ISINCode                        string          `json:"isinCode"`
	Status                          string          `json:"status"`
	Name                            string          `json:"name"`
	EnglishName                     string          `json:"englishName"`
	DetailName                      string          `json:"detailName"`
	Market                          StockInfoMarket `json:"market"`
	Group                           StockInfoGroup  `json:"group"`
	CompanyCode                     string          `json:"companyCode"`
	CompanyName                     string          `json:"companyName"`
	LogoImageUrl                    string          `json:"logoImageUrl"`
	Currency                        string          `json:"currency"`
	TradingSuspended                bool            `json:"tradingSuspended"`
	KrxTradingSuspended             bool            `json:"krxTradingSuspended"`
	NxtTradingSuspended             bool            `json:"nxtTradingSuspended"`
	CommonShare                     bool            `json:"commonShare"`
	Spac                            bool            `json:"spac"`
	SpacMergerDate                  *string         `json:"spacMergerDate"`
	LeverageFactor                  float64         `json:"leverageFactor"`
	Clearance                       bool            `json:"clearance"`
	RiskLevel                       string          `json:"riskLevel"`
	PurchasePrerequisite            string          `json:"purchasePrerequisite"`
	SharesOutstanding               int64           `json:"sharesOutstanding"`
	PrevListDate                    *string         `json:"prevListDate"`
	ListDate                        *string         `json:"listDate"`
	DelistDate                      *string         `json:"delistDate"`
	OfferingPrice                   *float64        `json:"offeringPrice"`
	WarrantsCode                    *string         `json:"warrantsCode"`
	WarrantsGroupCode               *string         `json:"warrantsGroupCode"`
	EtfTaxCode                      *string         `json:"etfTaxCode"`
	DaytimePriceSupported           bool            `json:"daytimePriceSupported"`
	OptionSupported                 bool            `json:"optionSupported"`
	OptionPennyPilotPriceSupported  bool            `json:"optionPennyPilotPriceSupported"`
	OptionOvertimeSupported         bool            `json:"optionOvertimeSupported"`
	OptionInstrument                *string         `json:"optionInstrument"`
	DerivativeEtp                   bool            `json:"derivativeEtp"`
	PoolingStock                    bool            `json:"poolingStock"`
	NxtSupported                    bool            `json:"nxtSupported"`
	UserTradingSuspended            bool            `json:"userTradingSuspended"`
	LimitOnCompetitiveTradingVolume bool            `json:"limitOnCompetitiveTradingVolume"`
	NxtOpenDate                     *string         `json:"nxtOpenDate"`
	NxtOpenDateRecent               *string         `json:"nxtOpenDateRecent"`
	DerivativeEtf                   bool            `json:"derivativeEtf"`
}

type stockInfoResponse struct {
	Result StockInfoResponse `json:"result"`
}

func (c *V2StockInfos) CodeOrSymbol(ctx context.Context, codeOrSymbol string) (*StockInfoResponse, error) {
	url := fmt.Sprintf("https://wts-info-api.tossinvest.com/api/v2/stock-infos/code-or-symbol/%s", codeOrSymbol)

	if v, ok := c.cache.Load(url); ok {
		return v.(*StockInfoResponse), nil
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	respBody := &stockInfoResponse{}
	err = c.getJson(httpReq, respBody)
	if err != nil {
		return nil, err
	}

	c.cache.Store(url, &respBody.Result)
	return &respBody.Result, err
}
