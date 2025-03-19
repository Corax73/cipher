package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	Password, OutputExt, Method string
}

func (state *State) ResetState() {
	state.Password, state.OutputExt = "", ""
}

type GuiApp struct {
	State
	InputPassword, InputOutputExt   *widget.Entry
	GeneratePasswordBtn, ProcessBtn *widget.Button
	Window                          *fyne.Window
	SelectMethod                    *widget.Select
}

func (gui *GuiApp) GetSelectMethod() *widget.Select {
	resp := widget.NewSelect([]string{"Encrypt", "Decrypt"}, func(value string) {
		gui.Method = value
	})
	resp.PlaceHolder = "Select method"
	return resp
}
