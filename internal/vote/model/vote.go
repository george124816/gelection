package model

import "time"

type Vote struct {
	Id          string    `json:"id"`
	InsertedAt  time.Time `json:"inserted_at"`
	ElectionId  int       `json:"election_id"`
	CandidateId int       `json:"candidate_id"`
}
