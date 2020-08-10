package main

import (
	"assignment/handlers"
	"assignment/mapenabler"

	"github.com/gin-gonic/gin"
)

const (
	locationMatchString = `^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`
)

func main() {
	// Creating new MapEnabler instance.
	// Here Maps Api_Key , URIs to be provided.
	// What catagories are going to be fetched and size that means number of items to return for each category
	me, _ := mapenabler.New("NQLeBf6xcolqAFhQyex0sHeAILpgHqSdTT45i1ahPdI", "https://places.ls.hereapi.com", "/places/v1/discover/explore?", 1, "petrol-station", "parking-facility", "restaurant")

	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/ping", handlers.Ping())

	mapsGroup := router.Group("/v1/maps/")
	{
		mapsGroup.GET("places", handlers.FetchPlaces(me))
	}
	router.Run()
}
