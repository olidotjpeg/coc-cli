package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input interface {
	Value() string
	Blur() tea.Msg
	Update(tea.Msg) (Input, tea.Cmd)
	View() string
}

type ShortAnswerField struct {
	textInput textinput.Model
}

func (sa *ShortAnswerField) Value() string {
	return sa.textInput.Value()
}

func (sa *ShortAnswerField) Blur() tea.Msg {
	return sa.textInput.Blur
}

func (sa *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sa.textInput, cmd = sa.textInput.Update(msg)

	return sa, cmd
}

func (sa *ShortAnswerField) View() string {
	return sa.textInput.View()
}

func NewShortAnswerField() *ShortAnswerField {
	ti := textinput.New()
	ti.Placeholder = "Your answer goes here"
	ti.Focus()
	return &ShortAnswerField{ti}
}

// Long Answer Fields
type LongAnswerField struct {
	textArea textarea.Model
}

func (la *LongAnswerField) Value() string {
	return la.textArea.Value()
}

func (la *LongAnswerField) Blur() tea.Msg {
	return la.textArea.Blur
}

func (la *LongAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	la.textArea, cmd = la.textArea.Update(msg)

	return la, cmd
}

func (la *LongAnswerField) View() string {
	return la.textArea.View()
}

// Text Area
func NewLongAnswerField() *LongAnswerField {
	ta := textarea.New()
	ta.Placeholder = "Your answer goes here"
	ta.Focus()
	return &LongAnswerField{ta}
}
