package main

import (
	"cipher/gui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("Text file encryptor/decryptor")

	btnExit := widget.NewButton("Exit", func() {
		app.Quit()
	})

	gui := gui.GuiApp{
		InputPassword:  widget.NewEntry(),
		InputOutputExt: widget.NewEntry(),
		Window:         &window,
	}
	gui.SelectMethod = gui.GetSelectMethod()
	gui.InputPassword.SetPlaceHolder("Enter password")
	gui.InputOutputExt.SetPlaceHolder("Enter the output file extension")

	content := container.NewGridWithRows(
		3,
		container.NewGridWithColumns(
			2,
			gui.InputPassword,
			gui.InputOutputExt,
		),
		gui.SelectMethod,
		btnExit,
	)
	window.SetContent(content)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(500, 400))
	window.ShowAndRun()
}
