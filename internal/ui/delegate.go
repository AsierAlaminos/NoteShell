package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/AsierAlaminos/NoteShell/internal/model"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type IdeaDelegate struct{}

func (d IdeaDelegate) Height() int { return 1 }
func (d IdeaDelegate) Spacing() int { return 0 }
func (d IdeaDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d IdeaDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(model.Idea)
	if !ok {
		return
	}

	name := fmt.Sprintf("%d. %s", index + 1, i.Name)
	categories := i.ParseCategories()

	nameFn := ItemStyle.Render
	catFn := CategoryStyle.Render
	if index == m.Index() {
		nameFn = func(s ...string) string {
			return SelectedItemStyle.Render("> " + strings.Join(s, " "))
		}
		catFn = SelectedCategoryStyle.Render
	}

	fmt.Fprint(w, nameFn(name))
	fmt.Fprint(w, catFn("\n" + categories))
}

type delegateKeyMap struct {
	up key.Binding
	down key.Binding
	quit key.Binding
	nextWindow key.Binding
	lastWindow key.Binding
}

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding {
		d.up,
		d.down,
		d.quit,
		d.nextWindow,
		d.lastWindow,
	}
}

func (d delegateKeyMap) FullHelp() []key.Binding {
	return []key.Binding {
		d.up,
		d.down,
		d.quit,
		d.nextWindow,
		d.lastWindow,
	}
}

func NewDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		up: key.NewBinding(
			key.WithKeys("k", "up"),
			key.WithHelp("k, ↑", "up"),
		),
		down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("j, ↓", "down"),
		),
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q, ctrl+c", "quit"),
		),
		nextWindow: key.NewBinding(
			key.WithKeys("n", "n"),
			key.WithHelp("n", "next Window"),
		),
		lastWindow: key.NewBinding(
			key.WithKeys("l", "l"),
			key.WithHelp("l", "last Window"),
		),
	}
}
