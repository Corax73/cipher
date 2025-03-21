package gui

import (
	"cipher/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	Password, OutputExt, Method, LoadedFilePath string
}

func (state *State) ResetState() {
	state.Password, state.OutputExt, state.Method, state.LoadedFilePath = "", "", "", ""
}

type GuiApp struct {
	State
	SelectedFile                                                   *widget.Label
	InputPassword, InputOutputExt                                  *widget.Entry
	GeneratePasswordBtn, ProcessBtn, ClearResultBtn, SelectFileBtn *widget.Button
	Window                                                         *fyne.Window
	SelectMethod                                                   *widget.Select
	FileURI                                                        fyne.URI
}

func (gui *GuiApp) GetSelectMethod() *widget.Select {
	resp := widget.NewSelect([]string{"Encrypt", "Decrypt"}, func(value string) {
		gui.Method = value
	})
	resp.PlaceHolder = "Select method"
	return resp
}

func (gui *GuiApp) GeneratePasswordBtnBtnHandler() *widget.Button {
	return widget.NewButton(
		"Generate password",
		func() {
			gui.InputPassword.SetText(core.PasswordGenerator(16))
		},
	)
}

func (gui *GuiApp) ClearWindowBtnHandler() *widget.Button {
	return widget.NewButton(
		"Clearing window data",
		func() {
			gui.InputPassword.SetText("")
			gui.InputOutputExt.SetText("")
			gui.SelectMethod.Selected = "Select method"
			gui.SelectMethod.Refresh()
			gui.FileURI = nil
			gui.SelectedFile.SetText("No file yet")
			gui.SelectedFile.Refresh()
			gui.ResetState()
		},
	)
}

func (gui *GuiApp) SelectFileBtnHandler() *widget.Button {
	return widget.NewButton(
		"Select file",
		func() {
			dialog.ShowFileOpen(
				func(reader fyne.URIReadCloser, err error) {
					saveFile := "No file yet"
					if err != nil {
						dialog.ShowError(err, *gui.Window)
						return
					}
					if reader == nil {
						return
					}
					saveFile = reader.URI().Path()
					gui.FileURI = reader.URI()
					gui.SelectedFile.SetText(saveFile)
				},
				*gui.Window,
			)
		},
	)
}
