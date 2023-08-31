package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gaurilab/ImageStoreService/database"
	"github.com/gaurilab/ImageStoreService/model"
	"github.com/gin-gonic/gin"
)

func RenderHomePage(c *gin.Context) {

	//var albums []model.Album
	albums, err := database.GetAllAlbum()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to  retrieve albums")
		return
	}
	for _, album := range albums {
		fmt.Println(album.AlbumName, album)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"albums": albums,
	})
}

func RenderCreateAlbumPage(c *gin.Context) {
	c.HTML(http.StatusOK, "create-album.html", nil)
}

func CreateAlbum(c *gin.Context) {

	albumName := c.PostForm("album-name")
	var album model.Album

	album.AlbumName = albumName
	album.CreatedAt = time.Now()
	fmt.Println(" \nalbumname: %s, album: %s", albumName, album)
	// Insert album to MongoDB
	err := database.CreateAlbum(album)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to create album")
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func GetAlbum(c *gin.Context) {
	albumIDStr := c.Param("albumID")

	// Check if the album exists
	var in model.Album
	in.ID = albumIDStr
	fmt.Println(" \nalbumid in: ", in)
	album, err := database.GetAlbumById(in)
	if err != nil {
		c.String(http.StatusNotFound, "Album not found")
		return
	}

	fmt.Println(" \nalbumid: ", album)
	images, err := database.GetAllImage(album)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch images")
		return
	}

	c.HTML(http.StatusOK, "album.html", gin.H{
		"album":  album,
		"images": images,
	})
}

func DeleteAlbum(c *gin.Context) {
	albumIDStr := c.Param("albumID")
	var in model.Album
	in.ID = albumIDStr
	err := database.DeleteAlbumById(in)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete album")
		return
	}
	err = database.DeleteImageByAlbum(in)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to delete associated images")
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}
