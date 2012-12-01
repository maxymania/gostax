Go Xml Apis
===========

Go Streaming API for XML (and More)

## Why? Go has an Xml De-/Serializer!

Well. The go package "encoding/xml" serializes Go Objects to Xml and deserializes Xml to Go Objects.
If you want to parse a specific XML Format, it is the Default Go Way to create a Datastructure in go, that matches to the
XML-Document-Structure. In Regular Cases this Model is very Productive.

But what is, if you just want a Generic inteface to XML such as StAX, SAX or DOM?

Therefore I created This Package.

## Package "xml.api/pull"

This Go Package defines the XML-Pull-API of the StAX-Pattern.

```
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

func (nt NodeType) String() string

type XmlReader interface {
    // Reads The Next Element. It returns true if an Element was Read
    // and false, for example if The Document has ended
    Read() bool

    // Gets The current Elements NodeType
    NodeType() NodeType

    // Gets The current Elements name, wich is an element-name or an attribute-name
    Name() string

    // Gets the current Elements value, wich is Text or an attribute-value
    Value() string

    // Gets the error that caused the Read()-Method to return false, if any
    GetError() error
}
```

## Package "xml.api/streamxmlp"

This Go Package implements a Basic but Fast (I hope) XML-Parser.

```
type PullXmlReader struct {
    // contains filtered or unexported fields
}
    This is a pull.XmlReader Implementation

func NewPullParser(src io.RuneReader) *PullXmlReader
    Creates a XmlReader From the given RuneReader.

func NewPullParserFromReader(src io.Reader) *PullXmlReader
    Creates a XmlReader From the given Reader. This equals
    NewPullParser(bufio.NewReader(src)).

func (xr *PullXmlReader) GetError() error

func (xr *PullXmlReader) Name() string

func (xr *PullXmlReader) NodeType() pull.NodeType

func (xr *PullXmlReader) Read() bool

func (xr *PullXmlReader) Value() string
```

## License

MIT-License

Copyright (c) 2012 Simon Schmidt

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
