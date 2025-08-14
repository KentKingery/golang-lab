package main

import (
	"io"
	"log/slog"
	"net/http"

	badger "github.com/dgraph-io/badger/v4"
	"github.com/gin-gonic/gin"
)

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	slog.Info("Startup")

	opts := badger.DefaultOptions("")
	opts.Logger = nil

	db, err := badger.Open(badger.DefaultOptions("./data"))
	if err != nil {
		slog.Error(err.Error())
	}

	defer db.Close()

	gin.DefaultWriter = io.Discard
	router := gin.Default()

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusOK) // Return 200 OK
	})

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8181")

	slog.Info("Shutdown")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	slog.Info("getAlbums called")
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	slog.Info("postAlbums called")
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
