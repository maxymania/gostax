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
