package main

import (
	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite))
	
	screen.Clear()

	screen.SetContent(2, 2, 'H', nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))

	screen.Show()

	for {
		event := screen.PollEvent()

		if key, ok := event.(*tcell.EventKey); ok {
			switch key.Key() {
			case tcell.KeyCtrlC:
				return
			case tcell.KeyEscape:
			}
		}
	}
}