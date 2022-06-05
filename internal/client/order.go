package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gitlab.com/modanisatech/marketplace/shared/httpkit"
	"go.uber.org/zap"
)

const ordersPath = "/orders"

var ErrOrderNotExists = errors.New("order not exists")

type OrderClient struct {
	Logger     *zap.Logger
	BaseURL    string
	httpClient http.Client
}

func New(baseURL string, logget *zap.Logger) *OrderClient {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	return &OrderClient{
		BaseURL:    baseURL,
		httpClient: httpkit.NewClientWithRoundTrippers(httpClient),
	}
}

func (o *OrderClient) GetOrder(ctx context.Context, id int) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%s/%d", o.BaseURL, ordersPath, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	res, err := o.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		orders := make(map[string]interface{})
		err = json.Unmarshal(body, &orders)
		if err != nil {
			return nil, err
		}

		return orders, nil
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, ErrOrderNotExists
	}

	return nil, errors.New("unexpected Error")
}
