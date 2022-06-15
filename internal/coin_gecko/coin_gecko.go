package coin_gecko

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"path"
	"strings"
	"time"
)

type APIClient struct {
	BaseURL    string
	httpClient *http.Client
}

type CoinGeckoResponse struct {
	ID              string      `json:"id"`
	Symbol          string      `json:"symbol"`
	Price           string      `json:"price"`
	Name            string      `json:"name"`
	AssetPlatformID interface{} `json:"asset_platform_id"`
	Platforms       struct {
		NAMING_FAILED string `json:""`
	} `json:"platforms"`
	BlockTimeInMinutes int           `json:"block_time_in_minutes"`
	HashingAlgorithm   string        `json:"hashing_algorithm"`
	Categories         []string      `json:"categories"`
	PublicNotice       interface{}   `json:"public_notice"`
	AdditionalNotices  []interface{} `json:"additional_notices"`
	Description        struct {
		En string `json:"en"`
	} `json:"description"`
	Links struct {
		Homepage                    []string    `json:"homepage"`
		BlockchainSite              []string    `json:"blockchain_site"`
		OfficialForumURL            []string    `json:"official_forum_url"`
		ChatURL                     []string    `json:"chat_url"`
		AnnouncementURL             []string    `json:"announcement_url"`
		TwitterScreenName           string      `json:"twitter_screen_name"`
		FacebookUsername            string      `json:"facebook_username"`
		BitcointalkThreadIdentifier interface{} `json:"bitcointalk_thread_identifier"`
		TelegramChannelIdentifier   string      `json:"telegram_channel_identifier"`
		SubredditURL                string      `json:"subreddit_url"`
		ReposURL                    struct {
			Github    []string      `json:"github"`
			Bitbucket []interface{} `json:"bitbucket"`
		} `json:"repos_url"`
	} `json:"links"`
	Image struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	CountryOrigin                string  `json:"country_origin"`
	GenesisDate                  string  `json:"genesis_date"`
	SentimentVotesUpPercentage   float64 `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64 `json:"sentiment_votes_down_percentage"`
	MarketCapRank                int     `json:"market_cap_rank"`
	CoingeckoRank                int     `json:"coingecko_rank"`
	CoingeckoScore               float64 `json:"coingecko_score"`
	DeveloperScore               float64 `json:"developer_score"`
	CommunityScore               float64 `json:"community_score"`
	LiquidityScore               float64 `json:"liquidity_score"`
	PublicInterestScore          float64 `json:"public_interest_score"`
	PublicInterestStats          struct {
		AlexaRank   int         `json:"alexa_rank"`
		BingMatches interface{} `json:"bing_matches"`
	} `json:"public_interest_stats"`
	StatusUpdates []interface{} `json:"status_updates"`
	LastUpdated   time.Time     `json:"last_updated"`
}

type CoinGeckoResponseEntity struct {
	ID              string      `bson:"id"`
	Symbol          string      `bson:"symbol"`
	Name            string      `bson:"name"`
	Price           string      `bson:"price"`
	AssetPlatformID interface{} `bson:"asset_platform_id"`
	Platforms       struct {
		NAMING_FAILED string `bson:""`
	} `bson:"platforms"`
	BlockTimeInMinutes int           `bson:"block_time_in_minutes"`
	HashingAlgorithm   string        `bson:"hashing_algorithm"`
	Categories         []string      `bson:"categories"`
	PublicNotice       interface{}   `bson:"public_notice"`
	AdditionalNotices  []interface{} `bson:"additional_notices"`
	Description        struct {
		En string `bson:"en"`
	} `bson:"description"`
	Links struct {
		Homepage                    []string    `bson:"homepage"`
		BlockchainSite              []string    `bson:"blockchain_site"`
		OfficialForumURL            []string    `bson:"official_forum_url"`
		ChatURL                     []string    `bson:"chat_url"`
		AnnouncementURL             []string    `bson:"announcement_url"`
		TwitterScreenName           string      `bson:"twitter_screen_name"`
		FacebookUsername            string      `bson:"facebook_username"`
		BitcointalkThreadIdentifier interface{} `bson:"bitcointalk_thread_identifier"`
		TelegramChannelIdentifier   string      `bson:"telegram_channel_identifier"`
		SubredditURL                string      `bson:"subreddit_url"`
		ReposURL                    struct {
			Github    []string      `bson:"github"`
			Bitbucket []interface{} `bson:"bitbucket"`
		} `bson:"repos_url"`
	} `bson:"links"`
	Image struct {
		Thumb string `bson:"thumb"`
		Small string `bson:"small"`
		Large string `bson:"large"`
	} `bson:"image"`
	CountryOrigin                string  `bson:"country_origin"`
	GenesisDate                  string  `bson:"genesis_date"`
	SentimentVotesUpPercentage   float64 `bson:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64 `bson:"sentiment_votes_down_percentage"`
	MarketCapRank                int     `bson:"market_cap_rank"`
	CoingeckoRank                int     `bson:"coingecko_rank"`
	CoingeckoScore               float64 `bson:"coingecko_score"`
	DeveloperScore               float64 `bson:"developer_score"`
	CommunityScore               float64 `bson:"community_score"`
	LiquidityScore               float64 `bson:"liquidity_score"`
	PublicInterestScore          float64 `bson:"public_interest_score"`
	PublicInterestStats          struct {
		AlexaRank   int         `bson:"alexa_rank"`
		BingMatches interface{} `bson:"bing_matches"`
	} `bson:"public_interest_stats"`
	StatusUpdates []interface{} `bson:"status_updates"`
	LastUpdated   time.Time     `bson:"last_updated"`
}

func NewClient(baseUrl string) APIClient {

	client := http.Client{}
	return APIClient{
		BaseURL:    strings.TrimRight(baseUrl, "/") + "/",
		httpClient: &client,
	}
}

func (c APIClient) GetCoin(coinID string) (*CoinGeckoResponse, error) {
	endPoint := c.BaseURL + path.Join("api", "v3", "coins", coinID)
	url := fmt.Sprintf("%s?localization=%s&tickers=%s&market_data=%s&community_data=%s&developer_data=%s&sparkline=%s", endPoint, "false", "false", "false", "false", "false", "false")
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, http.NoBody)
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

	var coin CoinGeckoResponse
	if err := json.NewDecoder(res.Body).Decode(&coin); err != nil {
		return nil, errors.Wrap(err, "error decoding get default address response")
	}

	return &coin, nil
}
