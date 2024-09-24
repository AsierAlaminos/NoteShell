package ui

import (
	"fmt"
	"strings"

	"github.com/AsierAlaminos/NoteShell/internal/files"
	"github.com/AsierAlaminos/NoteShell/internal/model"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Window int

const (
	List Window = iota
	File
	Todo
)

func (s Window) String() string {
	switch s {
	case List:
		return "List"
	case File:
		return "File"
	case Todo:
		return "Todo"
	default:
		return "unknown"
	}
}

var (
	TitleStyle = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1).Foreground(lipgloss.Color("240"))
	CategoryStyle = lipgloss.NewStyle().PaddingLeft(7).Foreground(lipgloss.Color("247"))
	SelectedCategoryStyle = lipgloss.NewStyle().PaddingLeft(7).Foreground(lipgloss.Color("250"))
)


type Model struct {
	List list.Model
	choice model.Idea
	Window Window
	currentState string
	inputName textinput.Model
	inputDesc textinput.Model
}

func (m *Model) Init() tea.Cmd {
	m.inputName = textinput.New()
	m.inputName.Placeholder = "Enter name"
	m.inputName.Focus()
	m.inputDesc = textinput.New()
	m.inputDesc.Placeholder = "Enter name"
	return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		if m.currentState == "" {

			switch keypress := msg.String(); keypress {
			case "q", "ctrl+c":
				return m, tea.Quit
			case " ":
				i, ok := m.List.SelectedItem().(model.Idea)
				if ok {
					m.choice = i
				}
				return m, nil
			case "n":
				if m.Window < 2 {
					m.Window++
				}
				return m, nil
			case "l":
				if m.Window > 0 {
					m.Window--
				}
				return m, nil
			case "c":
				m.currentState = "name"
				m.inputName.Reset()
				m.inputName.Focus()	
				return m, nil
			}
		} else {
			switch keypress := msg.String(); keypress {
			case "enter":
				if m.currentState == "name" {
					m.currentState = "description"
					m.inputDesc.Reset()
					m.inputDesc.Focus()	
					return m, nil
				} else if m.currentState == "description" {
					name := m.inputName.Value()
					categories := m.inputDesc.Value()
					files.CreateIdeaFiles(name, strings.Split(categories, "/"))
					newIdea := model.Idea {
						Name: name,
						DescFile: fmt.Sprintf("%s.md", name),
						Categories: strings.Split(categories, "/"),
					}
					m.inputName.Reset()
					m.inputDesc.Reset()
					m.currentState = ""
					
					items := m.List.Items()
					items = append(items, newIdea)
					m.List.SetItems(items)
					return m, tea.ClearScreen
				}
			case "esc":
				m.currentState = ""
				m.inputName.Reset()
				m.inputDesc.Reset()
				return m, nil
			}
		}
	}

	if m.currentState == "name" {
		var cmd tea.Cmd
		m.inputName, cmd = m.inputName.Update(msg)
		return m, cmd
	} else if m.currentState == "description" {
		var cmd tea.Cmd
		m.inputDesc, cmd = m.inputDesc.Update(msg)
		return m, cmd
	}
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	help := NewDelegateKeyMap().ShortHelp()
	helpText := ""
	for _, binding := range help {
		helpText += fmt.Sprintf("%s: %s   ", binding.Help().Key, binding.Help().Desc)
	}
	view := "\n" + m.List.View() + HelpStyle.Render(helpText) + "\n"
	if m.currentState == "name" {
		view += lipgloss.JoinVertical(lipgloss.Bottom, lipgloss.NewStyle().Foreground(lipgloss.Color("57")).Render("[*] Enter the idea name...")) + "\n"
		view += m.inputName.View()
	} else if m.currentState == "description" {
		view += lipgloss.JoinVertical(lipgloss.Right, lipgloss.NewStyle().Foreground(lipgloss.Color("57")).Render("[*] Enter the idea categories... (separate them by '/')")) + "\n"
		view += m.inputDesc.View()
	}
	return view
}
