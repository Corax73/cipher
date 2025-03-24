package gui

import (
	"cipher/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	Password, Method, LoadedFilePath, Lang string
	IsFileLoaded                           bool
	SelectedMethodNumber                   int
}

func (state *State) ResetState() {
	state.Password, state.Method, state.LoadedFilePath, state.Lang, state.SelectedMethodNumber = "", "", "", "en", 0
}

type Internationalization struct {
	DataByLang map[string]map[string]string
}

type GuiApp struct {
	State
	Internationalization
	SelectedFile                                                            *widget.Label
	InputPassword                                                           *widget.Entry
	GeneratePasswordBtn, ProcessBtn, ClearResultBtn, SelectFileBtn, BtnExit *widget.Button
	Window                                                                  *fyne.Window
	SelectMethod                                                            *widget.Select
	FileURI                                                                 fyne.URI
	LangToggler                                                             *widget.RadioGroup
}

func (gui *GuiApp) GetSelectMethod() *widget.Select {
	methods := make([]string, 2)
	methods[0] = gui.DataByLang[gui.Lang]["encryptMethod"]
	methods[1] = gui.DataByLang[gui.Lang]["decryptMethod"]
	resp := widget.NewSelect(
		methods,
		func(value string) {
			gui.Method = value
			if value == gui.DataByLang[gui.Lang]["encryptMethod"] {
				gui.SelectedMethodNumber = 1
			} else {
				gui.SelectedMethodNumber = 2
			}
		},
	)
	resp.PlaceHolder = gui.DataByLang[gui.Lang]["selectedMethodPlaceholder"]
	return resp
}

func (gui *GuiApp) GeneratePasswordBtnBtnHandler() *widget.Button {
	return widget.NewButton(
		gui.DataByLang[gui.Lang]["generatePasswordBtn"],
		func() {
			gui.InputPassword.SetText(core.PasswordGenerator(16))
		},
	)
}

func (gui *GuiApp) ClearWindowBtnHandler() *widget.Button {
	return widget.NewButton(
		gui.DataByLang[gui.Lang]["clearBtn"],
		func() {
			gui.InputPassword.SetText("")
			gui.SelectMethod.ClearSelected()
			gui.SelectMethod.Refresh()
			gui.FileURI = nil
			gui.SelectedFile.SetText(gui.DataByLang[gui.Lang]["selectedFilePlaceholder"])
			gui.SelectedFile.Refresh()
			gui.ResetState()
		},
	)
}

func (gui *GuiApp) SelectFileBtnHandler() *widget.Button {
	return widget.NewButton(
		gui.DataByLang[gui.Lang]["selectFileBtn"],
		func() {
			dialog.ShowFileOpen(
				func(reader fyne.URIReadCloser, err error) {
					saveFile := gui.DataByLang[gui.Lang]["selectedFilePlaceholder"]
					if err != nil {
						gui.IsFileLoaded = false
						dialog.ShowError(err, *gui.Window)
						return
					}
					if reader == nil {
						gui.IsFileLoaded = false
						return
					}
					saveFile = reader.URI().Path()
					gui.FileURI = reader.URI()
					gui.SelectedFile.SetText(saveFile)
					gui.IsFileLoaded = true
				},
				*gui.Window,
			)
		},
	)
}

func (gui *GuiApp) ProcessBtnHandler() *widget.Button {
	return widget.NewButton(
		gui.DataByLang[gui.Lang]["processFileBtn"],
		func() {
			gui.Password = gui.InputPassword.Text
			gui.LoadedFilePath = gui.SelectedFile.Text
			if gui.Password != "" &&
				len(gui.Password) == 16 &&
				gui.FileURI != nil &&
				gui.LoadedFilePath != "" &&
				gui.Method != "" {
				key := []byte(gui.Password)
				var data []byte
				var err error
				encryptVal := gui.DataByLang[gui.Lang]["encryptMethod"]
				decryptVal := gui.DataByLang[gui.Lang]["decryptMethod"]
				if gui.Method == encryptVal {
					dataStr, err := core.EncryptFile(key, gui.LoadedFilePath)
					data = []byte(dataStr)
					if err != nil {
						dialog.ShowError(err, *gui.Window)
						return
					}
				} else if gui.Method == decryptVal {
					data, err = core.DecryptFile(key, gui.LoadedFilePath)
					if err != nil {
						dialog.ShowError(err, *gui.Window)
						return
					}
				}
				if len(data) > 0 {
					dialog.ShowFileSave(
						func(writer fyne.URIWriteCloser, err error) {
							if err == nil && writer != nil {
								_, err := writer.Write(data)
								if err != nil {
									dialog.ShowError(err, *gui.Window)
								}
							}
						},
						*gui.Window,
					)
				}
			}
		},
	)
}

func (gui *GuiApp) LangTogglerHandler() *widget.RadioGroup {
	resp := widget.NewRadioGroup(
		[]string{"en", "ru"},
		func(value string) {
			gui.Lang = value
			gui.refreshAllCanvas()
		},
	)
	resp.Selected = gui.Lang
	resp.Required = true
	return resp
}

func (gui *GuiApp) refreshAllCanvas() {
	gui.InputPassword.SetPlaceHolder(gui.DataByLang[gui.Lang]["passwordPlaceholder"])
	gui.InputPassword.Refresh()
	gui.GeneratePasswordBtn.Text = gui.DataByLang[gui.Lang]["generatePasswordBtn"]
	gui.GeneratePasswordBtn.Refresh()
	gui.BtnExit.Text = gui.DataByLang[gui.Lang]["btnExit"]
	gui.BtnExit.Refresh()
	gui.ClearResultBtn.Text = gui.DataByLang[gui.Lang]["clearBtn"]
	gui.ClearResultBtn.Refresh()
	gui.SelectFileBtn.Text = gui.DataByLang[gui.Lang]["selectFileBtn"]
	gui.SelectFileBtn.Refresh()
	if !gui.IsFileLoaded {
		gui.SelectedFile.SetText(gui.DataByLang[gui.Lang]["selectedFilePlaceholder"])
		gui.InputPassword.Refresh()
	}
	gui.SelectMethod.SetOptions(
		[]string{
			gui.DataByLang[gui.Lang]["encryptMethod"],
			gui.DataByLang[gui.Lang]["decryptMethod"],
		})
	gui.SelectMethod.PlaceHolder = gui.DataByLang[gui.Lang]["selectedMethodPlaceholder"]
	if gui.SelectedMethodNumber == 1 {
		gui.SelectMethod.SetSelected(gui.DataByLang[gui.Lang]["encryptMethod"])
	} else if gui.SelectedMethodNumber == 2 {
		gui.SelectMethod.SetSelected(gui.DataByLang[gui.Lang]["decryptMethod"])
	} else {
		gui.SelectMethod.ClearSelected()
	}
	gui.SelectMethod.Refresh()
	gui.ProcessBtn.Text = gui.DataByLang[gui.Lang]["processFileBtn"]
	gui.ProcessBtn.Refresh()
}
