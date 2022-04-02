package main

// A simple example that shows how to render an animated progress bar. In this
// example we bump the progress by 25% every two seconds, animating our
// progress bar to its new target state.
//
// It's also possible to render a progress bar in a more static fashion without
// transitions. For details on that approach see the progress-static example.

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

//#region textinput
type model_4 struct {
	textInput textinput.Model
	err       error
}

func initialModel() model_4 {
	ti := textinput.New()
	ti.Placeholder = "Username / Password"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model_4{
		textInput: ti,
		err:       nil,
	}
}

type tickMsg2 struct{}
type errMsg2 error

func (m model_4) Init() tea.Cmd {
	return textinput.Blink
}

func (m model_4) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg2:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model_4) View() string {
	return fmt.Sprintf(
		"Log In \n\n%s\n\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

//#endregion

func main() {
	/*region http
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
	#endregion*/

	//#region check command
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
	//#endregion

	//#region picker
	var (
		helpStyle = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	)

	items := []list.Item{
		item("Jazzmine"),
		item("NodeJS"),
		item("Raylib"),
		item("Plib"),
		item("Pango"),
		item("Unity3D"),
		item("U36"),
		item("U36 F#"),
		item("Castrime"),
	}

	const defaultWidth = 20

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Pick a package"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m2 := model_2{list: l}

	if err := tea.NewProgram(m2).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
	//#endregion
	//#region progress
	m := model{
		progress: progress.New(progress.WithGradient(GetRandomColorInHex(), GetRandomColorInHex())),
	}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
	//#endregion
}
