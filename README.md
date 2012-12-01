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

## Package "xml.api/streamxmlp"

This Go Package implements a Basic but Fast (I hope) XML-Parser.

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
