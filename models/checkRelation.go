package models

// CheckRelation model returns a status value (true) according to the relation between 2 users
type CheckRelation struct {
	Status bool `json:"status"`
}