package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	FirstName   string              `json:"firstName" bson:"firstName"`
	LastName    string              `json:"lastName" bson:"lastName"`
	DisplayName string              `json:"displayName" bson:"displayName"`
	Age         int                 `json:"age" bson:"age"`
	Email       string              `json:"email" bson:"email"`
	Password    string              `json:"password" bson:"password"`
	Status      bool                `json:"status" bson:"status"`
	CreatedOn   primitive.Timestamp `json:"createdOn" bson:"createdOn"`
	UpdatedOn   primitive.Timestamp `json:"updatedOn" bson:"updatedOn"`
}
