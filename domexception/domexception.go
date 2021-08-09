package domexception

import (
	"sync"
	"syscall/js"

	"github.com/realPy/hogosuru/baseobject"
)

var singleton sync.Once

var domexceptioninterface js.Value

//DomException DomException struct
type DomException struct {
	baseobject.BaseObject
}

type DomExceptionFrom interface {
	DomException() DomException
}

func (d DomException) DomException() DomException {
	return d
}

//GetJSInterface get teh JS interface of broadcast channel
func GetInterface() js.Value {

	singleton.Do(func() {

		var err error
		if domexceptioninterface, err = js.Global().GetWithErr("DOMException"); err != nil {
			domexceptioninterface = js.Null()
		}

	})
	return domexceptioninterface
}

func NewFromJSObject(obj js.Value) (DomException, error) {
	var d DomException
	var err error
	if di := GetInterface(); !di.IsNull() {
		if obj.InstanceOf(di) {
			d.BaseObject = d.SetObject(obj)
		} else {
			err = ErrNotADOMException
		}
	} else {
		err = ErrNotImplemented
	}

	return d, err
}

func (d DomException) Message() (string, error) {
	return d.GetAttributeString("message")
}

func (d DomException) Name() (string, error) {
	return d.GetAttributeString("name")
}
