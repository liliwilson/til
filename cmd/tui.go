package cmd

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO major refactor for this lol. i don't need a list of inputs, just one for title will suffice.

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("submit"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
	area       textarea.Model
	cursorMode cursor.Mode
	width      int
	height     int
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "your markdown here..."
	ta.Cursor.Style = cursorStyle
	ta.ShowLineNumbers = true
    ta.CharLimit = 0
	ta.SetWidth(100)
	ta.SetHeight(20)

	m := model{
		inputs: make([]textinput.Model, 1),
		area:   ta,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle

		switch i {
		case 0:
			t.Placeholder = "your title here..."
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink, textinput.Blink)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// set focus to next input
		case "tab", "shift+tab", "enter":
			s := msg.String()

			// exit if submit pressed
			if s == "enter" && m.focusIndex == len(m.inputs)+1 {
				Save(m.inputs[0].Value(), m.area.Value())
				return m, tea.Quit
			}

			// don't cycle on enter when we are in the text editor
			var cmd tea.Cmd
			if s == "enter" && m.focusIndex == len(m.inputs) {
				m.area, cmd = m.area.Update(msg)
				return m, cmd
			}

			// cycle indexes
			if s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex == len(m.inputs)+1 {
			} else if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs) + 1
			}

			var i int
			cmds := make([]tea.Cmd, len(m.inputs)+1)
			for i = 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			if m.focusIndex == len(m.inputs) {
				cmds[i] = m.area.Focus()
			} else {
				m.area.Blur()
			}

			return m, tea.Batch(cmds...)
		}
	}

	// handle character input
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs)+1)

	var i int
	for i = range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	m.area, cmds[i] = m.area.Update(msg)

	return tea.Batch(cmds...)
}

func (m model) View() string {
	button := &blurredButton
	if m.focusIndex == len(m.inputs)+1 {
		button = &focusedButton
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Center,
			"what did you learn today?\n\n",
			m.inputs[0].View(),
			"\n\ntell me more!\n",
			m.area.View(),
			fmt.Sprintf("\n\n%s\n\n", *button),
		))
}
