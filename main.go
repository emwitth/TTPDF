package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
	defaultPath string
}

func NewFileLoader() *FileLoader {
	loader := &FileLoader{}

	// Load default path setting
	settingsFile, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("Error when opening config file: %s", err)
	}

	var settings AppSettings
	err = json.Unmarshal(settingsFile, &settings)
	if err != nil {
		fmt.Printf("Error when unmarshalling data: %s", err)
	}
	loader.defaultPath = settings.DefaultPath
	return loader
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	// Request file relative to the default path for security
	requestedFilename := h.defaultPath + req.URL.Path
	println("Requesting file:", requestedFilename)
	fileData, err := os.ReadFile(requestedFilename)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
	}

	res.Write(fileData)
}

func main() {
	// Create an instance of the app structure
	app := NewApp()
	fileHandler := NewFileLoader()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "TTPDF",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: fileHandler,
		},
		BackgroundColour: &options.RGBA{R: 50, G: 49, B: 58, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
