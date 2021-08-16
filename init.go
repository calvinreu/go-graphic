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
	BaseSprites []SpriteBaseConfig
}

//WindowConfig Render info for Window
type WindowConfig struct {
	Title                           string
	X, Y, Width, Height             int32
	WindowFlags, RendererFlags, FPS uint32
}

//SpriteBaseConfig base config to create sprites on
type SpriteBaseConfig struct {
	ImgPath string
	Sprites []SpriteConfig
}

//SpriteConfig config for a sprite based on a SpriteBaseConfig obj
type SpriteConfig struct {
	Name    string
	SrcRect sdl.Rect
}

func (graphic *Graphic) Init(config Config) {
	var err error
	*graphic, err = New(config.Window.Title, config.Window.X, config.Window.Y, config.Window.Width, config.Window.Height, config.Window.WindowFlags, config.Window.RendererFlags, config.Window.FPS)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Window initialized")

	for _, i := range config.BaseSprites {
		mainSpriteID := graphic.AddSprite(i.ImgPath, i.SrcRects[0])
		fmt.Println("Initialized sprite via file path: ", i.ImgPath)
		for _, iSrcRect := range i.SrcRects[1:] {
			graphic.AddSpriteByID(mainSpriteID, iSrcRect)
			fmt.Println("Initialized sprite via spriteID: ", mainSpriteID)
		}
	}
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
		fmt.Println()
		for _, j := range i.Sprites {
			fmt.Println("Source Rectangle of ", j.Name)
			fmt.Println("X: ", j.SrcRect.X)
			fmt.Println("Y: ", j.SrcRect.Y)
			fmt.Println("W: ", j.SrcRect.W)
			fmt.Println("H: ", j.SrcRect.H)
			fmt.Println()
		}
		fmt.Println("----------")
	}
	fmt.Println("----------")
}
