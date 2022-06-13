package coin_gecko

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"path"
	"strings"
)

type APIClient struct {
	BaseURL    string
	httpClient *http.Client
}

func NewClient(baseUrl string) APIClient {

	client := http.Client{}
	return APIClient{
		BaseURL:    strings.TrimRight(baseUrl, "/") + "/",
		httpClient: &client,
	}
}

func (c APIClient) GetCoin(coinID string) (*map[string]interface{}, error) {
	endPoint := c.BaseURL + path.Join("api", "v3", "coins", coinID)
	url := fmt.Sprintf("%s?localization=%s&tickers=%s&market_data=%s&community_data=%s&developer_data=%s&sparkline=%s", endPoint, "false", "false", "false", "false", "false", "false")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create request to get coin")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "unable to fetch coin from coin gecko")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.WithStack(fmt.Errorf("get coin failed"))
	}

	var coin map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&coin); err != nil {
		return nil, errors.Wrap(err, "error decoding get default address response")
	}

	return &coin, nil
}
