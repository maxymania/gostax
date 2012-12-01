package pull

type XmlReader interface{
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
