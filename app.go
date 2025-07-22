package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"encoding/csv"
	"encoding/json"
)

// App struct
type App struct {
	ctx context.Context
}

type PDFMapMonster struct {
	Id string `json:"id"`
    Name string `json:"name"`
    CR float32 `json:"cr"`
    MonsterType string `json:"monsterType"`
    Size string `json:"size"`
}

type PDFMapPage struct {
	Id string `json:"id"`
    PageNumber int `json:"pageNumber"`
    Items []PDFMapMonster `json:"items"`
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

// This code was taken from a wails v2 example, located here:
// https://github.com/tataDan/wails-v2-examples/blob/eb48c734627cc56bab846cd5b19e0078449458d4/examples/research/app.go#L801
// The examples are released under the MIT license
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

func (a *App) getPathStrings(path string) (string, string) {
	// Break file extension off of path
	pathSplitExtension := strings.Split(path, ".")
	// Break path into hierarchy to add a folder
	pathSplitHierarchy := strings.Split(pathSplitExtension[0], "\\")
	// Create path for json file from split
	dirPath := ""
	itemPath := ""
	endIndex := len(pathSplitHierarchy) - 1
	for i, pathItem := range pathSplitHierarchy {
		if i == endIndex {
			dirPath += "monster_map"
			itemPath = dirPath + "/" + pathItem + ".json"
		} else {
			dirPath += pathItem +  "/"
		}
	}
	return dirPath, itemPath
}

// Open a dialog to let the user choose a pdf then return the name to the frontend
func (a *App) GetPDF() string {
	pdfPath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions {
		Title: "Select File",
	})

	if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return ""
	}

	if pdfPath == "" {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
		return ""
	}

	// For now, we just open the pdf in a browser
	// so we don't have to worry about pulling in large pdfs
	// or exposing the file system to the frontend
	runtime.BrowserOpenURL(a.ctx, pdfPath)
	return pdfPath;
}

func (a *App) OpenPDFMap(pdfPath string) []PDFMapPage {
	pdfMap := []PDFMapPage{{}}
	dirPath, itemPath := a.getPathStrings(pdfPath)

	// Create monster directory
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		fmt.Println("Error creating the directory:", err)
	}
	
	// Create a new file or overwrite an existing file
    file, err := os.Create(itemPath)
    if err != nil {
		fmt.Println("Error creating the file:", err)
        return pdfMap
    }
    defer file.Close() // Ensure the file gets closed after we are done

	// Create a JSON decoder (and set indentation)
    decoder := json.NewDecoder(file)
	decoder.Decode(pdfMap)

	return pdfMap;
}

func (a *App) SavePDFMap(pdfMap []PDFMapPage, pdfPath string) {
	dirPath, itemPath := a.getPathStrings(pdfPath)

	a.msgDlgOk(runtime.ErrorDialog, itemPath, itemPath)

	// Create monster directory
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		fmt.Println("Error creating the directory:", err)
	}
	
	// Create a new file or overwrite an existing file
    file, err := os.Create(itemPath)
    if err != nil {
		fmt.Println("Error creating the file:", err)
        return
    }
    defer file.Close() // Ensure the file gets closed after we are done
	
    // Create a JSON encoder (and set indentation for readability)
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
	
    // Encode the pdfMap and write it to the file
    if err := encoder.Encode(pdfMap); err != nil {
		fmt.Println("Error encoding JSON:", err)
    }
}

func (a *App) SaveCSV() {
	path := "C:/Users/emwit/Desktop/TTRPG/MonsterMaps/monsters.csv"
	data := [][]string{
        {"blah", "blah"},
        {"test", "test"},
        {"potato", "strawberry"},
    }

	// create a file
    file, err := os.Create(path)
    if err != nil {
		a.msgDlgOk(runtime.ErrorDialog, "ERROR", err.Error())
    }
    defer file.Close()

	// initialize csv writer
    writer := csv.NewWriter(file)

    defer writer.Flush()

    // write all rows at once
    writer.WriteAll(data)
}
