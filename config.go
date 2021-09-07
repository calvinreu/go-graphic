package graphic

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

//InitLogger sets the logger output to a file in /tmp
func InitLogger(logger *log.Logger, name string) error {
	file, err := os.Create("/tmp/" + name + ".log")
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(file)

	logger.SetOutput(writer)

	return nil
}

//Load Config from json files
func (config *Config) Load(windowJSON, spriteJSON string) error {
	var windowLog, spriteLog log.Logger
	err := InitLogger(&windowLog, "WindowConfig")
	if err != nil {
		fmt.Println(err)
		fmt.Println("log for window config is disabled")
	}
	err = InitLogger(&spriteLog, "SpriteConfig")
	if err != nil {
		fmt.Println(err)
		fmt.Println("log for sprite config is disabled")
	}
	windowLog.Println("Window config file: ", windowJSON)
	spriteLog.Println("Sprite config file: ", spriteJSON)

	//load Window Config file
	dataWindow, err := ioutil.ReadFile(windowJSON)
	if err != nil {
		windowLog.Println(err)
		return err
	}
	err = json.Unmarshal(dataWindow, &config.Window)
	if err != nil {
		windowLog.Println(err)
		return err
	}

	windowLog.Println("Window config:")
	windowLog.Println("Title: ", config.Window.Title)
	windowLog.Println("X: ", config.Window.X)
	windowLog.Println("Y: ", config.Window.Y)
	windowLog.Println("Width: ", config.Window.Width)
	windowLog.Println("Height: ", config.Window.Height)
	windowLog.Println("WindowFlags: ", config.Window.WindowFlags)
	windowLog.Println("RendererFlags: ", config.Window.RendererFlags)
	windowLog.Println("FPS: ", config.Window.FPS)

	//load Image Config File
	dataImage, err := ioutil.ReadFile(spriteJSON)
	if err != nil {
		spriteLog.Println(err)
		return err
	}

	err = json.Unmarshal(dataImage, &config.BaseSprites)
	if err != nil {
		spriteLog.Println(err)
		return err
	}

	spriteLog.Println("Sprite config:")
	for _, i := range config.BaseSprites {
		spriteLog.Println("Image file: ", i.ImgPath)
		spriteLog.Println()
		for _, j := range i.Sprites {
			spriteLog.Println("Source Rectangle of ", j.Name)
			spriteLog.Print(" X:", j.SrcRect.X)
			spriteLog.Print(" Y:", j.SrcRect.Y)
			spriteLog.Print(" W:", j.SrcRect.W)
			spriteLog.Print(" H:", j.SrcRect.H)
			spriteLog.Println()
		}
		spriteLog.Println("----------")
	}

	return nil
}
