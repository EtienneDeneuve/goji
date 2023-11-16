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
