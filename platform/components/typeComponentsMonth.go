package components

import (
	"github.com/helmutkemper/webassembly/browser/html"
	"time"
)

//todo: disabled

// __monthOnInputEvent faz a captura de dados do event input
type __monthOnInputEvent struct {
	Value float64 `wasmGet:"value"`
}

// Month
//
// English:
//
//	 Component width label, input month and input number
//
//	Public methods
//	  * Max(int64/float64)
//	  * Min(int64/float64)
//	  * Step(int64/float64)
//	  * Value(int64/float64)
//
//	Methods for internal use, should not be used by you and should not be shadowed (rewritten)
//	  * OnChangeNumber
//	  * OnChangeMonth
//
//	     Example:
//
//	     +- panel --------------------------------------------------------+
//	     |                                                                |
//	     |  +- panelCel -----------------------------------------------+  |
//	     |  |                                                          |  |
//	     |  | +- labelCel -------------------------------------------+ |  |
//	     |  | | Label                                              ˇ | |  |
//	     |  | +- compCel --------------------------------------------+ |  |
//	     |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |
//	     |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |
//	     |  | | Text inside component  ⊢--*-----------------⊣  [ n ] | |  |
//	     |  | +------------------------------------------------------+ |  |
//	     |  |                                                          |  |
//	     |  +----------------------------------------------------------+  |
//	     |                                                                |
//	     +----------------------------------------------------------------+
//
//	    // ColorMonth Component type Month
//	    type ColorMonth struct {
//	      // Mount the component
//	      components.Month
//
//	      // Field with the value of the element must be type int64 or float64
//	      Color       float64              `wasmPanel:"type:value;min:0;max:50;step:1;default:0"`
//	      // [optional] Reference to the month component
//	      TagMonth    *html.TagInputMonth  `wasmPanel:"type:inputTagMonth"`
//	      // [optional] Reference to the number component
//	      TagNumber   *html.TagInputNumber `wasmPanel:"type:inputTagNumber"`
//	      // [optional] Event pointer. Create one for each event, change, click, etc.
//	      ColorChange *OnChangeEvent       `wasmPanel:"type:listener;event:change;func:OnChange"`
//	    }
//
//	    // Init is invoked when the component is initialized
//	    func (e *ColorMonth) Init() { // [optional]
//	      // anything you want
//
//	      // you can change the default properties
//	      e.Step(1)
//	      e.Max(255)
//	      e.Min(0)
//	      e.Value(0)
//	    }
//
//	    // OnChangeEvent Component data capture object. The values are isTrusted, min, max, step and value
//	    type OnChangeEvent struct {
//
//	      // Indicates that the value was generated by the user and not by another method
//	      IsTrusted bool    `wasmGet:"isTrusted"`
//	      // Returns the value of the component
//	      Value     float64 `wasmGet:"value"`
//	    }
//
//	    // OnChange is the function determined by the tag `wasmPanel:"type:listener;event:change;func:OnChange"`.
//	    // `func:OnChange` determines the name of the public function to be invoked in the event
//	    func (e *OnChangeEvent) OnChange(event OnChangeEvent) {
//	      log.Printf("Trusted: %+v", event.IsTrusted)
//	      log.Printf("Value:   %+v", event.Value)
//	    }
type Month struct {
	__max   any
	__min   any
	__step  any
	__value any

	// __monthOnInputEvent is the pointer sent when the `change` event happens
	__change *__monthOnInputEvent

	// reference of component elements
	__monthTag *html.TagInputMonth `wasmPanel:"type:inputTagMonth"`
}

// OnChangeMonth Event called when the component value changes.
//
//	This function must not be called by the user, and it must not be shadowed.
//	It is public so that reflect can access it.
func (e *Month) OnChangeMonth(stt __monthOnInputEvent, ref any) {
	e.__monthTag.Value(stt.Value)
}

func (e *Month) init() {
	if e.__max != nil {
		e.max(e.__max)
		e.__max = nil
	}

	if e.__min != nil {
		e.min(e.__min)
		e.__min = nil
	}

	if e.__step != nil {
		e.step(e.__step)
		e.__step = nil
	}

	if e.__value != nil {
		e.value(e.__value)
		e.__value = nil
	}
}

func (e *Month) max(max any) {

	switch converted := max.(type) {
	case time.Time:
		e.__monthTag.Max(converted.Format("2006-01"))
	default:
		e.__monthTag.Max(max)
	}

}

func (e *Month) min(min any) {

	switch converted := min.(type) {
	case time.Time:
		e.__monthTag.Min(converted.Format("2006-01"))
	default:
		e.__monthTag.Min(min)
	}

}

func (e *Month) step(step any) {

	switch converted := step.(type) {
	case time.Time:
		e.__monthTag.Step(converted.Format("2006-01"))
	default:
		e.__monthTag.Step(step)
	}

}

func (e *Month) value(value any) {

	switch converted := value.(type) {
	case time.Time:
		e.__monthTag.Value(converted.Format("2006-01"))
	default:
		e.__monthTag.Value(value)
	}

}

// Max Sets the maximum value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Month) Max(max any) {
	if e.__monthTag == nil {
		e.__max = max
		return
	}

	e.max(max)
}

// Min Sets the minimum value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Month) Min(min any) {
	if e.__monthTag == nil {
		e.__min = min
		return
	}

	e.min(min)
}

// Step Sets the step value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Month) Step(step any) {
	if e.__monthTag == nil {
		e.__step = step
		return
	}

	e.step(step)
}

// Value Sets the value of the component.
//
//	When the component was created, a field within the configuration struct received the tag `wasmPanel:"type:value"`
//	and this was created as an int64 or float64 type, therefore the value passed to the function must respect the type
func (e *Month) Value(value any) {
	if e.__monthTag == nil {
		e.__value = value
		return
	}

	e.value(value)
}
