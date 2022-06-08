package coin

import "time"

type CoinDTO struct {
	ID                 string        `json:"id"`
	Symbol             string        `json:"symbol"`
	Name               string        `json:"name"`
	AssetPlatformID    interface{}   `json:"asset_platform_id"`
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
	Tickers       []struct {
		Base   string `json:"base"`
		Target string `json:"target"`
		Market struct {
			Name                string `json:"name"`
			Identifier          string `json:"identifier"`
			HasTradingIncentive bool   `json:"has_trading_incentive"`
		} `json:"market"`
		Last          float64 `json:"last"`
		Volume        float64 `json:"volume"`
		ConvertedLast struct {
			Btc int     `json:"btc"`
			Eth float64 `json:"eth"`
			Usd int     `json:"usd"`
		} `json:"converted_last"`
		ConvertedVolume struct {
			Btc int `json:"btc"`
			Eth int `json:"eth"`
			Usd int `json:"usd"`
		} `json:"converted_volume"`
		TrustScore             string      `json:"trust_score"`
		BidAskSpreadPercentage float64     `json:"bid_ask_spread_percentage"`
		Timestamp              time.Time   `json:"timestamp"`
		LastTradedAt           time.Time   `json:"last_traded_at"`
		LastFetchAt            time.Time   `json:"last_fetch_at"`
		IsAnomaly              bool        `json:"is_anomaly"`
		IsStale                bool        `json:"is_stale"`
		TradeURL               string      `json:"trade_url"`
		TokenInfoURL           interface{} `json:"token_info_url"`
		CoinID                 string      `json:"coin_id"`
		TargetCoinID           string      `json:"target_coin_id,omitempty"`
	} `json:"tickers"`
}
