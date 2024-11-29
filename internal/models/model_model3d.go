package models

import "time"

type Model3D struct {
	ID        string     `bson:"_id,omitempty" json:"_id"`
	Name      string     `bson:"name" json:"name"`
	Format    string     `bson:"format" json:"format"`
	Size      int64      `bson:"size" json:"size"`
	CreatedAt time.Time  `bson:"createdAt" json:"createdAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
	Tags      []string   `bson:"tags" json:"tags"`
}
