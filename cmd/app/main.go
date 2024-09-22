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
	/*items := []list.Item {
		ui.Idea{Name: "name1", Categories: []string{"cat1", "cat2"}},
		ui.Idea{Name: "name2", Categories: []string{"cat3", "cat4"}},
		ui.Idea{Name: "name3", Categories: []string{"cat5", "cat6"}},
	}*/
	items := utils.CreateIdeaList("/home/asmus/.noteshell/ideas")

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

	/*fmt.Println("[*] Printing directories")
	files.ReadDirs("/home/asmus/.noteshell/ideas")
	fmt.Println("[*] Creating directories")
	for i := 0; i < 10; i++ {
		ideaName := fmt.Sprintf("idea%d", i)
		files.CreateIdeaFiles(ideaName)
	}*/


}
