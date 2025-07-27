package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx         context.Context
	defaultPath string
}

type AppSettings struct {
	DefaultPath string `json:"defaultPath"`
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// Load settings
	settingsFile, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("Error when opening file: %s", err)
		return
	}

	var settings AppSettings
	err = json.Unmarshal(settingsFile, &settings)
	if err != nil {
		fmt.Printf("Error when unmarshalling data: %s", err)
		return
	}
	a.defaultPath = settings.DefaultPath
}

/*
This function was taken from a wails v2 example, located here:

https://github.com/tataDan/wails-v2-examples/blob/eb48c734627cc56bab846cd5b19e0078449458d4/examples/research/app.go#L801

The examples are released under the MIT license.
*/
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

func (a *App) GetSettings() AppSettings {
	settings := &AppSettings{}
	settings.DefaultPath = a.defaultPath
	return *settings
}

func (a *App) SelectPdf() string {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
	})
	slashFile := filepath.ToSlash(file)

	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "Error:' file failed to open", err.Error())
		return ""
	}

	if slashFile == "" {
		a.msgDlgOk(runtime.ErrorDialog, "Error: failed to choose a file", "no file chosen")
		return ""
	}

	if !strings.Contains(slashFile, a.defaultPath) {
		a.msgDlgOk(runtime.ErrorDialog, "Error: failed to choose file", "file is not in "+a.defaultPath)
	}

	// This returns with quotes around it and I have no idea why or how to fix it on the go side
	return slashFile
}

func (a *App) GetPageCount(filePath string) int {
	return a.getPageCount(filePath)
}

func (a *App) GetPage(filePath string, pageCount int) []byte {
	// runtime.BrowserOpenURL(a.ctx, file)
	pdfFile := a.readPdfPage(filePath, pageCount)
	pdfData, err := io.ReadAll(pdfFile)
	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "Error: failed to read file", err.Error())
		return nil
	}
	return pdfData
}
