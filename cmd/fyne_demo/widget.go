package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeButtonTab() fyne.Widget {
	return widget.NewVBox(
		widget.NewLabel("Text label"),
		widget.NewButton("Text button", func() { fmt.Println("tapped text button") }),
		widget.NewButtonWithIcon("With icon", theme.ConfirmIcon(), func() { fmt.Println("tapped icon button") }),
	)
}

func makeInputTab() fyne.Widget {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Entry")

	return widget.NewVBox(
		entry,
		widget.NewCheck("Check", func(on bool) { fmt.Println("checked", on) }),
		widget.NewRadio([]string{"Item 1", "Item 2"}, func(s string) { fmt.Println("selected", s) }),
	)
}

func makeGroupTab() fyne.Widget {
	progress := widget.NewProgressBar()
	infProgress := widget.NewProgressBarInfinite()

	go func() {
		num := 0.0
		for num < 1.0 {
			time.Sleep(100 * time.Millisecond)
			progress.SetValue(num)
			num += 0.01
		}

		progress.SetValue(1)
	}()

	return widget.NewGroup("Grouped",
		widget.NewLabel("Grouped content"),
		progress,
		infProgress)
}

func makeScrollTab() fyne.Widget {
	logo := canvas.NewImageFromResource(theme.FyneLogo())
	logo.SetMinSize(fyne.NewSize(320, 320))
	scroll := widget.NewScroller(logo)
	scroll.Resize(fyne.NewSize(200, 200))

	return scroll
}

// Widget shows a window containing widget demos
func Widget(app fyne.App) {
	w := app.NewWindow("Widgets")

	w.SetContent(widget.NewTabContainer(
		widget.NewTabItem("Buttons", makeButtonTab()),
		widget.NewTabItem("Input", makeInputTab()),
		widget.NewTabItem("Group", makeGroupTab()),
		widget.NewTabItem("Scroll", makeScrollTab()),
	))

	w.Show()
}
