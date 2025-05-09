package filetree

import (
	"path/filepath"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/mistakenelf/fm/filesystem"
	"github.com/mistakenelf/fm/polish"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	if m.Disabled {
		return m, nil
	}

	switch msg := msg.(type) {
	case editorFinishedMsg:
		if msg.err != nil {
			m.err = msg.err
			return m, tea.Quit
		}
	case errorMsg:
		cmds = append(cmds, m.NewStatusMessageCmd(
			lipgloss.NewStyle().
				Foreground(polish.Colors.Red600).
				Bold(true).
				Render(string(msg))))
	case statusMessageTimeoutMsg:
		m.StatusMessage = ""
	case moveDirectoryItemMsg:
		m.State = IdleState

		return m, m.GetDirectoryListingCmd(m.CurrentDirectory)
	case copyToClipboardMsg:
		cmds = append(cmds, m.NewStatusMessageCmd(
			lipgloss.NewStyle().
				Bold(true).
				Render(string(msg))))
	case createFileMsg:
		m.State = IdleState

		return m, m.GetDirectoryListingCmd(m.CurrentDirectory)
	case createDirectoryMsg:
		m.State = IdleState

		return m, m.GetDirectoryListingCmd(m.CurrentDirectory)
	case renameDirectoryItemMsg:
		m.State = IdleState

		return m, m.GetDirectoryListingCmd(m.CurrentDirectory)
	case getDirectoryListingMsg:
		if msg.files != nil {
			m.files = msg.files
		} else {
			m.files = make([]DirectoryItem, 0)
		}

		m.CurrentDirectory = msg.workingDirectory
		m.Cursor = 0
		m.min = 0
		m.max = max(m.max, m.height-1)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.Down):
			if m.State != IdleState {
				return m, nil
			}

			m.Cursor++

			if m.Cursor >= len(m.files) {
				m.Cursor = len(m.files) - 1
			}

			if m.Cursor > m.max {
				m.min++
				m.max++
			}
		case key.Matches(msg, m.keyMap.Up):
			if m.State != IdleState {
				return m, nil
			}

			m.Cursor--

			if m.Cursor < 0 {
				m.Cursor = 0
			}

			if m.Cursor < m.min {
				m.min--
				m.max--
			}
		case key.Matches(msg, m.keyMap.GotoTop):
			if m.State != IdleState {
				return m, nil
			}

			m.Cursor = 0
			m.min = 0
			m.max = m.height
		case key.Matches(msg, m.keyMap.GotoBottom):
			if m.State != IdleState {
				return m, nil
			}

			m.Cursor = len(m.files) - 1
			m.min = len(m.files) - m.height
			m.max = len(m.files) - 1
		case key.Matches(msg, m.keyMap.PageDown):
			if m.State != IdleState {
				return m, nil
			}

			m.Cursor += m.height

			if m.Cursor >= len(m.files) {
				m.Cursor = len(m.files) - 1
			}

			m.min += m.height
			m.max += m.height

			if m.max >= len(m.files) {
				m.max = len(m.files) - 1
				m.min = m.max - m.height
			}
		case key.Matches(msg, m.keyMap.PageUp):
			if m.State != IdleState {
				return m, nil
			}

			m.Cursor -= m.height

			if m.Cursor < 0 {
				m.Cursor = 0
			}

			m.min -= m.height
			m.max -= m.height

			if m.min < 0 {
				m.min = 0
				m.max = m.min + m.height
			}
		case key.Matches(msg, m.keyMap.GoToHomeDirectory):
			if m.State != IdleState {
				return m, nil
			}

			return m, m.GetDirectoryListingCmd(filesystem.HomeDirectory)
		case key.Matches(msg, m.keyMap.GoToRootDirectory):
			if m.State != IdleState {
				return m, nil
			}

			return m, m.GetDirectoryListingCmd(filesystem.RootDirectory)
		case key.Matches(msg, m.keyMap.ToggleHidden):
			if m.State != IdleState {
				return m, nil
			}

			m.showHidden = !m.showHidden

			return m, m.GetDirectoryListingCmd(m.CurrentDirectory)
		case key.Matches(msg, m.keyMap.OpenDirectory):
			if m.State != IdleState {
				return m, nil
			}

			if len(m.files) == 0 {
				return m, nil
			}

			if m.files[m.Cursor].IsDirectory {
				return m, m.GetDirectoryListingCmd(m.files[m.Cursor].Path)
			}
		case key.Matches(msg, m.keyMap.PreviousDirectory):
			if m.State != IdleState {
				return m, nil
			}

			return m, m.GetDirectoryListingCmd(
				filepath.Dir(m.CurrentDirectory),
			)
		case key.Matches(msg, m.keyMap.CopyPathToClipboard):
			if m.State != IdleState {
				return m, nil
			}

			return m, copyToClipboardCmd(m.files[m.Cursor].Path)
		case key.Matches(msg, m.keyMap.CopyDirectoryItem):
			if m.State != IdleState {
				return m, nil
			}

			return m, tea.Sequence(
				copyDirectoryItemCmd(m.files[m.Cursor].Path, m.files[m.Cursor].IsDirectory),
				m.GetDirectoryListingCmd(m.CurrentDirectory),
			)
		case key.Matches(msg, m.keyMap.DeleteDirectoryItem):
			if m.State != IdleState {
				return m, nil
			}

			return m, tea.Sequence(
				deleteDirectoryItemCmd(m.files[m.Cursor].Path, m.files[m.Cursor].IsDirectory),
				m.GetDirectoryListingCmd(m.CurrentDirectory),
			)
		case key.Matches(msg, m.keyMap.ZipDirectoryItem):
			if m.State != IdleState {
				return m, nil
			}

			return m, tea.Sequence(
				zipDirectoryCmd(m.files[m.Cursor].Path),
				m.GetDirectoryListingCmd(m.CurrentDirectory),
			)
		case key.Matches(msg, m.keyMap.UnzipDirectoryItem):
			if m.State != IdleState {
				return m, nil
			}

			return m, tea.Sequence(
				unzipDirectoryCmd(m.files[m.Cursor].Name),
				m.GetDirectoryListingCmd(m.CurrentDirectory),
			)
		case key.Matches(msg, m.keyMap.ShowDirectoriesOnly):
			if m.State != IdleState {
				return m, nil
			}

			m.showDirectoriesOnly = !m.showDirectoriesOnly
			m.showFilesOnly = false

			return m, m.GetDirectoryListingCmd(m.CurrentDirectory)
		case key.Matches(msg, m.keyMap.ShowFilesOnly):
			if m.State != IdleState {
				return m, nil
			}

			m.showFilesOnly = !m.showFilesOnly
			m.showDirectoriesOnly = false

			return m, m.GetDirectoryListingCmd(m.CurrentDirectory)
		case key.Matches(msg, m.keyMap.WriteSelectionPath):
			if m.State != IdleState {
				return m, nil
			}

			if m.selectionPath != "" {
				return m, tea.Sequence(
					writeSelectionPathCmd(m.selectionPath, m.files[m.Cursor].Name),
					tea.Quit,
				)
			}
		case key.Matches(msg, m.keyMap.OpenInEditor):
			if m.State != IdleState {
				return m, nil
			}

			return m, openEditorCmd(m.files[m.Cursor].Name)
		case key.Matches(msg, m.keyMap.CreateFile):
			if m.State != IdleState {
				return m, nil
			}

			m.State = CreateFileState

			return m, nil
		case key.Matches(msg, m.keyMap.CreateDirectory):
			if m.State != IdleState {
				return m, nil
			}

			m.State = CreateDirectoryState

			return m, nil
		case key.Matches(msg, m.keyMap.RenameDirectoryItem):
			if m.State != IdleState {
				return m, nil
			}

			m.State = RenameState
		}
	}

	return m, tea.Batch(cmds...)
}
