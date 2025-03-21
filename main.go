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
		SelectedFile:   widget.NewLabel("Selected file"),
	}
	gui.SelectMethod = gui.GetSelectMethod()
	gui.InputPassword.SetPlaceHolder("Enter password, minimum length 16")
	gui.InputOutputExt.SetPlaceHolder("Enter the output file extension")
	gui.GeneratePasswordBtn = gui.GeneratePasswordBtnBtnHandler()
	gui.ClearResultBtn = gui.ClearWindowBtnHandler()
	gui.SelectFileBtn = gui.SelectFileBtnHandler()

	content := container.NewGridWithRows(
		3,
		container.NewGridWithColumns(
			2,
			container.NewGridWithRows(
				2,
				gui.InputPassword,
				gui.InputOutputExt,
			),
			gui.GeneratePasswordBtn,
		),
		container.NewGridWithColumns(
			3,
			gui.SelectMethod,
			container.NewGridWithRows(
				2,
				gui.SelectedFile,
				gui.SelectFileBtn,
			),
			gui.ClearResultBtn,
		),
		btnExit,
	)
	window.SetContent(content)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(600, 300))
	window.ShowAndRun()
}
