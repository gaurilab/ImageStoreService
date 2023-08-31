package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gaurilab/ImageStoreService/database"

	"github.com/gaurilab/ImageStoreService/model"
	"github.com/gin-gonic/gin"
)

func DeleteImage(c *gin.Context) {
	albumIDStr := c.Param("albumID")
	imageIDStr := c.Param("imageID")

	fmt.Println("\n Image id : ")
	fmt.Println(albumIDStr, imageIDStr)

	// Delete the image from the database
	var album model.Album
	album.ID = albumIDStr
	var image model.Image
	image.ID = imageIDStr

	err := database.DeleteImageById(album, image)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete image")
		return
	}

	c.Redirect(http.StatusSeeOther, "/albums/"+albumIDStr)
}

func RenderImage(c *gin.Context) {
	imageIDStr := c.Param("id")
	fmt.Println("\n Image id : ")
	fmt.Println(imageIDStr)

	var in model.Image
	in.ID = imageIDStr

	image, err := database.GetImageById(in)
	if err != nil {
		c.String(http.StatusNotFound, "Image not found db")
		return
	}

	c.Data(http.StatusOK, "image/jpeg", image.ImageData)
}

func RenderUploadImagePage(c *gin.Context) {
	albumID := c.Param("albumID")
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"albumID": albumID,
	})
}

func UploadImage(c *gin.Context) {
	albumIDStr := c.Param("albumID")

	// Check if the album exists
	var album model.Album
	album.ID = albumIDStr
	fmt.Println("UploadImage: ", albumIDStr)

	albumInDb, err := database.GetAlbumById(album)
	if err != nil {
		c.String(http.StatusNotFound, "Album not found")
		return
	}

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	defer file.Close()

	imageName := c.PostForm("image-name")
	// Read image data into a byte slice
	var imageData bytes.Buffer

	_, err = io.Copy(&imageData, file)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to read image data")
		return
	}

	// Insert image data to MongoDB with the album ID
	err = database.CreateImage(albumInDb, model.Image{
		AlbumID:    albumIDStr,
		ImageData:  imageData.Bytes(),
		UploadTime: time.Now(),
		ImageName:  imageName,
	})

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to store image data")
		return
	}

	c.Redirect(http.StatusSeeOther, "/albums/"+albumIDStr)
}
