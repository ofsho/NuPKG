package main

// A simple program that makes a GET request and prints the response status.

import (
	"fmt"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type model_3 struct {
	status int
	err    error
}

type statusMsg int

type errMsg struct{ error }

func (e errMsg) Error() string { return e.error.Error() }

func (m model_3) Init() tea.Cmd {
	return checkServer
}

func (m model_3) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		default:
			return m, nil
		}

	case statusMsg:
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, nil

	default:
		return m, nil
	}
}

func (m model_3) View() string {
	s := fmt.Sprintf("Checking %s...", url)
	if m.err != nil {
		s += fmt.Sprintf("something went wrong: %s", m.err)
	} else if m.status != 0 {
		s += fmt.Sprintf("%d %s", m.status, http.StatusText(m.status))
	}
	return s + "\n"
}

func checkServer() tea.Msg {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}
	return statusMsg(res.StatusCode)
}