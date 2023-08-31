package model

import "time"

type Image struct {
	ID         string    `bson:"_id,omitempty"`
	AlbumID    string    `bson:"album_id"`
	ImageData  []byte    `bson:"image_data"`
	UploadTime time.Time `bson:"upload_time"`
	ImageName  string    `bson:"name"`
}

type Album struct {
	ID        string    `bson:"_id,omitempty"`
	AlbumName string    `bson:"name"`
	CreatedAt time.Time `bson:"created_at"`
}
