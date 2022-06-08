package coin_service

import (
	"fmt"
	"github.com/hberkayozdemir/hibemi-be/internal/coin"
	gecko "github.com/superoo7/go-gecko/v3"
	"log"
	http "net/http"
)

type CGclient struct {
	Repository coin.Repository
}

func NewCoinGeckoClient(repository coin.Repository) *CGclient {
	return &CGclient{
		repository,
	}
}

func (cgg *CGclient) FetchCoins() {
	http := &http.Client{}
	cg := gecko.NewClient(http)
	coin, err := cg.CoinsID("dogecoin", true, true, true, true, true, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coin)
}
