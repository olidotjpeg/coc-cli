package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type Attribute struct {
	Type  string
	Value string
}

func (i Attribute) Title() string       { return i.Type }
func (i Attribute) Description() string { return i.Value }
func (i Attribute) FilterValue() string { return i.Type }

type model struct {
	list     list.Model
	editing  bool
	editText string
}

func (m model) Init() tea.Cmd {
	return nil
}

var newText string

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			if !m.editing {
				m.editing = true
			} else {
				m.editing = false
				selectedItem := m.list.SelectedItem().(Attribute)
				selectedIndex := m.list.Index()
				selectedItem.Value = m.editText
				// Clear the editText after saving
				m.editText = ""

				return m, m.list.SetItem(selectedIndex, selectedItem)

			}
			return m, nil
		}
		if m.editing {
			switch msg.String() {
			case "backspace":
				// Handle backspace when editing
				if len(m.editText) > 0 {
					m.editText = m.editText[:len(m.editText)-1]
				}
			default:
				// Capture user input while editing
				m.editText += msg.String()
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var inputField string

	if m.editing {
		// Show the input field when editing
		selectedItem := m.list.Items()[m.list.Index()]

		// Assuming your list items are of type Attribute
		if attr, ok := selectedItem.(Attribute); ok {
			inputField = fmt.Sprintf("Editing: %s (original: %s)\n%s", attr.Value, attr.Value, m.editText)
		}
	} else {
		inputField = ""
	}

	return docStyle.Render(fmt.Sprintf("%s\n%s", m.list.View(), inputField))
}

func main() {
	items := []list.Item{
		Attribute{Type: "Strength", Value: "60"},
		Attribute{Type: "Dexterity", Value: "70"},
		Attribute{Type: "Constitution", Value: "55"},
		Attribute{Type: "Intelligence", Value: "75"},
		Attribute{Type: "Power", Value: "40"},
		Attribute{Type: "Appearance", Value: "65"},
		Attribute{Type: "Size", Value: "65"},
		Attribute{Type: "Education", Value: "80"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My CoC Character"

	p := tea.NewProgram(m, tea.WithAltScreen())

	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
