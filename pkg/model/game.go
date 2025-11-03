package model

type GamesStream <-chan Game

type Achievements struct {
	Total    int `json:"total"`
	Unlocked int `json:"unlocked"`
}

type Game struct {
	Title        string        `json:"title"`
	HoursPlayed  float64       `json:"hours_played"`
	Achievements *Achievements `json:"achievements"`
	Platform     string        `json:"platform"`
}
