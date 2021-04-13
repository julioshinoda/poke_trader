package trade

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func Handler(r chi.Router) {

	r.Post("/trade", TradeHandler)
	r.Get("/trade", TradeListHandler)
	r.Get("/trade/{id}", GetTradeByIDHandler)

}
func GetTradeByIDHandler(w http.ResponseWriter, r *http.Request) {
	tradeID := chi.URLParam(r, "id")
	service := NewTradeManager()
	trade, err := service.TradeByID(tradeID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		response, _ := json.Marshal(map[string]interface{}{"error": err.Error()})
		w.Write(response)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	response, _ := json.Marshal(trade)
	w.Write(response)
}
func TradeListHandler(w http.ResponseWriter, r *http.Request) {
	service := NewTradeManager()
	list, err := service.TradeList()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(list)
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func TradeHandler(w http.ResponseWriter, r *http.Request) {
	var request Trade
	json.NewDecoder(r.Body).Decode(&request)
	service := NewTradeManager()
	good, err := service.TradeCalculator(request)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		response, _ := json.Marshal(map[string]string{"error": err.Error()})
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	response, _ := json.Marshal(map[string]bool{"fair_trade": good})
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
