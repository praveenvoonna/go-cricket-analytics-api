package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/praveenvoonna/go-cricket-analytics-api/model"
)

var (
	PlayersByYear  map[string][]model.Player
	MostRunsByYear map[string]model.Player
)

func ReadCSVLoadMaps(filePath string) ([]model.Player, error) {
	var players []model.Player

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	PlayersByYear = make(map[string][]model.Player)
	MostRunsByYear = make(map[string]model.Player)

	for i, line := range lines {
		if i == 0 {
			continue
		}

		matches, _ := strconv.Atoi(line[3])
		innings, _ := strconv.Atoi(line[4])
		notOuts, _ := strconv.Atoi(line[5])
		runs, _ := strconv.Atoi(line[6])
		average, _ := strconv.ParseFloat(line[8], 64)
		ballsFaced, _ := strconv.Atoi(line[9])
		strikeRate, _ := strconv.ParseFloat(line[10], 64)
		hundreds, _ := strconv.Atoi(line[11])
		fifties, _ := strconv.Atoi(line[12])
		ducks, _ := strconv.Atoi(line[13])
		player := model.Player{
			ID:         i,
			Name:       line[1],
			Span:       line[2],
			Matches:    matches,
			Innings:    innings,
			NotOuts:    notOuts,
			Runs:       runs,
			HighScore:  line[7],
			Average:    average,
			BallsFaced: ballsFaced,
			StrikeRate: strikeRate,
			Hundreds:   hundreds,
			Fifties:    fifties,
			Ducks:      ducks,
		}
		players = append(players, player)

		yearsActive := strings.Split(player.Span, "-")
		for _, year := range yearsBetween(yearsActive[0], yearsActive[1]) {
			PlayersByYear[year] = append(PlayersByYear[year], player)
			updateMostRuns(year, player)
		}
	}

	return players, nil
}

func updateMostRuns(year string, player model.Player) {
	existingPlayer, exists := MostRunsByYear[year]
	if !exists || player.Runs > existingPlayer.Runs {
		MostRunsByYear[year] = player
	}
}

func yearsBetween(start, end string) []string {
	startYear, _ := strconv.Atoi(start)
	endYear, _ := strconv.Atoi(end)

	var years []string
	for i := startYear; i <= endYear; i++ {
		years = append(years, strconv.Itoa(i))
	}
	return years
}
