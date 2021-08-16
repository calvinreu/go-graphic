package graphic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/veandco/go-sdl2/sdl"
)

//Config data loaded form config json
type Config struct {
	Window      WindowConfig
	BaseSprites []SpriteConfig
}

//WindowConfig Render info for Window
type WindowConfig struct {
	Title                           string
	X, Y, Width, Height             int32
	WindowFlags, RendererFlags, FPS uint32
}

//ImageConfig contains info about one img
type SpriteConfig struct {
	ImgPath  string
	SrcRects []sdl.Rect
}

//Init Config from json files
func (config *Config) Init(windowJSON, spriteJSON string) {
	fmt.Println("Config files:")
	fmt.Println("Window config file: ", windowJSON)
	fmt.Println("Sprite config file: ", spriteJSON)
	fmt.Println("----------")

	//load Window Config file
	dataWindow, err := ioutil.ReadFile(windowJSON)
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(dataWindow, &config.Window)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Window config:")
	fmt.Println("Title: ", config.Window.Title)
	fmt.Println("X: ", config.Window.X)
	fmt.Println("Y: ", config.Window.Y)
	fmt.Println("Width: ", config.Window.Width)
	fmt.Println("Height: ", config.Window.Height)
	fmt.Println("WindowFlags: ", config.Window.WindowFlags)
	fmt.Println("RendererFlags: ", config.Window.RendererFlags)
	fmt.Println("FPS: ", config.Window.FPS)
	fmt.Println("----------")

	//load Image Config File
	dataImage, err := ioutil.ReadFile(spriteJSON)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(dataImage, &config.BaseSprites)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Sprite config:")
	for _, i := range config.BaseSprites {
		fmt.Println("Image file: ", i.ImgPath)
		for _, iRect := range i.SrcRects {
			fmt.Println("Source Rectangle:")
			fmt.Println("X: ", iRect.X)
			fmt.Println("Y: ", iRect.Y)
			fmt.Println("W: ", iRect.W)
			fmt.Println("H: ", iRect.H)
			fmt.Println()
		}
		fmt.Println("----------")
	}
	fmt.Println("----------")
}
