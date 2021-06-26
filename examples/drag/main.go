//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontStyle"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColorNames"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryFont"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryInk"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryText"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"time"
)

var imgSpace html.Image

func main() {

	done := make(chan struct{}, 0)
	stage := global.Global.Stage

	fontText := factoryFont.NewFont(
		14,
		factoryFontFamily.NewHelvetica(),
		factoryFontStyle.NewNotSet(),
	)

	inkText := factoryInk.NewInk(
		0,
		factoryColorNames.NewBlack(),
	)

	text := factoryText.NewText(
		"text",
		&inkText,
		fontText,
		"<-- Drad me",
		135,
		40,
	)
	text.SetDraggableToDesktop()
	stage.AddToDraw(text)

	imgSpace = factoryBrowserImage.NewImage(
		global.Global.Html,
		global.Global.Document.SelfDocument,
		map[string]interface{}{
			"width": 29,
			"heght": 50,
			"id":    "spacecraft",
			"src":   "./small.png",
		},
		true,
		false,
	)

	i := factoryImage.NewImage(
		"id_image",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		nil,
		imgSpace.Get(),
		-100,
		-100,
		29,
		50,
		global.Global.Density,
		global.Global.DensityManager,
	)
	i.SetDraggableToDesktop()
	i.Move(100, 30)
	i.SetOnDragEndFunc(func(x, y int) {
		factoryTween.NewLinear(
			300*time.Millisecond,
			float64(x),
			100,
			false,
			nil,
			func(x float64, arguments ...interface{}) {
				i.MoveX(100)
			},
			nil,
			nil,
			nil,
			func(x, p float64, ars ...interface{}) {
				//log.Printf("x: %v", x)
				//i.Dimensions.X = int(x)
				//i.OutBoxDimensions.X = int(x)
				//i.MoveY(int(x), int(100))
				i.MoveX(int(x))
				//i.Draw()
			},
			0,
		)
		factoryTween.NewLinear(
			300*time.Millisecond,
			float64(y),
			30,
			false,
			nil,
			func(y float64, arguments ...interface{}) {
				i.MoveY(30)
			},
			nil,
			nil,
			nil,
			func(y, p float64, ars ...interface{}) {
				//log.Printf("x: %v", x)
				//i.Dimensions.X = int(x)
				//i.OutBoxDimensions.X = int(x)
				//i.MoveY(int(x), int(100))
				i.MoveY(int(y))
				//i.Draw()
			},
			0,
		)
	})
	stage.AddToDraw(i)

	<-done
}
