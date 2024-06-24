package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/components"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryAlgorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryEasingTween"
	"log"
	"math"
	"time"
)

type ComponentControlPanel struct {
	components.Components

	Panel *ControlPanel `wasmPanel:"type:panel"`
}

func (e *ComponentControlPanel) Init() (panel *html.TagDiv, err error) {
	panel, err = e.Components.Init(e)
	return
}

type ControlPanel struct {
	Header string `wasmPanel:"type:headerText;label:Control panel"`
	Body   *Body  `wasmPanel:"type:panelBody"`
}

type Body struct {
	BoatAnimation *BoatAdjust `wasmPanel:"type:component;label:Boat dragging effect"`
}

type OnChangeEvent struct {
	IsTrusted bool    `wasmGet:"isTrusted"`
	Value     float64 `wasmGet:"value"`
	Min       float64 `wasmGet:"min"`
	Max       float64 `wasmGet:"max"`
	Type      string  `wasmGet:"type"`
}

func (e *OnChangeEvent) OnChange(event OnChangeEvent, reference DraggingEffect) {
	log.Printf("reference.TagNumber.GetValue(): %v", reference.TagNumber.GetValue())
	log.Printf("reference.TagRange.GetValue(): %v", reference.TagRange.GetValue())

	var value = reference.TagNumber.GetValue()

	factoryEasingTween.NewInOutSine(
		time.Duration(value)*time.Second,
		0,
		10000,
		tagDivRocket.EasingTweenWalkingAndRotateIntoPoints,
		0,
	).
		SetArgumentsFunc(any(tagDivRocket)).
		SetDoNotReverseMotion()
}

func (e *OnChangeEvent) OnInput(event OnChangeEvent, reference DraggingEffect) {
	switch event.Type {
	case "range":
		reference.TagNumber.Value(reference.MathematicalFormula(event.Min, event.Max, reference.TagRange.GetValue()))
	case "number":
		reference.TagRange.Value(reference.MathematicalFormula(event.Min, event.Max, reference.TagNumber.GetValue()))
	}
}

type DraggingEffect struct {
	components.Range

	TagRange    *html.TagInputRange  `wasmPanel:"type:inputTagRange"`
	TagNumber   *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
	Color       float64              `wasmPanel:"type:value;min:2;max:50;step:1;default:15"`
	ColorChange *OnChangeEvent       `wasmPanel:"type:listener;event:change;func:OnChange"`
	RangeChange *OnChangeEvent       `wasmPanel:"type:listener;event:input;func:OnInput"`
}

func (e *DraggingEffect) MathematicalFormula(min, max, value float64) (result float64) {
	return (max - value) + min
}

func (e *DraggingEffect) Init() {
	e.TagNumber.Value(e.MathematicalFormula(2, 50, e.TagRange.GetValue()))
	e.TagRange.Value(e.MathematicalFormula(2, 50, e.TagNumber.GetValue()))
}

type BoatAdjust struct {
	Dragging *DraggingEffect `wasmPanel:"type:range;label:effect"`
}

type OnClickEvent struct {
	IsTrusted bool   `wasmGet:"isTrusted"`
	Value     string `wasmGet:"value"`
}

func (e *OnClickEvent) OnClick(event OnClickEvent, ref ButtonEvent) {
	log.Printf("Trusted: %v", event.IsTrusted)
	log.Printf("Value:   %v", event.Value)
	ref.Value("Clicked")
}

type ButtonEvent struct {
	components.Button

	Label      string        `wasmPanel:"type:value;default:Ok"`
	RunCommand *OnClickEvent `wasmPanel:"type:listener;event:click;func:OnClick"`
}

func (e *ButtonEvent) Init() {
	e.Value("Initialized")
}

var canvas *html.TagCanvas
var tagDivRocket *html.TagDiv

func main() {
	var err error
	var panel *html.TagDiv

	stage := factoryBrowser.NewStage()

	controlPanel := ComponentControlPanel{}
	if panel, err = controlPanel.Init(); err != nil {
		panic(err)
	}

	canvas = factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight())
	stage.Append(canvas)

	border := 50.0
	wight := 400.0
	height := 400.0

	var bezier = BezierCurve(border, wight, height)
	for _, point := range *bezier.GetProcessed() {
		AddDotBlue(int(point.X), int(point.Y))
	}

	tagDivRocket = factoryBrowser.NewTagDiv().
		Class("animate").
		AddPointsToEasingTween(bezier).
		SetDeltaX(-25).
		SetDeltaY(-25).
		RotateDelta(-math.Pi).
		SetXY(int(1*wight+border), int(0*height+border)).
		Html("<img src=\"boat.png\" alt=\"Imagem\">")
	stage.Append(tagDivRocket)

	stage.Append(panel)

	done := make(chan struct{})
	done <- struct{}{}

}

func BezierCurve(border, wight, height float64) (bezier *algorithm.BezierCurve) {

	// E.g.: P0 (1,0) = (1*wight,0*height)
	// E.g.: P1 (2,0) = (2*wight,0*height)
	// E.g.: P2 (2,1) = (2*wight,1*height)
	//
	//     (0,0)            (1,0)            (2,0)
	//       +----------------+----------------+
	//       | P7            P0             P1 |
	//       |                                 |
	//       |                                 |
	//       |                                 |
	// (0,1) + P6                           P2 + (2,1)
	//       |                                 |
	//       |                                 |
	//       |                                 |
	//       | P5            P4             P3 |
	//       +----------------+----------------+
	//     (0,2)            (1,2)            (2,2)

	bezier = factoryAlgorithm.NewBezierCurve()
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Process()

	return
}

func AddDotBlue(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewBlueHalfTransparent()).
		Arc(x, y, 0.5, 0, 2*math.Pi, false).
		Fill()
}