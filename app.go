package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) msgDlgOk(dlgType runtime.DialogType, title string, msg string) {
	_, err := runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
		Type:    dlgType,
		Title:   title,
		Message: msg,
	})
	if err != nil {
		fmt.Print(err)
	}
}

func (a *App) GetPDF() string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions {
		Title: "Select File",
	})

	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return ""
	}

	if file == "" {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return ""
	}

	runtime.BrowserOpenURL(a.ctx, file)
	return file;
}
