package streamxmlp

import (
	"io"
	"bytes"
	"unicode"
	"errors"
	"fmt"
	"html"
)

const (
	Fail = iota
	BeginElem
	EndElem
	Text
	Attr
)

func isXmlName(r rune,inner bool) bool{
	if inner {
		return (unicode.IsLetter(r) || unicode.IsDigit(r) ||
			r==':' || r=='.' || r=='-' )
	}
	return unicode.IsLetter(r)
}

// This is the underlying XmlParser-Implementation. Please don't use it directly.
type XmlParser struct{
	Rd io.RuneReader
	Buffer bytes.Buffer
	Chr rune
	Err error
	AwaitAttrib bool
	elemstack []string
}
func (x *XmlParser)popElem() (r string){
	if x.elemstack==nil { return "" }
	l := len(x.elemstack)
	if l==0 { return "" }
	r = x.elemstack[l-1]
	x.elemstack = x.elemstack[:l-1]
	return
}
func (x *XmlParser)pushElem(r string) {
	x.elemstack = append(x.elemstack,r)
}
func (x *XmlParser)IsOK() bool{
	return x.Err==nil
}
func (x *XmlParser)Next(){ x.Chr,_,x.Err = x.Rd.ReadRune() }
func (x *XmlParser)Trunc(){ x.Buffer.Truncate(0) }
func (x *XmlParser)Write(){ x.Buffer.WriteRune(x.Chr) }

func (x *XmlParser)XmlName() bool{
	x.Trunc()
	inner := false
	for isXmlName(x.Chr,inner) {
		x.Write()
		x.Next()
		inner=true
		if !x.IsOK() { break }
	}
	return inner
}
func (x *XmlParser)Text() {
	x.Trunc()
	for x.Chr!='<' {
		x.Write()
		x.Next()
		if !x.IsOK() { break }
	}
}
func (x *XmlParser)AttrValue() bool{
	x.Trunc()
	switch x.Chr {
	case '"':
		x.Next()
		for x.Chr!='"' {
			x.Write()
			x.Next()
			if !x.IsOK() { break }
		}
		x.Next()
		return true
	case '\'':
		x.Next()
		for x.Chr!='\'' {
			x.Write()
			x.Next()
			if !x.IsOK() { break }
		}
		x.Next()
		return true
	}
	return false
}
func (x *XmlParser)SkipSpace() {
	for unicode.IsSpace(x.Chr) { x.Next() }
}

func (x *XmlParser)xmlDefinition() bool{
	if x.Chr=='?' {
		x.Next()
		i := true
		for !(i && x.Chr=='>') {
			i = (x.Chr=='?')
			x.Next()
		}
		x.Next()
		return true
	}
	return false
}

func (x *XmlParser)commentOrElse() bool{
	if x.Chr=='!' {
		x.Next()
		if x.Chr!='-' { return true }
		x.Next()
		if x.Chr!='-' { return true }
		x.Next()
		i := 0
		for !(i==2 && x.Chr=='>') {
			if x.Chr=='-'{
				if i<2 { i++ }
			} else {
				i = 0
			}
			x.Next()
		}
		x.Next()
		return true
	}
	return false
}

// This Method returns BeginElem, EndElem, Text, Attr or Fail.
// And it sets the 'name', and 'value' variables.
func (x *XmlParser)ReadElement(name,value *string) int{
	readElementStart:
	*name=""
	*value=""
	if !x.IsOK() { return Fail }
	if x.AwaitAttrib {
		x.SkipSpace()
		ending := false
		switch x.Chr {
		case '/':
			x.Next()
			ending = true
			fallthrough
		case '>':
			x.Next()
			x.AwaitAttrib=false
			*name = x.popElem()
			if ending { return EndElem }
		default:
			if !x.XmlName() {
				x.Err = errors.New("expected xml-Name or '(/)?>'")
				return Fail
			}
			*name = x.Buffer.String()
			if x.Chr!='=' {
				x.Err = errors.New(fmt.Sprintf(" expected '=' but found '%c'",x.Chr))
				return Fail
			}
			x.Next()
			if !x.AttrValue() {
				x.Err = errors.New("expected Value or '(/)?>'")
				return Fail
			}
			*value = html.UnescapeString(x.Buffer.String())
			return Attr
		}
	}
	end := false
	switch x.Chr {
	case '<':
		x.Next()
		x.SkipSpace()
		if x.commentOrElse() { goto readElementStart }
		if x.xmlDefinition() { goto readElementStart }
		if x.Chr=='/' {
			end = true
			x.Next()
			x.SkipSpace()
		}
		if !x.XmlName() {
			x.Err = errors.New("expected xml-Name")
			return Fail
		}
		*name = x.Buffer.String()
		if end {
			x.popElem()
			x.SkipSpace()
			if x.Chr!='>' {
				x.Err = errors.New(fmt.Sprintf(" a tag has to end with '>' but ended '%c'",x.Chr))
				return Fail
			}
			x.Next()
			return EndElem
		} else {
			x.pushElem(*name)
			x.AwaitAttrib = true
			return BeginElem
		}
	default:
		x.Text()
		*value = html.UnescapeString(x.Buffer.String())
		return Text
	}
	return Fail
}
