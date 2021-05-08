package trade

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"poketrader/pkg/database/postgres"
	"time"
)

type Repository interface {
	Save(tr Trade) error
	Get() ([]Trade, error)
	////////////////////Trade
	GetByID(id string) ([]Trade, error)
}

type repo struct {
	DBClient *sql.DB
}

func NewRepo() Repository {
	conn, _ := postgres.GetConnection()
	return repo{DBClient: conn}
}

//////////////////////////////////Trade
func (r repo) GetByID(id string) ([]Trade, error) {
	lista := []Trade{}
	q := fmt.Sprintf(`SELECT id,trainerOne,trainerTwo,created_at,fair FROM public.trade where id=%s`, id)
	rows, err := r.DBClient.Query(q)
	if err != nil {
		log.Printf("Error on get trade: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		tr := Trade{}
		var line1, line2 []byte
		if err := rows.Scan(&tr.ID, &line1, &line2, &tr.CreatedAt, &tr.Fair); err != nil {
			log.Printf("Error on get trade: %v\n", err)
			return nil, err

		}
		json.Unmarshal(line1, &tr.FirstTrainerList)

		json.Unmarshal(line2, &tr.SecondTrainerList)

		lista = append(lista, tr)

	}

	return lista, nil
	// return Trade{}, nil

}

func (r repo) Save(tr Trade) error {
	line1, err := json.Marshal(tr.FirstTrainerList)
	if err != nil {
		log.Printf("Error on save trade: %v\n", err)

		return err
	}
	line2, err := json.Marshal(tr.SecondTrainerList)
	if err != nil {
		log.Printf("Error on save trade: %v\n", err)

		return nil
	}
	query := fmt.Sprintf(`INSERT INTO trade (trainerone,trainertwo,created_at,fair)
		VALUES ($1,$2,$3,$4);`)
	_, err = r.DBClient.Exec(query, line1, line2, time.Now().Format("2006-01-02"), tr.Fair)
	if err != nil {
		log.Printf("Error on save trade: %v\n", err)
		return err
	}
	return nil
}

func (r repo) Get() ([]Trade, error) {
	list := []Trade{}
	q := fmt.Sprintf(`SELECT  id,trainerone,trainertwo, fair,created_at FROM trade ORDER BY created_at DESC;`)
	rows, err := r.DBClient.Query(q)
	if err != nil {
		log.Printf("Error on get trade: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		tr := Trade{}
		var line1, line2 []byte
		if err := rows.Scan(&tr.ID, &line1, &line2, &tr.Fair, &tr.CreatedAt); err != nil {
			log.Printf("Error on get trade: %v\n", err)
			return nil, err
		}
		json.Unmarshal(line1, &tr.FirstTrainerList)

		json.Unmarshal(line2, &tr.SecondTrainerList)

		list = append(list, tr)

	}
	return list, nil
}
