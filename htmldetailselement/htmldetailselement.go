package htmldetailselement

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/document"
	"github.com/realPy/hogosuru/element"
	"github.com/realPy/hogosuru/htmlelement"
)

var singleton sync.Once

var htmldetailselementinterface js.Value

//HtmlDetailsElement struct
type HtmlDetailsElement struct {
	htmlelement.HtmlElement
}

type HtmlDetailsElementFrom interface {
	HtmlDetailsElement() HtmlDetailsElement
}

func (h HtmlDetailsElement) HtmlDetailsElement() HtmlDetailsElement {
	return h
}

func GetInterface() js.Value {

	singleton.Do(func() {
		var err error
		if htmldetailselementinterface, err = js.Global().GetWithErr("HTMLDetailsElement"); err != nil {
			htmldetailselementinterface = js.Null()
		}

	})

	baseobject.Register(htmldetailselementinterface, func(v js.Value) (interface{}, error) {
		return NewFromJSObject(v)
	})

	return htmldetailselementinterface
}

func New(d document.Document) (HtmlDetailsElement, error) {
	var err error

	var h HtmlDetailsElement
	var e element.Element

	if e, err = d.CreateElement("details"); err == nil {
		h, err = NewFromElement(e)
	}
	return h, err
}

func NewFromElement(elem element.Element) (HtmlDetailsElement, error) {
	var h HtmlDetailsElement
	var err error

	if hci := GetInterface(); !hci.IsNull() {
		if elem.BaseObject.JSObject().InstanceOf(hci) {
			h.BaseObject = h.SetObject(elem.BaseObject.JSObject())

		} else {
			err = ErrNotAnHtmlDetailsElement
		}
	} else {
		err = ErrNotImplemented
	}

	return h, err
}

func NewFromJSObject(obj js.Value) (HtmlDetailsElement, error) {
	var h HtmlDetailsElement

	if hci := GetInterface(); !hci.IsNull() {
		if obj.InstanceOf(hci) {

			h.BaseObject = h.SetObject(obj)
			return h, nil
		}
	}
	return h, ErrNotAnHtmlDetailsElement
}

func (h HtmlDetailsElement) Open() (bool, error) {
	return h.GetAttributeBool("open")
}

func (h HtmlDetailsElement) SetOpen(value bool) error {
	return h.SetAttributeBool("open", value)
}
