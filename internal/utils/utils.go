package utils

import (
	"fmt"

	"github.com/AsierAlaminos/NoteShell/internal/files"
	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/term"
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

func getTerminalSize() (width, height int){

	w, h, err := term.GetSize(0)
	if err != nil {
		return 
	}
	return w, h
}
