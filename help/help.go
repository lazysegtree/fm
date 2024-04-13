// Package help implements a help bubble which can be used
// to display help information such as keymaps.
package help

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mistakenelf/fm/polish"
)

const (
	keyWidth = 12
)

type TitleColor struct {
	Background lipgloss.AdaptiveColor
	Foreground lipgloss.AdaptiveColor
}

// Entry represents a single entry in the help bubble.
type Entry struct {
	Key         string
	Description string
}

// Model represents the properties of a help bubble.
type Model struct {
	Viewport         viewport.Model
	Entries          []Entry
	Title            string
	TitleColor       TitleColor
	ViewportDisabled bool
}

func generateHelpScreen(
	title string,
	titleColor TitleColor,
	entries []Entry,
	width, height int,
) string {
	helpScreen := ""

	for _, content := range entries {
		keyText := lipgloss.NewStyle().
			Bold(true).
			Foreground(polish.AdaptiveColors.DefaultText).
			Width(keyWidth).
			Render(content.Key)

		descriptionText := lipgloss.NewStyle().
			Foreground(polish.AdaptiveColors.DefaultText).
			Render(content.Description)

		row := lipgloss.JoinHorizontal(lipgloss.Top, keyText, descriptionText)
		helpScreen += fmt.Sprintf("%s\n", row)
	}

	titleText := lipgloss.NewStyle().Bold(true).
		Background(titleColor.Background).
		Foreground(titleColor.Foreground).
		Padding(0, 1).
		Italic(true).
		Render(title)

	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Render(lipgloss.JoinVertical(
			lipgloss.Top,
			titleText,
			helpScreen,
		))
}

// New creates a new instance of a help bubble.
func New(
	title string,
	titleColor TitleColor,
	entries []Entry,
) Model {
	viewPort := viewport.New(0, 0)
	viewPort.SetContent(generateHelpScreen(title, titleColor, entries, 0, 0))

	return Model{
		Viewport:         viewPort,
		Entries:          entries,
		Title:            title,
		ViewportDisabled: false,
		TitleColor:       titleColor,
	}
}

// SetSize sets the size of the help bubble.
func (m *Model) SetSize(w, h int) {
	m.Viewport.Width = w
	m.Viewport.Height = h

	m.Viewport.SetContent(
		generateHelpScreen(
			m.Title,
			m.TitleColor,
			m.Entries,
			m.Viewport.Width,
			m.Viewport.Height,
		),
	)
}

// SetViewportDisabled toggles the state of the viewport.
func (m *Model) SetViewportDisabled(disabled bool) {
	m.ViewportDisabled = disabled
}

// GotoTop jumps to the top of the viewport.
func (m *Model) GotoTop() {
	m.Viewport.GotoTop()
}

// GotoBottom jumps to the bottom of the viewport.
func (m *Model) GotoBottom() {
	m.Viewport.GotoBottom()
}

// SetTitleColor sets the color of the title.
func (m *Model) SetTitleColor(color TitleColor) {
	m.TitleColor = color

	m.Viewport.SetContent(
		generateHelpScreen(
			m.Title,
			m.TitleColor,
			m.Entries,
			m.Viewport.Width,
			m.Viewport.Height,
		),
	)
}

// Update handles UI interactions with the help bubble.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	if !m.ViewportDisabled {
		m.Viewport, cmd = m.Viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// View returns a string representation of the help bubble.
func (m Model) View() string {
	return m.Viewport.View()
}
