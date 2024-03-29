package trade

import (
	"errors"
	"log"
	"os"
	"poketrader/pkg/pokeapi"
	"strconv"
	"sync"
)

type TradeManager interface {
	TradeList() ([]Trade, error)
	TradeCalculator(trade Trade) (bool, error)
	//////////////////////Trade
	TradeByID(id string) ([]Trade, error)
	DeleteTrade(id string) (string, error)
	UpdateTrade(trade Trade, id string) (bool, error)
	FilterByFair(boolFilter string) ([]Trade, error)
}

type tradeManager struct {
	Repo Repository
}

func NewTradeManager() TradeManager {
	return tradeManager{Repo: NewRepo()}
}

func (tm tradeManager) FilterByFair(boolFilter string) ([]Trade, error) {
	return tm.Repo.FilterByFair(boolFilter)
}

func (tm tradeManager) UpdateTrade(trade Trade, id string) (bool, error) {
	var wg sync.WaitGroup
	indexOne := int(0)
	indexTwo := int(0)
	if err := tm.validateQuantity(trade.FirstTrainerList, trade.SecondTrainerList); err != nil {
		return false, err
	}
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
	if err := tm.Repo.UpdateTrade(trade, id); err != nil {
		log.Printf("Error on save trade: %v\n", err)
		return false, err
	}

	return trade.Fair, nil

}

func (tm tradeManager) DeleteTrade(id string) (string, error) {
	return tm.Repo.DeleteByID(id)
}

/////////////////////////////////////////////Trade
func (tm tradeManager) TradeByID(id string) ([]Trade, error) {
	return tm.Repo.GetByID(id)

}

func (tm tradeManager) TradeList() ([]Trade, error) {
	return tm.Repo.Get()
}

func (tm tradeManager) TradeCalculator(trade Trade) (bool, error) {
	var wg sync.WaitGroup
	indexOne := int(0)
	indexTwo := int(0)
	if err := tm.validateQuantity(trade.FirstTrainerList, trade.SecondTrainerList); err != nil {
		return false, err
	}
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

func (tm tradeManager) validateQuantity(line1 []*Pokemon, line2 []*Pokemon) error {
	if len(line1) > 6 || len(line1) < 1 {
		return errors.New("only allowed between 1 and 6 pokemons")
	}
	if len(line2) > 6 || len(line2) < 1 {
		return errors.New("only allowed between 1 and 6 pokemons")
	}
	return nil
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
