// This package implements a Streaming API for writing XML.
// 
// The Programmer manually calls StartElement(), EndElement(), Attribute() and Text() Methods.
// 
// This Package has only Intefaces, Not Implementations
package push

import "errors"

var MisplacedAttribute error = errors.New("MisplacedAttribute")

type XmlWriter interface{
	StartElement(name string) error
	EndElement() error
	Text(text string) error
	
	// Must be Called after StartElement(..)
	Attribute(name,value string) error
	
	// Flushes the Buffer
	Flush() error
}

type XmlWriteCloser interface{
	XmlWriter
	
	// Close implies Flush
	Close() error
}

