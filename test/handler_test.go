package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/praveenvoonna/go-cricket-analytics-api/csv"
	"github.com/praveenvoonna/go-cricket-analytics-api/handler"
	"github.com/praveenvoonna/go-cricket-analytics-api/model"
)

func TestMostRunsByYear_Success(t *testing.T) {
	csv.MostRunsByYear = map[string]model.Player{
		"2012": model.Player{Name: "SR Tendulkar (INDIA)", Runs: 18426},
	}

	router := gin.Default()
	router.GET("/players/most_runs", handler.MostRunsByYear)

	req, _ := http.NewRequest("GET", "/players/most_runs?year=2012", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	assert.Contains(t, resp.Body.String(), `"playerWithMost":{"ID":0,"Name":"SR Tendulkar (INDIA)","Span":"","Matches":0,"Innings":0,"NotOuts":0,"Runs":18426,"HighScore":"","Average":0,"BallsFaced":0,"StrikeRate":0,"Hundreds":0,"Fifties":0,"Ducks":0},"year":"2012"`)

}

func TestMostRunsByYear_NotFound(t *testing.T) {
	csv.MostRunsByYear = make(map[string]model.Player)

	router := gin.Default()
	router.GET("/players/most_runs", handler.MostRunsByYear)

	req, _ := http.NewRequest("GET", "/players/most_runs?year=2021", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestActivePlayersByYear_Success(t *testing.T) {
	csv.PlayersByYear = map[string][]model.Player{
		"2012": {
			{ID: 1, Name: "Player 1", Runs: 100},
			{ID: 2, Name: "Player 2", Runs: 150},
		},
	}

	router := gin.Default()
	router.GET("/players/active", handler.ActivePlayersByYear)

	req, _ := http.NewRequest("GET", "/players/active?year=2012", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	actualResponse := resp.Body.String()
	t.Logf("Actual response: %s", actualResponse)

	assert.Equal(t, http.StatusOK, resp.Code)

	expectedResponse := `{"activePlayers":[{"ID":1,"Name":"Player 1","Span":"","Matches":0,"Innings":0,"NotOuts":0,"Runs":100,"HighScore":"","Average":0,"BallsFaced":0,"StrikeRate":0,"Hundreds":0,"Fifties":0,"Ducks":0},{"ID":2,"Name":"Player 2","Span":"","Matches":0,"Innings":0,"NotOuts":0,"Runs":150,"HighScore":"","Average":0,"BallsFaced":0,"StrikeRate":0,"Hundreds":0,"Fifties":0,"Ducks":0}],"year":"2012"}`
	assert.JSONEq(t, expectedResponse, resp.Body.String())
}

func TestActivePlayersByYear_NotFound(t *testing.T) {
	csv.PlayersByYear = make(map[string][]model.Player)

	router := gin.Default()
	router.GET("/players/active", handler.ActivePlayersByYear)

	req, _ := http.NewRequest("GET", "/players/active?year=2021", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	expectedResponse := `{"error":"No player data found for the given year","year":"2021"}`
	assert.JSONEq(t, expectedResponse, resp.Body.String())
}
