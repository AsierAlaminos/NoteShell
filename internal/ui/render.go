package ui

import (
	"fmt"

	"github.com/AsierAlaminos/NoteShell/internal/model"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1).Foreground(lipgloss.Color("240"))
	CategoryStyle = lipgloss.NewStyle().PaddingLeft(7).Foreground(lipgloss.Color("247"))
	SelectedCategoryStyle = lipgloss.NewStyle().PaddingLeft(7).Foreground(lipgloss.Color("247"))
)


type Model struct {
	List list.Model
	choice model.Idea
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
			case "q", "ctrl+c":
				return m, tea.Quit
			case " ", "enter":
				i, ok := m.List.SelectedItem().(model.Idea)
				if ok {
					m.choice = i
				}
				return m, nil
		}
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	help := NewDelegateKeyMap().ShortHelp()
	helpText := ""
	for _, binding := range help {
		helpText += fmt.Sprintf("%s: %s   ", binding.Help().Key, binding.Help().Desc)
	}
	return "\n" + m.List.View() + HelpStyle.Render(helpText) + "\n"
}
