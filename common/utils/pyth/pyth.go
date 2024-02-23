package pyth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Mu-Exchange/Mu-End/common/utils/hexutil"
)

type EvmPriceServiceConnection struct {
	url string
	ctx context.Context
}

func NewEvmPriceServiceConnection(url string) *EvmPriceServiceConnection {
	return &EvmPriceServiceConnection{
		url: url,
	}
}

func (c *EvmPriceServiceConnection) GetPriceFeedsUpdateData(priceIds []string) ([]string, error) {
	params := url.Values{}
	for _, priceId := range priceIds {
		params.Add("ids", priceId)
	}
	fullURL := fmt.Sprintf("%s/api/latest_vaas?%s", c.url, params.Encode())
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get price feeds update data failed, status code: %d, body: %s", resp.StatusCode, body)
	}

	var latestVaas []string
	if err := json.Unmarshal(body, &latestVaas); err != nil {
		return nil, err
	}

	result := make([]string, 0, len(latestVaas))
	for _, vaa := range latestVaas {
		bytes, err := base64.StdEncoding.DecodeString(vaa)
		if err != nil {
			return nil, err
		}
		result = append(result, hexutil.Encode(bytes))
	}

	return result, nil
}
