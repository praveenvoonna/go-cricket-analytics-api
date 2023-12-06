package model

type Player struct {
	ID         int
	Name       string
	Span       string
	Matches    int
	Innings    int
	NotOuts    int
	Runs       int
	HighScore  string
	Average    float64
	BallsFaced int
	StrikeRate float64
	Hundreds   int
	Fifties    int
	Ducks      int
}
