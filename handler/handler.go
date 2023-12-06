package handler

import (
	"net/http"

	"github.com/praveenvoonna/go-cricket-analytics-api/csv"

	"github.com/gin-gonic/gin"
)

func MostRunsByYear(c *gin.Context) {
	year := c.Query("year")

	playerWithMostRuns, exists := csv.MostRunsByYear[year]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No player data found for the given year",
			"year":  year,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"year":           year,
		"playerWithMost": playerWithMostRuns,
	})
}

func ActivePlayersByYear(c *gin.Context) {
	year := c.Query("year")

	activePlayers, exists := csv.PlayersByYear[year]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No player data found for the given year",
			"year":  year,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"year":          year,
		"activePlayers": activePlayers,
	})
}

func SetupHandlers() *gin.Engine {
	router := gin.Default()

	router.GET("/players/most_runs", MostRunsByYear)
	router.GET("/players/active", ActivePlayersByYear)

	return router
}
