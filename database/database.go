package database

import (
	"context"
	"fmt"

	"github.com/gaurilab/ImageStoreService/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "imagestore"
const imagesCollName = "images"
const albumsCollName = "albums"

var (
	Client          *mongo.Client
	Db              *mongo.Database
	AlbumCollection *mongo.Collection
	ImageCollection *mongo.Collection
)

func ConnectDB() error {
	clientOptions := options.Client().ApplyURI("mongodb://10.152.183.179:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	Client = client
	Db = client.Database(dbName)

	// Create or get collections
	AlbumCollection = Db.Collection(albumsCollName)
	ImageCollection = Db.Collection(imagesCollName)

	return nil
}

func CreateCollections() {
	AlbumCollection = Db.Collection(albumsCollName)
	ImageCollection = Db.Collection(imagesCollName)

}

func CreateAlbum(album model.Album) error {

	_, err := AlbumCollection.InsertOne(context.Background(), album)
	if err != nil {

		return err
	}
	return nil
}

func DeleteAlbumByName(album model.Album) error {
	// Delete associated blogs
	_, err := AlbumCollection.DeleteOne(context.TODO(), bson.M{"name": album.AlbumName})
	if err != nil {
		return err
	}
	return nil

}

func DeleteAlbumById(in model.Album) error {
	var albumIDStr = in.ID
	albumID, err := primitive.ObjectIDFromHex(albumIDStr)
	if err != nil {

		return err
	}
	// Delete associated blogs
	_, err = AlbumCollection.DeleteOne(context.Background(), bson.M{"_id": albumID})
	if err != nil {
		return err
	}
	return nil
}

func GetAlbumByName(in model.Album) (model.Album, error) {
	var album model.Album
	err := AlbumCollection.FindOne(context.Background(), bson.M{"name": in.AlbumName}).Decode(&album)
	if err != nil {
		return model.Album{}, err
	}
	return album, nil
}

func GetAlbumById(in model.Album) (model.Album, error) {

	albumID, err := primitive.ObjectIDFromHex(in.ID)
	if err != nil {
		fmt.Println(" GetAlbumById failed converting ", in.ID, err)
		return model.Album{}, err
	}

	var album model.Album
	err = AlbumCollection.FindOne(context.Background(), bson.M{"_id": albumID}).Decode(&album)
	if err != nil {
		fmt.Println(" GetAlbumById failed  ", in.ID, err)
		return model.Album{}, err
	}
	return album, nil
}

func GetAllAlbum() ([]model.Album, error) {
	var albums []model.Album
	cursor, err := AlbumCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &albums); err != nil {
		return nil, err
	}
	return albums, nil
}

func CreateImage(album model.Album, image model.Image) error {
	var albumIDStr = album.ID
	albumID, err := primitive.ObjectIDFromHex(albumIDStr)
	if err != nil {

		return err
	}

	// Check if the album exists
	var albuminDb model.Album
	err = AlbumCollection.FindOne(context.Background(), bson.M{"_id": albumID}).Decode(&albuminDb)
	if err != nil {
		return err
	}

	// Insert image data to MongoDB with the album ID
	_, err = ImageCollection.InsertOne(context.Background(), image)
	if err != nil {
		return err
	}

	return nil
}

func DeleteImageByAlbum(album model.Album) error {
	var albumIDStr = album.ID

	// Delete image data to MongoDB with the album ID
	_, err := ImageCollection.DeleteMany(context.Background(), bson.M{"album_id": albumIDStr})
	if err != nil {
		return err
	}

	return nil
}

func DeleteImageById(album model.Album, image model.Image) error {
	var albumIDStr = album.ID
	albumID, err := primitive.ObjectIDFromHex(albumIDStr)
	if err != nil {

		return err
	}

	// Check if the album exists
	var albuminDb model.Album
	err = AlbumCollection.FindOne(context.Background(), bson.M{"_id": albumID}).Decode(&albuminDb)
	if err != nil {
		return err
	}

	imageId, err := primitive.ObjectIDFromHex(image.ID)
	if err != nil {
		return err
	}

	// Delete image data from MongoDB with the album ID
	_, err = ImageCollection.DeleteOne(context.Background(), bson.M{"_id": imageId, "album_id": albuminDb.ID})
	if err != nil {
		return err
	}

	return nil
}

func GetImageByName(in model.Album, in2 model.Image) (model.Image, error) {

	var albumIDStr = in.ID
	albumID, err := primitive.ObjectIDFromHex(albumIDStr)
	if err != nil {

		return model.Image{}, err
	}

	// Check if the album exists
	var albuminDb model.Album
	err = AlbumCollection.FindOne(context.Background(), bson.M{"_id": albumID}).Decode(&albuminDb)
	if err != nil {
		return model.Image{}, err
	}

	filter := bson.M{}
	filter["album_id"] = albuminDb.ID
	filter["name"] = in2.ImageName

	// Insert image data to MongoDB with the album ID
	var image model.Image
	err = ImageCollection.FindOne(context.Background(), filter).Decode(&image)
	if err != nil {
		return model.Image{}, err
	}

	return image, nil

}

func GetImageById(in2 model.Image) (model.Image, error) {

	imageId, err := primitive.ObjectIDFromHex(in2.ID)
	if err != nil {
		return model.Image{}, err
	}

	filter := bson.M{}
	filter["_id"] = imageId

	// Insert image data to MongoDB with the album ID
	var image model.Image
	err = ImageCollection.FindOne(context.Background(), filter).Decode(&image)
	if err != nil {
		return model.Image{}, err
	}

	return image, nil
}

func GetImageById2(in model.Album, in2 model.Image) (model.Image, error) {
	var albumIDStr = in.ID
	albumID, err := primitive.ObjectIDFromHex(albumIDStr)
	if err != nil {
		return model.Image{}, err
	}

	// Check if the album exists
	var albuminDb model.Album
	err = AlbumCollection.FindOne(context.Background(), bson.M{"_id": albumID}).Decode(&albuminDb)
	if err != nil {
		return model.Image{}, err
	}

	imageId, err := primitive.ObjectIDFromHex(in2.ID)
	if err != nil {
		return model.Image{}, err
	}

	filter := bson.M{}
	filter["album_id"] = albuminDb.ID
	filter["_id"] = imageId

	// Insert image data to MongoDB with the album ID
	var image model.Image
	err = ImageCollection.FindOne(context.Background(), filter).Decode(&image)
	if err != nil {
		return model.Image{}, err
	}

	return image, nil
}

// TODO: implement paging

func GetAllImage(in model.Album) ([]model.Image, error) {

	filter := bson.M{}
	filter["album_id"] = in.ID

	var images []model.Image
	cursor, err := ImageCollection.Find(context.Background(), filter)
	if err != nil {
		fmt.Println("GetAllImage : Could not find album ", in.ID)
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &images); err != nil {
		return nil, err
	}
	fmt.Println("GetAllImage : found images")
	return images, nil

}
