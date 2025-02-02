package config

import "github.com/muandane/goji/pkg/models"

type Config struct {
	Types            []models.CommitType
	Scopes           []string
	Symbol           bool
	SkipQuestions    []string
	Questions        map[string]string
	SubjectMaxLength int
}

type Gitmoji struct {
	Emoji       string `json:"emoji"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Name        string `json:"name"`
}
type initConfig struct {
	Types            []Gitmoji `json:"Types"`
	Scopes           []string  `json:"Scopes"`
	Symbol           bool      `json:"Symbol"`
	SkipQuestions    []string  `json:"SkipQuestions"`
	SubjectMaxLength int       `json:"SubjectMaxLength"`
}
