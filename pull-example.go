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
