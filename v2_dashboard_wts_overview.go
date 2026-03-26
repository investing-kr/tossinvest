package tossinvest

import (
	"context"
	"net/http"
)

type V2DashboardWtsOverview service

type DashboardWtsOverviewRankingRequest struct {
	ID      string   `json:"id"`
	Filters []string `json:"filters"`
	Tag     string   `json:"tag"`
}

type RankingProduct struct {
	ProductCode  string `json:"productCode"`
	Name         string `json:"name"`
	LogoImageURL string `json:"logoImageUrl"`
	Rank         int    `json:"rank"`
	Price        struct {
		Base                   float64 `json:"base"`
		BaseKrw                float64 `json:"baseKrw"`
		Close                  float64 `json:"close"`
		CloseKrw               float64 `json:"closeKrw"`
		PriceType              string  `json:"priceType"`
		TossSecuritiesAmount   float64 `json:"tossSecuritiesAmount"`
		TossSecuritiesVolume   int     `json:"tossSecuritiesVolume"`
	} `json:"price"`
	ExtraInfo struct {
		TossSecuritiesBuy  int `json:"tossSecuritiesBuy"`
		TossSecuritiesSell int `json:"tossSecuritiesSell"`
	} `json:"extraInfo"`
}

type DashboardWtsOverviewRankingResponse struct {
	BasedAt               string           `json:"basedAt"`
	Duration              string           `json:"duration"`
	Filters               []string         `json:"filters"`
	RankingID             string           `json:"rankingId"`
	RankingIDWithDuration string           `json:"rankingIdWithDuration"`
	Tag                   string           `json:"tag"`
	Type                  string           `json:"type"`
	Products              []RankingProduct `json:"products"`
}

type dashboardWtsOverviewRankingResponse struct {
	Result DashboardWtsOverviewRankingResponse `json:"result"`
}

func (c *V2DashboardWtsOverview) Ranking(ctx context.Context, req *DashboardWtsOverviewRankingRequest) (*DashboardWtsOverviewRankingResponse, error) {
	respBody := &dashboardWtsOverviewRankingResponse{}

	if req.ID != "" {
		req.ID = "biggest_total_amount"
	}

	if req.Tag != "" {
		req.Tag = "us"
	}

	httpReq, err := c.newRequest(ctx, http.MethodPost, "/api/v2/dashboard/wts/overview/ranking", req)
	if err != nil {
		return nil, err
	}


	err = c.postJson(httpReq, &respBody)
	return &respBody.Result, err
}
