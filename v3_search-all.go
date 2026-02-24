package tossinvest

import (
	"context"
	"net/http"
)

type V3SearchAll service

type SearchAutoCompleteRequest struct {
	Query    string          `json:"query"`
	Sections []SearchSection `json:"sections"`
}

type SearchSection struct {
	Type string `json:"type"`
}

type SearchAutoCompleteResponseDataItem struct {
	Keyword      string `json:"keyword"`
	SubKeyword   string `json:"subKeyword"`
	ProductCode  string `json:"productCode"`
	ProductName  string `json:"productName"`
	Symbol       string `json:"symbol"`
	CompanyCode  string `json:"companyCode"`
	LogoImageUrl string `json:"logoImageUrl"`
	Market       string `json:"market"`
	Base         struct {
		Krw int     `json:"krw"`
		Usd float64 `json:"usd"`
	} `json:"base"`
	Close struct {
		Krw int     `json:"krw"`
		Usd float64 `json:"usd"`
	} `json:"close"`
	StockStatus     string `json:"stockStatus"`
	AutoComplete    bool   `json:"autoComplete"`
	Code            string `json:"code"`
	SubSectionQuery string `json:"subSectionQuery"`
}

type SearchAutoCompleteResponseData struct {
	Items       []SearchAutoCompleteResponseDataItem `json:"items"`
	SubSections []string                             `json:"subSections"`
}

type SearchAutoCompleteResponse struct {
	Type string                         `json:"type"`
	Data SearchAutoCompleteResponseData `json:"data"`
}

type searchAutoCompleteResponse struct {
	Result []SearchAutoCompleteResponse `json:"result"`
}

func (c *V3SearchAll) WtsAutoComplete(ctx context.Context, req *SearchAutoCompleteRequest) (*searchAutoCompleteResponse, error) {
	respBody := &searchAutoCompleteResponse{}
	httpReq, err := c.newRequest(ctx, http.MethodPost, "/api/v3/search-all/wts-auto-complete", req)
	if err != nil {
		return nil, err
	}

	// Override base URL to use info API
	httpReq.URL.Host = "wts-info-api.tossinvest.com"
	httpReq.URL.Scheme = "https"

	err = c.postJson(httpReq, &respBody)
	return respBody, err
}
