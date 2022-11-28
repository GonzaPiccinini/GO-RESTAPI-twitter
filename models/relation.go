package models

// Relation struct is the relationship of one user with another
type Relation struct {
	UserID         string `bson:"user_id" json:"userId"`
	UserRelationID string `bson:"user_relation_id" json:"userRelationId"`
}