package coinmarketcap

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

func FetchETHPrice(apiKey string) (float64, error) {
	url := "https://pro-api.coinmarketcap.com/v2/cryptocurrency/quotes/latest?symbol=ETH"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return 0, err
	}
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.Header.Add("Accept", "*/*")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	cmc := &CMCResp{}
	err = json.Unmarshal(body, cmc)
	if err != nil {
		return 0, err
	}
	if 0 != cmc.Status.ErrorCode {
		return 0, errors.New(cmc.Status.ErrorMessage.(string))
	}
	if len(cmc.Data.ETH) == 0 {
		return 0, errors.New("no price data")
	}
	return cmc.Data.ETH[0].Quote.USD.Price, nil
}

type CMCResp struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		ETH []struct {
			Id             int       `json:"id"`
			Name           string    `json:"name"`
			Symbol         string    `json:"symbol"`
			Slug           string    `json:"slug"`
			NumMarketPairs int       `json:"num_market_pairs"`
			DateAdded      time.Time `json:"date_added"`
			Tags           []struct {
				Slug     string `json:"slug"`
				Name     string `json:"name"`
				Category string `json:"category"`
			} `json:"tags"`
			MaxSupply                     interface{} `json:"max_supply"`
			CirculatingSupply             interface{} `json:"circulating_supply"`
			TotalSupply                   interface{} `json:"total_supply"`
			IsActive                      int         `json:"is_active"`
			Platform                      interface{} `json:"platform"`
			CmcRank                       int         `json:"cmc_rank"`
			IsFiat                        int         `json:"is_fiat"`
			SelfReportedCirculatingSupply interface{} `json:"self_reported_circulating_supply"`
			SelfReportedMarketCap         interface{} `json:"self_reported_market_cap"`
			TvlRatio                      interface{} `json:"tvl_ratio"`
			LastUpdated                   time.Time   `json:"last_updated"`
			Quote                         struct {
				USD struct {
					Price                 float64     `json:"price"`
					Volume24H             float64     `json:"volume_24h"`
					VolumeChange24H       float64     `json:"volume_change_24h"`
					PercentChange1H       float64     `json:"percent_change_1h"`
					PercentChange24H      float64     `json:"percent_change_24h"`
					PercentChange7D       float64     `json:"percent_change_7d"`
					PercentChange30D      float64     `json:"percent_change_30d"`
					PercentChange60D      float64     `json:"percent_change_60d"`
					PercentChange90D      float64     `json:"percent_change_90d"`
					MarketCap             float64     `json:"market_cap"`
					MarketCapDominance    float64     `json:"market_cap_dominance"`
					FullyDilutedMarketCap float64     `json:"fully_diluted_market_cap"`
					Tvl                   interface{} `json:"tvl"`
					LastUpdated           time.Time   `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} `json:"ETH"`
	} `json:"data"`
}
