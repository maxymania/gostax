package streamxmlp

import (
	"io"
	"bufio"
	"xml.api/pull"
)

// This is a pull.XmlReader Implementation
type PullXmlReader struct{
	parser XmlParser
	state int
	nodeType pull.NodeType
	name string
	value string
}

// Creates a XmlReader From the given Reader.
// This equals NewPullParser(bufio.NewReader(src)).
func NewPullParserFromReader(src io.Reader) *PullXmlReader{
	pp := new(PullXmlReader)
	pp.parser.Rd = bufio.NewReader(src)
	pp.parser.Next()
	return pp
}

// Creates a XmlReader From the given RuneReader.
func NewPullParser(src io.RuneReader) *PullXmlReader{
	pp := new(PullXmlReader)
	pp.parser.Rd = src
	pp.parser.Next()
	return pp
}

func (xr *PullXmlReader) Read() bool {
	const (
		stateBegin = iota
		stateNormal
		stateEnd
	)
	xr.name=""
	xr.value=""
	switch xr.state {
	case stateBegin:
		xr.nodeType = pull.StartDocument
		xr.state = stateNormal
		return true
	case stateNormal:
		switch xr.parser.ReadElement(&xr.name,&xr.value) {
		case Fail:
			if xr.parser.Err==io.EOF {
				xr.nodeType = pull.EndDocument
				xr.state = stateEnd
				return true
			}
			return false
		case BeginElem: xr.nodeType = pull.StartElement
		case EndElem: xr.nodeType = pull.EndElement
		case Text: xr.nodeType = pull.Text
		case Attr: xr.nodeType = pull.Attribute
		}
		return true
	case stateEnd:
		xr.nodeType = pull.None
		return false
	}
	return false
}
func (xr *PullXmlReader) NodeType() pull.NodeType { return xr.nodeType }
func (xr *PullXmlReader) Name() string { return xr.name }
func (xr *PullXmlReader) Value() string { return xr.value }
func (xr *PullXmlReader) GetError() error { return xr.parser.Err }

