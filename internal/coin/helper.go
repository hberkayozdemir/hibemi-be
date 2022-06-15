package coin

import (
	"github.com/hberkayozdemir/hibemi-be/internal/coin_gecko"
)

func CoinGeckoEntityToModel(entity coin_gecko.CoinGeckoResponseEntity) coin_gecko.CoinGeckoResponse {
	return coin_gecko.CoinGeckoResponse{
		ID:              entity.ID,
		Symbol:          entity.Symbol,
		Name:            entity.Name,
		Price:           entity.Price,
		AssetPlatformID: entity.AssetPlatformID,
		Platforms: struct {
			NAMING_FAILED string `json:""`
		}(entity.Platforms),
		BlockTimeInMinutes: entity.BlockTimeInMinutes,
		HashingAlgorithm:   entity.HashingAlgorithm,
		Categories:         entity.Categories,
		PublicNotice:       entity.PublicNotice,
		AdditionalNotices:  entity.AdditionalNotices,
		Description: struct {
			En string `json:"en"`
		}(entity.Description),
		Links: struct {
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
		}(entity.Links),
		Image: struct {
			Thumb string `json:"thumb"`
			Small string `json:"small"`
			Large string `json:"large"`
		}(entity.Image),
		CountryOrigin:                entity.CountryOrigin,
		GenesisDate:                  entity.GenesisDate,
		SentimentVotesUpPercentage:   entity.SentimentVotesUpPercentage,
		SentimentVotesDownPercentage: entity.SentimentVotesDownPercentage,
		MarketCapRank:                entity.MarketCapRank,
		CoingeckoRank:                entity.CoingeckoRank,
		CoingeckoScore:               entity.CoingeckoScore,
		DeveloperScore:               entity.DeveloperScore,
		CommunityScore:               entity.CommunityScore,
		LiquidityScore:               entity.LiquidityScore,
		PublicInterestScore:          entity.PublicInterestScore,
		PublicInterestStats: struct {
			AlexaRank   int         `json:"alexa_rank"`
			BingMatches interface{} `json:"bing_matches"`
		}(entity.PublicInterestStats),
		StatusUpdates: entity.StatusUpdates,
		LastUpdated:   entity.LastUpdated,
	}
}
