package models

import "time"

type Model3D struct {
	ID        string     `bson:"_id,omitempty"`
	Name      string     `bson:"name"`
	Format    string     `bson:"format"`
	Size      int64      `bson:"size"`
	CreatedAt time.Time  `bson:"createdAt"`
	DeletedAt *time.Time `bson:"deletedAt,omitempty"`
	Tags      []string   `bson:"tags"`
}
