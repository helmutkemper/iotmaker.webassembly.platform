package canvas

import "syscall/js"

func NewCanvasWith2DContext(document js.Value, id string, width, height int) Canvas {
	el := Canvas{}
	el.SelfElement = document

	el.InitializeContext2DById(id)

	el.SelfElement.Set("width", width)
	el.SelfElement.Set("height", height)
	el.SelfContext = el.SelfElement.Call("getContext", "2d")

	return el
}