package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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

func FilterIdeas(value string, ideas []list.Item) []list.Item {
	var filteredIdeas []list.Item

	for _, item := range ideas {
		idea := item.(model.Idea)
		if strings.ToLower(idea.Name)[:len(value)] ==  strings.ToLower(value) {
			filteredIdeas = append(filteredIdeas, item)
		}
	}

	return filteredIdeas
}
