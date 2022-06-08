package coin_service

type Coin struct {
	BlockTimeInMin      int32                   `json:"block_time_in_minutes" bson:"block_time_in_minutes"`
	Categories          []string                `json:"categories" bson:"categories"`
	Localization        map[string]string       `json:"localization" bson:"localization"`
	Description         map[string]string       `json:"description" bson:"description"`
	Links               *map[string]interface{} `json:"links" bson:"links"`
	Image               ImageItem               `json:"image" bson:"image"`
	CountryOrigin       string                  `json:"country_origin" bson:"country_origin"`
	GenesisDate         string                  `json:"genesis_date" bson:"genesis_date"`
	MarketCapRank       uint16                  `json:"market_cap_rank" bson:"market_cap_rank"`
	CoinGeckoRank       uint16                  `json:"coingecko_rank" bson:"coingecko_rank"`
	CoinGeckoScore      float32                 `json:"coingecko_score" bson:"coingecko_score"`
	DeveloperScore      float32                 `json:"developer_score" bson:"developer_score"`
	CommunityScore      float32                 `json:"community_score" bson:"community_score"`
	LiquidityScore      float32                 `json:"liquidity_score" bson:"liquidity_score"`
	PublicInterestScore float32                 `json:"public_interest_score" bson:"public_interest_score"`
	MarketData          *interface{}            `json:"market_data" bson:"market_data"`
	CommunityData       *interface{}            `json:"community_data" bson:"community_data"`
	DeveloperData       *interface{}            `json:"developer_data" bson:"developer_data"`
	PublicInterestStats *interface{}            `json:"public_interest_stats" bson:"public_interest_stats"`
	StatusUpdates       *[]interface{}          `json:"status_updates" bson:"status_updates"`
	LastUpdated         string                  `json:"last_updated" bson:"last_updated"`
	Tickers             *[]interface{}          `json:"tickers" bson:"tickers"`
}

type ImageItem struct {
	Thumb string `json:"thumb" bson:"thumb"`
	Small string `json:"small" bson:"small"`
	Large string `json:"large" bson:"large"`
}
