package test

import (
	"fmt"
	"github.com/beevik/etree"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试xml解析
*/
func TestParseXml(t *testing.T) {
	document := etree.NewDocument()
	if err := document.ReadFromFile("./xml_demo.xml"); err != nil {
		panic(err)
	}

	root := document.SelectElement("bookstore")
	fmt.Println("ROOT element:", root.Tag)

	for _, book := range root.SelectElements("book") {
		fmt.Println("CHILD element:", book.Tag)
		if title := book.SelectElement("title"); title != nil {
			lang := title.SelectAttrValue("lang", "unknown")
			fmt.Printf("  TITLE: %s (%s)\n", title.Text(), lang)
		}
		for _, attr := range book.Attr {
			fmt.Printf("  ATTR: %s=%s\n", attr.Key, attr.Value)
		}
	}

}
