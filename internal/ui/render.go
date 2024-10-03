package ui

import (
	"fmt"
	"strings"

	"github.com/AsierAlaminos/NoteShell/internal/files"
	"github.com/AsierAlaminos/NoteShell/internal/model"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
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
	BannerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("93"))
	MarginStyle = lipgloss.NewStyle()
	TitleStyle = lipgloss.NewStyle().MarginLeft(2)
	ItemStyle = lipgloss.NewStyle().PaddingLeft(4)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	HelpStyle = list.DefaultStyles().HelpStyle.MarginLeft(0).MarginTop(2).PaddingLeft(4).PaddingBottom(1).Foreground(lipgloss.Color("240"))
	CategoryStyle = lipgloss.NewStyle().PaddingLeft(7).Foreground(lipgloss.Color("247"))
	SelectedCategoryStyle = lipgloss.NewStyle().PaddingLeft(7).Foreground(lipgloss.Color("250"))
	inputTextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("57"))
	listStyle = lipgloss.NewStyle().Align(lipgloss.Left)
)

type Model struct {
	List list.Model
	choice model.Idea
	Window Window
	currentState string
	errorCreatingIdea bool
	inputName textinput.Model
	inputDesc textinput.Model
	textArea textarea.Model
	writted bool
	removeIdea bool
	height int
	width int
}

func (m *Model) cleanInputs() {
	m.inputName.Reset()
	m.inputDesc.Reset()
	m.currentState = ""
}

func (m *Model) Init() tea.Cmd {
	m.inputName = textinput.New()
	m.inputName.Placeholder = "Enter name"
	m.inputName.Focus()
	m.inputDesc = textinput.New()
	m.inputDesc.Placeholder = "Enter name"

	m.textArea = textarea.New()
	m.textArea.Placeholder = ""
	m.textArea.Focus()
	return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.List.SetSize(msg.Width, 14)
		m.textArea.SetWidth(msg.Width - 10)
		m.textArea.SetHeight(msg.Height - len(m.List.Items()) * 2 - 10)
		return m, nil
	case tea.KeyMsg:
		switch m.Window {
		case List:
			if m.currentState == "" {
				if m.removeIdea {
					switch keypress := msg.String(); keypress{
					case "y":
						selected := m.List.Index()
						newItems := files.DeleteIdea(selected)
						m.List.SetItems(newItems)
						m.removeIdea = false
						return m, tea.ClearScreen
					case "n", "esc":
						m.removeIdea = false
						return m, nil
					}
				} else {
					switch keypress := msg.String(); keypress {
					case "q", "ctrl+c":
						return m, tea.Quit
					case " ":
						i, ok := m.List.SelectedItem().(model.Idea)
						if ok {
							m.choice = i
						}
						m.Window = File
						m.textArea.Focus()
						idea := m.List.SelectedItem().(model.Idea)
						m.textArea.SetValue(files.ReadDescription(idea.Name))
						return m, nil
					case "c":
						m.currentState = "name"
						m.inputName.Reset()
						m.inputName.Focus()	
						return m, nil
					case "d":
						m.removeIdea = true
						i, ok := m.List.SelectedItem().(model.Idea)
						if ok {
							m.choice = i
						}
						return m, nil
					}
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
						if name == "" || categories == "" {
							m.errorCreatingIdea = true
							m.cleanInputs()
							return m, nil
						}
						files.CreateIdea(name, strings.Split(categories, "/"))
						newIdea := model.Idea {
							Name: name,
							DescFile: fmt.Sprintf("%s.md", name),
							Categories: strings.Split(categories, "/"),
						}
						m.cleanInputs()
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
				default:
					if m.currentState == "name" {
						var cmd tea.Cmd
						m.inputName, cmd = m.inputName.Update(msg)
						return m, cmd
					} else if m.currentState == "description" {
						var cmd tea.Cmd
						m.inputDesc, cmd = m.inputDesc.Update(msg)
						return m, cmd
					}
				}
			}
		case File:
			switch keypress := msg.String(); keypress {
			case "esc":
				m.Window = List
				m.writted = true
				return m, nil
			default:
				if !m.textArea.Focused() {
					m.textArea.Focus()
				}
				var cmd tea.Cmd
				m.textArea, cmd = m.textArea.Update(msg)
				return m, cmd
			}
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}
	view := "\n" + BannerStyle.Render(files.Banner()) + "\n"
	switch m.Window {
	case List:
		view += listStyle.Render(m.List.View()) + "\n"

		if m.removeIdea {
			view += inputTextStyle.Render(fmt.Sprintf("[*] Do you want to remove %s? (y/n)", m.choice.Name)) + "\n"
		}
		if m.currentState == "name" {
			view += inputTextStyle.Render("[*] Enter the idea name...") + "\n"
			view += m.inputName.View() + "\n"
		} else if m.currentState == "description" {
			view += lipgloss.JoinVertical(lipgloss.Right, inputTextStyle.Render("[*] Enter the idea categories... (separate them by '/')")) + "\n"
			view += m.inputDesc.View() + "\n"
		}
		if m.errorCreatingIdea {
			m.errorCreatingIdea = false
			view += lipgloss.JoinVertical(lipgloss.Left, lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Render("[!] Error creating new idea..."))
		}
		help := NewDelegateKeyMap().ListHelp()
		helpText := ""
		for _, binding := range help {
			helpText += fmt.Sprintf("%s: %s   ", binding.Help().Key, binding.Help().Desc)
		}
		view += HelpStyle.Render(helpText)
		if m.writted {
			view += lipgloss.JoinVertical(lipgloss.Left, lipgloss.NewStyle().Foreground(lipgloss.Color("50")).Render("value: " + m.textArea.Value()))
			m.writted = false
			idea := m.List.SelectedItem().(model.Idea)
			files.WriteDescription(idea.Name, m.textArea.Value())
			m.textArea.Reset()
		}
		view = lipgloss.NewStyle().Width(m.width).Height(m.height).Align(lipgloss.Center, lipgloss.Center).Render(view)
	case File:
		view += "\n" + m.textArea.View() + "\n\n"
		help := NewDelegateKeyMap().FileHelp()
		helpText := ""
		for _, binding := range help {
			helpText += fmt.Sprintf("%s: %s   ", binding.Help().Key, binding.Help().Desc)
		}
		view += HelpStyle.Render(helpText)
	}
	return MarginStyle.Render(view)
}
