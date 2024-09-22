package utils

import (
	"fmt"

	"github.com/AsierAlaminos/NoteShell/internal/files"
	"github.com/charmbracelet/bubbles/list"
)

func CreateIdeaList(ideasPath string) []list.Item {
	var items []list.Item

	dirs := files.ReadDirs(ideasPath)

	for _,d := range dirs {
		jsonPath := fmt.Sprintf("%s/%s/%s.json", ideasPath, d, d)
		idea := files.ReadJsonIdea(jsonPath)
		items = append(items, idea)
	}

	return items
}
