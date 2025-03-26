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
	gui := gui.GuiApp{
		InputPassword: widget.NewEntry(),
		Window:        &window,
	}
	gui.Lang = "en"
	gui.DataByLang = map[string]map[string]string{
		"ru": map[string]string{
			"passwordPlaceholder":       "Введите пароль, длина должна быть равна 16",
			"btnExit":                   "Выход",
			"generatePasswordBtn":       "Сгенерировать пароль",
			"clearBtn":                  "Очистить все данные окна",
			"selectFileBtn":             "Выберите файл",
			"selectedFilePlaceholder":   "Еще не выбран файл",
			"processFileBtn":            "Обработать файл",
			"selectedMethodPlaceholder": "Выберите метод",
			"encryptMethod":             "Зашифровать",
			"decryptMethod":             "Расшифровать",
			"errPasswordLength":         "Длина пароля должна быть равна 16",
			"missingFile":               "Загрузите файл",
			"missingMethod":             "Выберите метод",
		},
		"en": map[string]string{
			"passwordPlaceholder":       "Enter password, length must be 16 characters",
			"btnExit":                   "Exit",
			"generatePasswordBtn":       "Generate password",
			"clearBtn":                  "Clearing window data",
			"selectFileBtn":             "Select file",
			"selectedFilePlaceholder":   "No file yet",
			"processFileBtn":            "Process the file",
			"selectedMethodPlaceholder": "Select method",
			"encryptMethod":             "Encrypt",
			"decryptMethod":             "Decrypt",
			"errPasswordLength":         "Password length must be 16 characters",
			"missingFile":               "Upload file",
			"missingMethod":             "Select method",
		},
	}
	gui.BtnExit = widget.NewButton(gui.DataByLang[gui.Lang]["btnExit"], func() {
		app.Quit()
	})
	gui.SelectedFile = widget.NewLabel(gui.DataByLang[gui.Lang]["selectedFilePlaceholder"])
	gui.LangToggler = gui.LangTogglerHandler()
	gui.SelectMethod = gui.GetSelectMethod()
	gui.InputPassword.SetPlaceHolder(gui.DataByLang[gui.Lang]["passwordPlaceholder"])
	gui.GeneratePasswordBtn = gui.GeneratePasswordBtnBtnHandler()
	gui.ClearResultBtn = gui.ClearWindowBtnHandler()
	gui.SelectFileBtn = gui.SelectFileBtnHandler()
	gui.ProcessBtn = gui.ProcessBtnHandler()

	content := container.NewGridWithRows(
		3,
		container.NewGridWithColumns(
			2,
			gui.LangToggler,
			container.NewGridWithRows(
				2,
				gui.InputPassword,
				gui.GeneratePasswordBtn,
			),
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
		container.NewGridWithRows(
			2,
			gui.ProcessBtn,
			gui.BtnExit,
		),
	)
	window.SetContent(content)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(600, 300))
	window.ShowAndRun()
}
