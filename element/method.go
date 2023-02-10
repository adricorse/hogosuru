package element

import (
	"errors"
	"syscall/js"

	"github.com/realPy/hogosuru/array"
	"github.com/realPy/hogosuru/attr"
	"github.com/realPy/hogosuru/baseobject"
	"github.com/realPy/hogosuru/domrect"
	"github.com/realPy/hogosuru/domrectlist"
	"github.com/realPy/hogosuru/htmlcollection"
	"github.com/realPy/hogosuru/node"
	"github.com/realPy/hogosuru/nodelist"
	"github.com/realPy/hogosuru/object"
)

func (e Element) attachShadow() {
	//TODO IMPLEMENT
}

func (e Element) After(params ...interface{}) error {
	var err error
	var arrayJS []interface{}

	for _, param := range params {
		switch p := param.(type) {
		case node.Node:
			arrayJS = append(arrayJS, p.JSObject())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	_, err = e.Call("after", arrayJS...)

	return err
}

func (e Element) Animate(keyframes, options interface{}) error {
	var argCall []interface{}

	var err error
	if keyframesObject, ok := keyframes.(array.ArrayFrom); ok {
		argCall = append(argCall, keyframesObject.Array_().JSObject())

	}

	if keyframesObject, ok := keyframes.(object.ObjectFrom); ok {
		argCall = append(argCall, keyframesObject.Object_().JSObject())
	}

	if optionsObject, ok := keyframes.(object.ObjectFrom); ok {
		argCall = append(argCall, optionsObject.Object_().JSObject())
	} else {
		argCall = append(argCall, js.ValueOf(options))
	}
	_, err = e.Call("animate")
	return err
}

func (e Element) Append(params ...interface{}) error {
	var err error
	var arrayJS []interface{}

	for _, param := range params {
		switch p := param.(type) {
		case node.Node:
			arrayJS = append(arrayJS, p.JSObject())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	_, err = e.Call("append", arrayJS...)

	return err
}

func (e Element) Before(params ...interface{}) error {
	var err error
	var arrayJS []interface{}

	for _, param := range params {
		switch p := param.(type) {
		case node.Node:
			arrayJS = append(arrayJS, p.JSObject())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	_, err = e.Call("before", arrayJS...)

	return err
}

func (e Element) Closest(selectors string) (Element, error) {
	var err error
	var obj js.Value
	var elem Element

	if obj, err = e.Call("closest", js.ValueOf(selectors)); err == nil {
		return NewFromJSObject(obj)
	}

	return elem, err
}

func (e Element) computedStyleMap() {
	//TODO IMPLEMENT
}

func (e Element) getAnimations() {
	//TODO IMPLEMENT
}

func (e Element) GetAttribute(attributeName string) (string, error) {
	var err error
	var obj js.Value
	var newstr string

	if obj, err = e.Call("getAttribute", js.ValueOf(attributeName)); err == nil {
		if obj.IsNull() {
			return newstr, ErrAttributeEmpty
		}
		return obj.String(), err
	}
	return newstr, err
}

func (e Element) GetAttributeNames() (array.Array, error) {
	var err error
	var obj js.Value
	var arr array.Array

	if obj, err = e.Call("getAttributeNames"); err == nil {
		if obj.IsNull() {
			return arr, ErrAttributeEmpty
		}
		return array.NewFromJSObject(obj)
	}
	return arr, err
}

func (e Element) GetAttributeNode(attrName string) (attr.Attr, error) {
	var err error
	var obj js.Value
	var newobj attr.Attr

	if obj, err = e.Call("getAttributeNode", js.ValueOf(attrName)); err == nil {
		if obj.IsNull() {
			return newobj, ErrAttributeEmpty
		}
		return attr.NewFromJSObject(obj)
	}
	return newobj, err
}

func (e Element) GetAttributeNodeNS(namespace, nodeName string) (attr.Attr, error) {
	var err error
	var obj js.Value
	var newobj attr.Attr

	if obj, err = e.Call("getAttributeNodeNS", js.ValueOf(namespace), js.ValueOf(nodeName)); err == nil {
		if obj.IsNull() {
			return newobj, ErrAttributeEmpty
		}
		return attr.NewFromJSObject(obj)
	}
	return newobj, err
}

func (e Element) GetAttributeNS(namespace, name string) (string, error) {
	var err error
	var obj js.Value
	var newobj string

	if obj, err = e.Call("getAttributeNS", js.ValueOf(namespace), js.ValueOf(name)); err == nil {
		if obj.IsNull() {
			return newobj, ErrAttributeEmpty
		}
		return obj.String(), err
	}
	return newobj, err
}

func (e Element) GetBoundingClientRect() (domrect.DOMRect, error) {
	var err error
	var obj js.Value
	var newdomrect domrect.DOMRect

	if obj, err = e.Call("getBoundingClientRect"); err == nil {
		return domrect.NewFromJSObject(obj)
	}
	return newdomrect, err
}

// retourne un DOMRectList
func (e Element) GetClientRects() (domrectlist.DOMRectList, error) {
	var err error
	var obj js.Value
	var arr domrectlist.DOMRectList

	if obj, err = e.Call("getClientRects"); err == nil {
		return domrectlist.NewFromJSObject(obj)
	}
	return arr, err
}

func (e Element) GetElementsByClassName(names string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = e.Call("getElementsByClassName", js.ValueOf(names)); err == nil {
		if !obj.IsNull() {
			return htmlcollection.NewFromJSObject(obj)
		}
		return collection, ErrElementsNotFound
	}

	return collection, err
}

func (e Element) GetElementsByTagName(tagName string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = e.Call("getElementsByTagName", js.ValueOf(tagName)); err == nil {
		if obj.IsNull() || obj.IsUndefined() {
			return collection, ErrElementsNotFound
		}
		return htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (e Element) GetElementsByTagNameNS(namespaceURI, localName string) (htmlcollection.HtmlCollection, error) {
	var err error
	var obj js.Value
	var collection htmlcollection.HtmlCollection

	if obj, err = e.Call("getElementsByTagNameNS", js.ValueOf(namespaceURI), js.ValueOf(localName)); err == nil {
		if obj.IsNull() || obj.IsUndefined() {
			return collection, ErrElementsNotFound
		}
		return htmlcollection.NewFromJSObject(obj)
	}

	return collection, err
}

func (e Element) HasAttribute(name string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("hasChildNodes", js.ValueOf(name)); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), err
		}
		return result, baseobject.ErrObjectNotBool
	}

	return result, err

}

func (e Element) HasAttributeNS(namespace, localName string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("hasAttributesNS", js.ValueOf(namespace), js.ValueOf(localName)); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), err
		}
		return result, baseobject.ErrObjectNotBool
	}
	return result, err
}

func (e Element) HasAttributes() (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("hasAttributes"); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), err
		}
		return result, baseobject.ErrObjectNotBool
	}
	return result, err
}

func (e Element) HasPointerCapture(pointerId int) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("hasPointerCapture", js.ValueOf(pointerId)); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), err
		}
		return result, baseobject.ErrObjectNotBool
	}
	return result, err
}

func (e Element) InsertAdjacentElement(position string, element Element) (Element, error) {
	var elemObject js.Value
	var newelem Element
	var err error

	if elemObject, err = e.Call("insertAdjacentElement", js.ValueOf(position), element.JSObject()); err == nil {
		if elemObject.IsNull() {
			return newelem, ErrInsertAdjacent
		}
		return element, err
	}
	return newelem, err
}

func (e Element) InsertAdjacentHTML(position string, text string) error {
	var err error

	_, err = e.Call("insertAdjacentHTML", js.ValueOf(position), js.ValueOf(text))
	return err
}

func (e Element) InsertAdjacentText(where string, data string) error {
	var err error

	_, err = e.Call("insertAdjacentText", js.ValueOf(where), js.ValueOf(data))
	return err
}

func (e Element) Matches(selectors string) (bool, error) {
	var err error
	var obj js.Value
	var result bool

	if obj, err = e.Call("matches", js.ValueOf(selectors)); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), err
		}
		return result, baseobject.ErrObjectNotBool
	}
	return result, err
}

func (e Element) Prepend(params ...interface{}) error {
	var err error
	var arrayJS []interface{}

	for _, param := range params {
		switch p := param.(type) {
		case node.Node:
			arrayJS = append(arrayJS, p.JSObject())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	_, err = e.Call("preprend", arrayJS...)

	return err
}

func (e Element) QuerySelector(selectors string) (node.Node, error) {
	var err error
	var obj js.Value
	var nod node.Node

	if obj, err = e.Call("querySelector", js.ValueOf(selectors)); err == nil {
		if !obj.IsNull() {
			return node.NewFromJSObject(obj)
		}
		return nod, errors.New(ErrElementNotFound.Error() + "" + selectors)
	}
	return nod, err
}

func (e Element) QuerySelectorAll(selectors string) (nodelist.NodeList, error) {
	var err error
	var obj js.Value
	var nlist nodelist.NodeList

	if obj, err = e.Call("querySelectorAll", js.ValueOf(selectors)); err == nil {
		if !obj.IsNull() {
			return nodelist.NewFromJSObject(obj)
		}
		return nlist, errors.New(ErrElementsNotFound.Error() + "" + selectors)
	}
	return nlist, err
}

func (e Element) ReleasePointerCapture(pointerId int) error {
	var err error
	_, err = e.Call("releasePointerCapture", js.ValueOf(pointerId))
	return err
}
func (e Element) Remove() error {
	var err error
	_, err = e.Call("remove")
	return err
}

func (e Element) RemoveAttribute(attrName string) error {
	var err error
	_, err = e.Call("removeAttribute", js.ValueOf(attrName))
	return err
}

func (e Element) removeAttributeNode() {
	//TODO IMPLEMENT
}

func (e Element) RemoveAttributeNS(namespace, attrName string) error {
	var err error
	_, err = e.Call("removeAttributeNS", js.ValueOf(namespace), js.ValueOf(attrName))
	return err
}

func (e Element) ReplaceChildren(params ...interface{}) error {
	var err error
	var arrayJS []interface{}
	for _, param := range params {
		switch p := param.(type) {
		case node.Node:
			arrayJS = append(arrayJS, p.JSObject())
		case string:
			arrayJS = append(arrayJS, js.ValueOf(p))
		default:
			return ErrSendUnknownType
		}
	}

	_, err = e.Call("replaceChildren", arrayJS...)

	return err
}

func (e Element) RequestFullscreen() error {
	var err error
	_, err = e.Call("requestFullscreen")
	return err
}

func (e Element) RequestPointerLock() error {
	var err error
	_, err = e.Call("requestPointerLock")
	return err
}

func (e Element) Scroll(params ...interface{}) error {
	var err error
	var optJSValue []interface{}

	if len(params) == 1 {
		if options, ok := params[0].(map[string]interface{}); ok {
			optJSValue = append(optJSValue, js.ValueOf(options))
			_, err = e.Call("scroll", optJSValue)
			return err
		}
		return ErrAttributeEmpty
	}
	if len(params) == 2 {
		if _, ok := params[0].(int); !ok {
			return ErrSendUnknownType
		}
		if _, ok := params[1].(int); !ok {
			return ErrSendUnknownType
		}
		optJSValue = append(optJSValue, js.ValueOf(params[0]))
		optJSValue = append(optJSValue, js.ValueOf(params[1]))
		_, err = e.Call("scroll", optJSValue...)
		return err
	}
	return ErrAttributeEmpty
}

func (e Element) ScrollBy(params ...interface{}) error {
	var err error
	var optJSValue []interface{}

	if len(params) == 1 {
		if options, ok := params[0].(map[string]interface{}); ok {
			optJSValue = append(optJSValue, js.ValueOf(options))
			_, err = e.Call("scrollBy", optJSValue)
			return err
		}
		return ErrAttributeEmpty
	}
	if len(params) == 2 {
		if _, ok := params[0].(int); !ok {
			return ErrSendUnknownType
		}
		if _, ok := params[1].(int); !ok {
			return ErrSendUnknownType
		}
		optJSValue = append(optJSValue, js.ValueOf(params[0]))
		optJSValue = append(optJSValue, js.ValueOf(params[1]))
		_, err = e.Call("scrollBy", optJSValue...)
		return err
	}
	return ErrAttributeEmpty
}

func (e Element) ScrollIntoView() error {
	var err error
	_, err = e.Call("scrollBy")
	return err
}

func (e Element) ScrollTo(params ...interface{}) error {
	var err error
	var optJSValue []interface{}

	if len(params) == 1 {
		if options, ok := params[0].(map[string]interface{}); ok {
			optJSValue = append(optJSValue, js.ValueOf(options))
			_, err = e.Call("scrollTo", optJSValue)
			return err
		}
		return ErrAttributeEmpty
	}
	if len(params) == 2 {
		if _, ok := params[0].(int); !ok {
			return ErrSendUnknownType
		}
		if _, ok := params[1].(int); !ok {
			return ErrSendUnknownType
		}
		optJSValue = append(optJSValue, js.ValueOf(params[0]))
		optJSValue = append(optJSValue, js.ValueOf(params[1]))
		_, err = e.Call("scrollTo", optJSValue...)
		return err
	}
	return ErrAttributeEmpty
}

func (e Element) SetAttribute(name, value string) error {
	var err error
	_, err = e.Call("setAttribute", js.ValueOf(name), js.ValueOf(value))
	return err
}

func (e Element) setAttributeNode() {
	//TODO IMPLEMENT
}

func (e Element) setAttributeNodeNS() {
	//TODO IMPLEMENT
}

func (e Element) SetAttributeNS(namespace, name, value string) error {
	var err error
	_, err = e.Call("setAttributeNS", js.ValueOf(namespace), js.ValueOf(name), js.ValueOf(value))
	return err
}

func (e Element) SetPointerCapture(pointerId int) error {
	var err error
	_, err = e.Call("setPointerCapture", js.ValueOf(pointerId))
	return err
}

func (e Element) ToggleAttribute(name string, force ...interface{}) (bool, error) {
	var err error
	var optJSValue []interface{}
	var obj js.Value
	var result bool

	optJSValue = append(optJSValue, js.ValueOf(name))
	if force != nil && len(force) == 1 {
		optJSValue = append(optJSValue, js.ValueOf(force[0]))
	}

	if obj, err = e.Call("toggleAttribute", optJSValue...); err == nil {
		if obj.Type() == js.TypeBoolean {
			return obj.Bool(), err
		}
		return result, baseobject.ErrObjectNotBool
	}
	return result, err
}
