package main

import (
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

	db, err := badger.Open(badger.DefaultOptions("./data"))
	if err != nil {
		slog.Error(err.Error())
	}

	defer db.Close()

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")

	slog.Info("Shutdown")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
