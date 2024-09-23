package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/AsierAlaminos/NoteShell/internal/ui"
	"github.com/AsierAlaminos/NoteShell/internal/utils"
)



func main() {
	items := utils.CreateIdeaList("/home/asmus/.noteshell/ideas")

	l := list.New(items, ui.IdeaDelegate{}, 20, 4 + len(items))
	l.Title = "Idea list"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = ui.TitleStyle
	l.Styles.PaginationStyle = ui.PaginationStyle
	l.Styles.HelpStyle = ui.HelpStyle

	m := ui.Model{List: l, Window: ui.List}

	program := tea.NewProgram(&m, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
