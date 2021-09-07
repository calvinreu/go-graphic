//Package graphic is using the sdl2 go interface from (c)https://github.com/veandco/go-sdl2/ under the BSD 3 License
package graphic

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

//Window contains the information required to render a window with diffrent Sprites
type Window struct {
	Sprites  []Sprite
	Renderer *sdl.Renderer
	window   *sdl.Window
	fps      uint32
	logger   log.Logger
}

//RunOutput will render FPS frames every second until running is false
func (graphic Window) RunOutput(stop chan bool) {
	var timeStamp, frameTime uint32
	frameTime = 1000 / graphic.fps

	for {
		select {
		case <-stop:
			return
		default:
			timeStamp = sdl.GetTicks()
			graphic.Render()
			if sdl.GetTicks()-timeStamp < frameTime {
				sdl.Delay(frameTime - (sdl.GetTicks() - timeStamp))
			}
		}
	}
}

//Render renders the information from the graphic object to the screen
func (graphic *Window) Render() {

	graphic.Renderer.SetDrawColor(0, 0, 0, 1)
	graphic.Renderer.Clear()

	for _, i := range graphic.Sprites {
		for j := i.instances.Front(); j != nil; j = j.Next() {
			if instance, ok := j.Value.(*Instance); ok {
				graphic.Renderer.CopyExF(i.texture, &i.srcRect, &instance.DestRect, instance.Angle, &instance.Center, sdl.FLIP_HORIZONTAL)
			} else {
				fmt.Println("list of sprite does not contain Instances")
			}
		}
	}
	graphic.Renderer.Present()
}

//New returns a Window object with initialized renderer and window note that Sprites have to be added manual
func (graphic *Window) Init(config WindowConfig) error {
	err := InitLogger(&graphic.logger, config.Title)
	if err != nil {
		graphic.logger.Println(err)
		fmt.Println("logs are now disabled")
	}

	err = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER)
	if err != nil {
		graphic.logger.Println(err)
		return err
	}

	graphic.window, err = sdl.CreateWindow(config.Title, config.X, config.Y, config.Width, config.Height, config.WindowFlags)
	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO | sdl.INIT_TIMER)
		graphic.logger.Println(err)
		return err
	}

	graphic.Renderer, err = sdl.CreateRenderer(graphic.window, -1, config.RendererFlags)

	if err != nil {
		sdl.QuitSubSystem(sdl.INIT_VIDEO | sdl.INIT_TIMER)
		graphic.logger.Println(err)
		return err
	}

	graphic.logger.Println("Created Window " + config.Title + " succesfully")
	return nil
}

//LoadSprites from config object returns map with sprite IDs linked to the sprite name
func (graphic *Window) LoadSprites(config []SpriteBaseConfig) (map[string]uint32, error) {
	spriteIDs := make(map[string]uint32)

	for _, i := range config {
		spriteID, err := graphic.AddSprite(i.ImgPath, i.Sprites[0].SrcRect)
		spriteIDs[i.Sprites[0].Name] = spriteID
		if err != nil {
			graphic.logger.Println("called from LoadSprites")
			return spriteIDs, err
		}

		for _, j := range i.Sprites[1:] {
			spriteIDs[j.Name] = graphic.AddSpriteByID(spriteID, j.SrcRect)
		}
	}

	return spriteIDs, nil
}

//AddSprite adds another sprite which can be used be creating a instance of it see Sprite.NewInstance
func (graphic *Window) AddSprite(imgPath string, srcRect sdl.Rect) (uint32, error) {
	var err error
	var sprite Sprite
	retIndex := len(graphic.Sprites)

	sprite, err = NewSprite(graphic.Renderer, imgPath, srcRect)
	if err != nil {
		graphic.logger.Println(err)
		return 0, err
	}
	graphic.Sprites = append(graphic.Sprites, sprite)

	return uint32(retIndex), nil
}

//AddSpriteByID adds another sprite with the same texture as sprite with id spriteID
func (graphic *Window) AddSpriteByID(spriteID uint32, srcRect sdl.Rect) uint32 {
	var sprite Sprite

	if len(graphic.Sprites)-1 < (int)(spriteID) {
		graphic.logger.Println("sprite: ", spriteID, " does not exist")
		return 0
	}

	sprite.texture = graphic.Sprites[spriteID].texture
	sprite.srcRect = srcRect

	retIndex := len(graphic.Sprites)
	return uint32(retIndex)
}
