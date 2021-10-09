package htmltextareaelement

import (
	"syscall/js"

	"github.com/realPy/hogosuru/event"
)

func (h HtmlTextAreaElement) OnInput(handler func(e event.Event)) (js.Func, error) {

	return h.AddEventListener("input", handler)
}
