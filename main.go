package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

type FileLoader struct {
    http.Handler
}

func NewFileLoader() *FileLoader {
    return &FileLoader{}
}

//Test - try "http://localhost:34115/DishWorker.js.LICENSE.txt"
func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    var err error

    requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	
    println("Requesting file:", requestedFilename)
    fileData, err := os.ReadFile(requestedFilename)
    if err != nil {
        res.WriteHeader(http.StatusBadRequest)
        res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
    }

    res.Write(fileData)
}


func main() {

fileload := NewFileLoader()
// Create an instance of the app structure
app := NewApp()

err := wails.Run(&options.App{
	Title:             	"CyberChef",
	Width:             	1024,
	Height:            	768,
	Assets:            	assets,
	HideWindowOnClose: 	false,
	BackgroundColour:  	&options.RGBA{R: 27, G: 38, B: 54, A: 1},
	Menu:              	nil,
	Logger:            	nil,
	LogLevelProduction: logger.ERROR,
	AssetsHandler:      fileload,
	OnStartup:         	func(ctx context.Context){
							ctx = context.WithValue(ctx, "token", "myValue...1234")
							app.SetContext(ctx)
						},
	OnDomReady:        	app.domReady,
	OnBeforeClose:     	app.beforeClose,
	OnShutdown:        	app.shutdown,
	WindowStartState:  	options.Normal,
	Bind: []interface{}{
		app,
	},
// Windows platform specific options
	Windows: &windows.Options{
		WebviewIsTransparent: false,
		WindowIsTranslucent:  false,
		DisableWindowIcon:    false,
		WebviewUserDataPath: "",
	},
// Mac platform specific options
	Mac: &mac.Options{
		TitleBar: &mac.TitleBar{
			TitlebarAppearsTransparent: true,
			HideTitle:                  false,
			HideTitleBar:               false,
			FullSizeContent:            false,
			UseToolbar:                 false,
			HideToolbarSeparator:       true,
		},
		Appearance:           mac.NSAppearanceNameDarkAqua,
		WebviewIsTransparent: true,
		WindowIsTranslucent:  true,
		About: &mac.AboutInfo{
			Title:   "CyberChef",
			Message: "",
			Icon:    icon,
		},
	},
})

if err != nil {
	log.Fatal(err)
}
}
