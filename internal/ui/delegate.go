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
	createIdea key.Binding
	updateIdea key.Binding
	deleteIdea key.Binding
	save key.Binding
	quitNotSaving key.Binding
	quitSaving key.Binding
	filter key.Binding
	resetList key.Binding
}

func (d delegateKeyMap) ListHelp() []key.Binding {
	return []key.Binding {
		d.up,
		d.down,
		d.createIdea,
		d.updateIdea,
		d.deleteIdea,
		d.quit,
		d.filter,
		d.resetList,
	}
}

func (d delegateKeyMap) FileHelp() []key.Binding {
	return []key.Binding {
		d.save,
		d.quitSaving,
		d.quitNotSaving,
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
		createIdea: key.NewBinding(
			key.WithKeys("c", "c"),
			key.WithHelp("c", "create new idea"),
		),
		updateIdea: key.NewBinding(
			key.WithKeys("u", "u"),
			key.WithHelp("u", "update idea"),
		),
		deleteIdea: key.NewBinding(
			key.WithKeys("d", "d"),
			key.WithHelp("d", "delete idea"),
		),
		quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q, ctrl+c", "quit"),
		),
		save: key.NewBinding(
			key.WithKeys("esc", "esc"),
			key.WithHelp("esc", "save file"),
		),
		quitSaving: key.NewBinding(
			key.WithKeys("ctrl+w", "ctrl+w"),
			key.WithHelp("ctrl+w", "save and exit file"),
		),
		quitNotSaving: key.NewBinding(
			key.WithKeys("ctrl+q", "ctrl+q"),
			key.WithHelp("ctrl+q", "save and exit file"),
		),
		filter: key.NewBinding(
			key.WithKeys("f", "f"),
			key.WithHelp("f", "filter list"),
		),
		resetList: key.NewBinding(
			key.WithKeys("r", "r"),
			key.WithHelp("r", "reset list"),
		),
	}
}
