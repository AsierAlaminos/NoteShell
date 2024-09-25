package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AsierAlaminos/NoteShell/internal/model"
	"github.com/charmbracelet/bubbles/list"
	"golang.org/x/term"
)

func CreateIdeaList(ideasPath string) []list.Item {

	byteValue, err := os.ReadFile(ideasPath)
	if err != nil {
		fmt.Println("[!] Exiting... (invalid json file)")
		os.Exit(1)
	}

	var items []list.Item
	var ideas []model.Idea

	if err := json.Unmarshal(byteValue, &ideas); err != nil {
		fmt.Println("[!] Error in json file")
	}

	for _, i := range ideas {
		items = append(items, i)
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
