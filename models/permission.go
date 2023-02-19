package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Permission struct {
	Id                    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Permission            string             `json:"firstName" bson:"firstName"`
	PermissionCode        int                `json:"lastName" bson:"lastName"`
	ValidGroup            []string           `json:"validGroup" bson:"validGroup"`
	PermissionDescription string             `json:"permissionDescription" bson:"permissionDescription"`
	ParentPermission      int                `json:"parentPermission" bson:"parentPermission"`
}
