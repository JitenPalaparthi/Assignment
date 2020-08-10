// Package handlers is to define/create handler.
// At preset there is only one handler
package handlers

import (
	"assignment/mapenabler"
	"assignment/models"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	// Regular expression for location string
	locationMatchString = `^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`
)

// Ping is a test call that retuns pong on success
func Ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

// FetchPlaces fetches places based on catagories of mapenabler type
// It is a concurrent function
func FetchPlaces(me *mapenabler.MapEnabler) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			chanResult := make(chan *models.Result)                        // Channel is used to send data of each fetch
			chanSignal := make(chan bool)                                  // Signal to close channel and ensure to fetch all data
			var Results []models.Result                                    // Each fetch request is added to this slice
			location := c.Query("loc")                                     // Must be a valid location and location is mandatory
			pass, err := regexp.MatchString(locationMatchString, location) // to validate location parameters
			if !pass || err != nil {
				c.JSON(http.StatusBadRequest, "given Location values(longitude or latitude) are wrong")
				c.Abort()
				return
			}
			for _, cat := range me.Categories {
				go me.FetchMapsDataWithChan(location, cat, chanResult) // Concurrent FetchMapsdata . upon fetch data is passed to chanResult channel
			}
			go func() {
				counter := 0
				for {
					select {
					case result := <-chanResult:
						if result != nil {
							Results = append(Results, *result) //When data is received from the channel and added to the slice
						}
						counter++
					default:
						if counter == len(me.Categories) {
							close(chanResult)  // When to close the channel it is trickey part. In this scenario  it is based on number of catagories
							chanSignal <- true // Once channel is closed can signal for the further process.. Can use workGroups as well but this is simple way
							return
						}
					}

				}
			}()
			<-chanSignal // Two purposes 1. Ensure all routines are given data 2. Stop the below code untle everything is fetched
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, Results)
		}
	}
}

// FetchPlacesByQueries fetches places based on query string parameters.
// It is a concurrent function
func FetchPlacesByQueries() func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method == "GET" {
			chanResult := make(chan *models.Result) // Channel is used to send data of each fetch
			chanSignal := make(chan bool)           // Signal to close channel and ensure to fetch all data
			var Results []models.Result             // Each fetch request is added to this slice

			location := c.Query("loc")                                     // Must be a valid location and location is mandatory
			pass, err := regexp.MatchString(locationMatchString, location) // to validate location parameters
			if !pass || err != nil {
				c.JSON(http.StatusBadRequest, "given Location values(longitude or latitude) are wrong")
				c.Abort()
				return
			}

			categories := c.QueryArray("cat")

			if len(categories) <= 0 {
				c.JSON(http.StatusBadRequest, "Categorie(s) to be provided")
				c.Abort()
				return
			}

			apiKey := c.Query("api_key")
			if apiKey == "" {
				c.JSON(http.StatusBadRequest, "api_key to be provided")
				c.Abort()
				return
			}

			size := c.Query("size")
			if size == "" {
				c.JSON(http.StatusBadRequest, "size to be provided")
				c.Abort()
				return
			}

			_size, err := strconv.Atoi(size)
			if err != nil {
				c.JSON(http.StatusBadRequest, "invalid size parameter")
				c.Abort()

			}
			me, err := mapenabler.New(apiKey, "https://places.ls.hereapi.com", "/places/v1/discover/explore?", _size)
			if err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				c.Abort()

			}
			me.Categories = categories

			for _, cat := range me.Categories {
				go me.FetchMapsDataWithChan(location, cat, chanResult) // Concurrent FetchMapsdata . upon fetch data is passed to chanResult channel
			}
			go func() {
				counter := 0
				for {
					select {
					case result := <-chanResult:
						if result != nil {
							Results = append(Results, *result) //When data is received from the channel and added to the slice
						}
						counter++
					default:
						if counter == len(me.Categories) {
							close(chanResult)  // When to close the channel it is trickey part. In this scenario  it is based on number of catagories
							chanSignal <- true // Once channel is closed can signal for the further process.. Can use workGroups as well but this is simple way
							return
						}
					}

				}
			}()
			<-chanSignal // Two purposes 1. Ensure all routines are given data 2. Stop the below code untle everything is fetched
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, Results)
		}
	}
}
