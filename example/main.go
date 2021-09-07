package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"gitlab.com/go-graphic"
)

func main() {
	var config graphic.Config
	var window graphic.Window
	config.Load("config/window.json", "config/sprite.json")
	window.Init(config.Window)
	spriteIDs, _ := window.LoadSprites(config.BaseSprites)
	_, water, gras, desert := spriteIDs["creature"], spriteIDs["water"], spriteIDs["gras"], spriteIDs["desert"]
	fmt.Println(spriteIDs)
	fmt.Println(len(window.Sprites))

	window.Sprites[water].NewInstance(0, &sdl.FPoint{25, 25})
	window.Sprites[water].NewInstance(0, &sdl.FPoint{125, 75})
	window.Sprites[gras].NewInstance(0, &sdl.FPoint{75, 25})
	window.Sprites[gras].NewInstance(0, &sdl.FPoint{75, 75})
	window.Sprites[desert].NewInstance(0, &sdl.FPoint{125, 25})
	window.Sprites[desert].NewInstance(0, &sdl.FPoint{25, 75})
	fmt.Println("starting render")
	window.Render()
	time.Sleep(2 * time.Second)
	sdl.Quit()
}
