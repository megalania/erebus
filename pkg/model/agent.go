package model

import (
	"time"
)

type Agent struct {
	ID        string    `db:"id" json:"ID" protobuf:"bytes,1,opt,name=id"`
	CreatedAt time.Time `db:"created_at" json:"createdAt" protobuf:"bytes,2,opt,name=createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"UpdatedAt" protobuf:"bytes,3,opt,name=updatedAt"`
}
