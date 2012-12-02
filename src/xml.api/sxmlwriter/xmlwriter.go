package sxmlwriter

import "io"
import "fmt"
import "html"
import "xml.api/push"

const xmlheader = `<?xml version="1.0" encoding="UTF-8" ?>` + "\r\n"

type xmlFprinter struct{
	dest io.Writer
	inStartElement bool
	stack []string
}
func (x *xmlFprinter) push(v string) { x.stack=append(x.stack,v) }
func (x *xmlFprinter) pop() (v string) {
	if x.stack!=nil && len(x.stack)>0 {
		l := len(x.stack)-1
		v,x.stack = x.stack[l],x.stack[:l]
	}
	return
}
func (x *xmlFprinter) Flush() error {
	if x.inStartElement {
		x.inStartElement=false
		_,e := fmt.Fprintf(x.dest,`>`)
		return e
	}
	return nil
}
func (x *xmlFprinter) StartElement(name string) error {
	if e:=x.Flush(); e!=nil { return e }
	x.push(name)
	_,e := fmt.Fprintf(x.dest,`<%s`,name)
	x.inStartElement=true
	return e
}
func (x *xmlFprinter) EndElement() error {
	if x.inStartElement{
		x.inStartElement=false
		x.pop()
		_,e := fmt.Fprintf(x.dest,`/>`)
		return e
	}
	if e:=x.Flush(); e!=nil { return e }
	_,e := fmt.Fprintf(x.dest,`</%s>`,x.pop())
	return e
}
func (x *xmlFprinter) Text(text string) error {
	if e:=x.Flush(); e!=nil { return e }
	_,e := fmt.Fprintf(x.dest,`%s`,html.EscapeString(text))
	return e
}
func (x *xmlFprinter) Attribute(name,value string) error {
	if !x.inStartElement { return push.MisplacedAttribute }
	_,e := fmt.Fprintf(x.dest,` %s="%s"`,name,html.EscapeString(value))
	return e
}

func NewXmlWriter(w io.Writer) push.XmlWriter{
	// xmlheader
	fmt.Fprintf(w,xmlheader)
	return &xmlFprinter{dest:w}
}
