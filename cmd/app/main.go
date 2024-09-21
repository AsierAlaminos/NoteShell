package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/AsierAlaminos/cmd/files"
	"github.com/AsierAlaminos/NoteShell/internal/ui"
)

func main() {
	items := []list.Item {
		ui.Idea{Name: "name1", Categories: []string{"cat1", "cat2"}},
		ui.Idea{Name: "name2", Categories: []string{"cat3", "cat4"}},
		ui.Idea{Name: "name3", Categories: []string{"cat5", "cat6"}},
	}

	l := list.New(items, ui.IdeaDelegate{}, 20, 4 + len(items))
	l.Title = "Idea list"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.Styles.Title = ui.TitleStyle
	l.Styles.PaginationStyle = ui.PaginationStyle
	l.Styles.HelpStyle = ui.HelpStyle

	m := ui.Model{List: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program: ", err)
		os.Exit(1)
	}
}
