package trade

import "time"

type Trade struct {
	ID                int        `json:"id"`
	FirstTrainerList  []*Pokemon `json:"first_trainer_list"`
	SecondTrainerList []*Pokemon `json:"second_trainer_list"`
	Fair              bool       `json:"fair,omitempty"`
	CreatedAt         time.Time  `json:"created_at,omitempty" `
}

type Pokemon struct {
	BaseExperience int    `json:"base_experience"`
	ID             int    `json:"id"`
	Name           string `json:"name"`
}
