package model

import (
	"fmt"
)

type Idea struct {
	Name string `json:"name"`
	DescFile string `json:"descfile"`
	Categories []string `json:"categories"`
}

func (i Idea) Title() string { return i.Name }
func (i Idea) Description() string { return i.ParseCategories() }
func (i Idea) FilterValue() string { return i.Name + " " + i.ParseCategories() }
func (idea Idea) ParseCategories() string {
	parsed := ""
	if len(idea.Categories) == 0 {
		return parsed
	}

	for i := 0; i < len(idea.Categories) - 1; i++ {
		parsed += fmt.Sprintf("%s, ", idea.Categories[i])
	}

	parsed += fmt.Sprintf("%s", idea.Categories[len(idea.Categories) - 1])

	return parsed
}
