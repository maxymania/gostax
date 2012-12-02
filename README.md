Go Xml Apis
===========

Go Streaming API for XML (and More)

## Why? Go has an Xml De-/Serializer!

Go has a package called "encoding/xml". It Supports the serialization and deserialization of XML as well as Streaming XML Parsing.
When I Created this Package, I thought that the "encoding/xml" package is just about Marhsaling and Unmarshaling of Structures (Oops!).

I wanted a Generic inteface to XML such as StAX, SAX or DOM?

Therefore I created This Package and I hope that it will be useful.

## How I Can get it?

Download The Package or git-clone it:
```
git clone https://github.com/maxymania/gostax.git
```

## Examples

pull-example.go:
```go
package main

import "io"
// import "bytes"
import "strings"
import "fmt"
import "xml.api/streamxmlp"
import "xml.api/pull"

const examplexml = `<?xml version="1.0" encoding="UTF-8" ?>
<note> 
	<to>Tove</to> 
	<from>Jani</from> 
	<heading>Reminder</heading> 
	<body>Don't forget me this weekend!</body> 
</note>`


func main(){
	// The RuneReader to Read the XML From
	var r io.RuneReader
	
	// The XmlReader Object
	var xr pull.XmlReader
	r = strings.NewReader(examplexml)
	
	// Creating a new XmlReader (Implementation : streamxmlp)
	xr = streamxmlp.NewPullParser(r)
	
	for xr.Read() {
		fmt.Println(xr.NodeType(),",",xr.Name(),",",xr.Value())
	}
	fmt.Println("EOP")
}
```

push-example.go:
```go
package main

import "bytes"
import "fmt"
import "xml.api/push"
import "xml.api/sxmlwriter"


func main(){
	// The Write to write the XML to
	buf := new(bytes.Buffer)
	
	// The XmlWriter Object
	var xw push.XmlWriter
	
	// Creating a new XmlWriter (Implementation : sxmlwriter)
	xw = sxmlwriter.NewXmlWriter(buf)
	
	// Lets Write some XML
	xw.StartElement("test")
	xw.Attribute("hallo","welt")
	xw.Text("hallo welt")
	xw.EndElement()
	
	// Lets Show the result
	fmt.Println(buf)
}
```


## Go Docs Can be found here

http://maxymania.github.com/gostax/
Or in the wiki

## License

This Software is licensed under the MIT-License.