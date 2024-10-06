# 📓 NoteShell

**NoteShell** is a TUI program to take notes in a fuild and organized way on the terminal.

## Features
- [x] Create ideas
- [x] Read ideas
- [X] Update ideas
- [x] Delete ideas
- [X] Idea description
- [ ] Filter idea
- [ ] Kanban todo list
- [ ] File encryption

## Libraries
This is a project that uses Charm libraries:

### [Bubble Tea](https://github.com/charmbracelet/bubbletea)

### [Lipgloss](https://github.com/charmbracelet/lipgloss)

### [Bubbles](https://github.com/charmbracelet/bubbles)

## Example

![NoteShell](assets/noteshell.gif)

## Instalation

Clone the repository:
```bash
git clone https://github.com/AsierAlaminos/NoteShell.git
```

Enter the repository:
```bash
cd NoteShell
```

Install dependencies:
```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/lipgloss
```

Update mod.go
```bash
go mod tidy
```

Create program:
```bash
make
```

Run:
```bash
./noteshell
```
