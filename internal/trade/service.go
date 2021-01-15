package trade

import (
	"log"
	"os"
	"poketrader/pkg/pokeapi"
	"strconv"
	"sync"
)

type TradeManager interface {
	TradeList() ([]Trade, error)
	TradeCalculator(trade Trade) (bool, error)
}

type tradeManager struct {
	Repo Repository
}

func NewTradeManager() TradeManager {
	return tradeManager{Repo: NewRepo()}
}

func (tm tradeManager) TradeList() ([]Trade, error) {
	return tm.Repo.Get()
}

func (tm tradeManager) TradeCalculator(trade Trade) (bool, error) {
	var wg sync.WaitGroup
	indexOne := int(0)
	indexTwo := int(0)
	for _, pk := range trade.FirstTrainerList {
		wg.Add(1)

		go func(p *Pokemon) {

			pokemon := tm.getPokemon(p.Name)
			p.ID = pokemon.ID
			p.BaseExperience = pokemon.BaseExperience
			indexOne += pokemon.BaseExperience
			wg.Done()
		}(pk)
	}
	for _, pk := range trade.SecondTrainerList {
		wg.Add(1)

		go func(p *Pokemon) {
			pokemon := tm.getPokemon(p.Name)
			p.ID = pokemon.ID
			p.BaseExperience = pokemon.BaseExperience
			indexTwo += pokemon.BaseExperience
			wg.Done()
		}(pk)
	}
	wg.Wait()
	trade.Fair = tm.isFairTrade(indexOne, indexTwo)
	if err := tm.Repo.Save(trade); err != nil {
		log.Printf("Error on save trade: %v\n", err)
		return false, err
	}
	return trade.Fair, nil
}

func (tm tradeManager) getPokemon(pokemon string) Pokemon {
	pkm, err := pokeapi.GetPokemon(pokemon)
	if err != nil {
		log.Printf("Error on Get pokemon %s : %s", pokemon, err.Error())
	}
	return Pokemon{
		ID:             pkm.ID,
		Name:           pkm.Name,
		BaseExperience: pkm.BaseExperience,
	}
}

func (tm tradeManager) isFairTrade(indexOne, indexTwo int) bool {
	fairIndex, _ := strconv.Atoi(os.Getenv("FAIR_INDEX"))
	difference := indexOne - indexTwo
	if difference < 0 {
		difference *= -1
	}
	return fairIndex >= difference
}
