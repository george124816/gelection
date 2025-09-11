package candidate

type Candidate struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	ElectionId int64  `json:"election_id"`
}
