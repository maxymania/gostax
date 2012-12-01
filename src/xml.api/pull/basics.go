// This package implements a Streaming API for reading XML.
// Traditionally, XML APIs are either:
// 
// * DOM based - the entire document is read into memory as a tree structure for random access by the calling application;
// 
// * event based - the application registers to receive events as entities are encountered within the source document.
// 
// The traditional Go way is
// * object Based - the entire document is read into memory as a object structure (de-/serializing).
//
// In Contrast pull/StAX has a Reader-Based Approach:
// The Programmer Manually Calls Read() for any Element, thus the Document Model isn't Bound to a Specific Model (Go types).
// 
// This Package has only Intefaces, Not Implementations
package pull

import "strconv"

type NodeType int16

const (
	None NodeType = iota
	
	// The First Node of all Nodes in a XML-Document
	StartDocument
	
	// The Last Node of all Nodes in a XML-Document
	EndDocument
	
	// The Start of an Element
	StartElement
	
	// The End of an Element
	EndElement
	
	// Just Text
	Text
	
	// The Attribute of an StartElement
	// it is assigned to the Previously received StartElement-Element
	Attribute
)

var ntNames = map[NodeType]string {
	None          :"None",
	StartDocument :"StartDocument",
	EndDocument   :"EndDocument",
	StartElement  :"StartElement",
	EndElement    :"EndElement",
	Text          :"Text",
	Attribute     :"Attribute"}

func (nt NodeType) String() string {
	s, ok := ntNames[nt]
	if ok { return s }
	return strconv.Itoa(int(nt))
}
