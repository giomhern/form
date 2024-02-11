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

func (sa *ShortAnswerField) Value() string {
	return sa.textinput.Value()
}

func (sa *ShortAnswerField) Blur() tea.Msg {
	return sa.textinput.Blur
}

func (sa *ShortAnswerField) View() string {
	return sa.textinput.View()
}

type ShortAnswerField struct {
	textinput textinput.Model
}

/*==================LONG ANSWER FIELD================*/
func (la *LongAnswerField) Value() string {
	return la.textarea.Value()
}

func (la *LongAnswerField) Blur() tea.Msg {
	return la.textarea.Blur
}

type LongAnswerField struct {
	textarea textarea.Model
}

func (sa *ShortAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	sa.textinput, cmd = sa.textinput.Update(msg)
	return sa, cmd
}
func (la *LongAnswerField) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	la.textarea, cmd = la.textarea.Update(msg)
	return la, cmd
}

func (la *LongAnswerField) View() string {
	return la.textarea.View()
}

func NewShortAnswerField() *ShortAnswerField {
	field := textinput.New()
	field.Placeholder = "Type your answer here"
	field.Focus()
	return &ShortAnswerField{textinput: field}

}

func NewLongAnswerField() *LongAnswerField {
	field := textarea.New()
	field.Placeholder = "Type your answer here"
	field.Focus()
	return &LongAnswerField{textarea: field}
}
