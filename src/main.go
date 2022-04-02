package main

// A simple example that shows how to render an animated progress bar. In this
// example we bump the progress by 25% every two seconds, animating our
// progress bar to its new target state.
//
// It's also possible to render a progress bar in a more static fashion without
// transitions. For details on that approach see the progress-static example.

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
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

	m := model{
		progress: progress.New(progress.WithGradient(GetRandomColorInHex(), GetRandomColorInHex())),
	}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
