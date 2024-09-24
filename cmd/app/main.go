package main

import (
	"fmt"
	"os"

	"github.com/AsierAlaminos/NoteShell/internal/ui"
	"github.com/AsierAlaminos/NoteShell/internal/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	items := utils.CreateIdeaList("/home/asmus/.noteshell/ideas")

	l := list.New(items, ui.IdeaDelegate{}, 20, len(items) * 2)
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
