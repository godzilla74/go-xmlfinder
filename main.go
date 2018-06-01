package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/subchen/go-xmldom"
)

func main() {

	filePtr := flag.String("file", "something.xml", "file location")
	findPtr := flag.String("find", "", "element to find in XML")

	flag.Parse()

	fmt.Println("[*]Opening file:", *filePtr)

	b, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		fmt.Println("[!]Error: ", err)
	} else {
		fmt.Println("[*]Successfully opened:", *filePtr)
	}

	// convert bytes to readable string
	xml := string(b)

	doc := xmldom.Must(xmldom.ParseXML(xml))
	root := doc.Root
	elem := root.FindOneByName(*findPtr)
	parent := elem.Parent

	// stores the XML DOM strings
	var elems []string

	// fmt.Println(rootSlice[0])

	elems = append(elems, elem.XMLPretty())

	for elem.Parent != root {
		parent = elem.Parent
		// slice tricks -- takes the current slice and prunes each child off one by one
		// from what I can tell we only want to do this with 'attribute' tags
		if strings.Contains(parent.XMLPretty(), "attributes") {
			for i := len(parent.Children) - 1; i >= 0; i-- {
				parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			}
		}

		// trim the attributes off the element
		for i := len(parent.Attributes) - 1; i >= 0; i-- {
			parent.Attributes = append(parent.Attributes[:i], parent.Attributes[i+1:]...)
		}

		// append the element to the slice
		elems = append(elems, parent.XMLPretty())
		// the elem should now equal the found parent
		elem = parent
	}

	// take the root and append it to the slice
	rootSlice := strings.Split(root.XMLPretty(), "\n")
	elems = append(elems, rootSlice[0])

	// reverse the slice
	for i, j := 0, len(elems)-1; i < j; i, j = i+1, j-1 {
		elems[i], elems[j] = elems[j], elems[i]
	}

	// iterate over slice to create structs
	var trimmedSlice []string
	for _, e := range elems {
		r := regexp.MustCompile(`\w+`)
		match := r.FindString(e)
		trimmedSlice = append(trimmedSlice, match)
	}

	// The final slice
	finalSlice := strings.Join(trimmedSlice, "\n")
	fmt.Printf("[*]The hierarchy for your element looks like:\n%v\n", finalSlice)
}
