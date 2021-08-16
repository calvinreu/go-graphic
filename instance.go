package graphic

import "github.com/veandco/go-sdl2/sdl"

//Instance position Angle and the Center of an instance of a sprite
type Instance struct {
	DestRect sdl.FRect
	Angle    float64
	Center   sdl.FPoint
}

//NewPosition sets the position of this instance Center is the Center of the instances new position
func (instance *Instance) NewPosition(Center *sdl.FPoint) {
	instance.Center.X = Center.X
	instance.Center.Y = Center.Y
	instance.DestRect.X = Center.X - (instance.DestRect.W / 2)
	instance.DestRect.Y = Center.Y - (instance.DestRect.H / 2)
}

//ChangePosition moves the instance by x, y
func (instance *Instance) ChangePosition(x, y float32) {
	instance.Center.X += x
	instance.Center.Y += y
	instance.DestRect.X = instance.Center.X - (instance.DestRect.W / 2)
	instance.DestRect.Y = instance.Center.Y - (instance.DestRect.H / 2)
}
