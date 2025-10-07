package model

import "time"

type Vote struct {
	Id          string
	InsertedAt  time.Time
	ElectionId  int
	CandidateId int
}
