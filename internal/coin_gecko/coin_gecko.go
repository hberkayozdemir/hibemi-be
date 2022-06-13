package coin_gecko

import (
	"net/http"
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
