Quickly find an element in XML DOM and its parent elements.

## About
This is my first GoLang program.  I created it because I needed to quickly find an element and its parent elements in a large XML file and I didn't want to have to
sift through thousands of lines of the file and attempt to determine how the element was nested in order to create the structs I needed in another Go project.

## Usage
```
go run main.go -file=afile.xml -find=anelement
```

## Example
Let's say you have an XML file like this:
```
<?xml version="1.0"?>
<catalog>
   <book id="bk101">
      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>
      <publish_date>2000-10-01</publish_date>
      <description>An in-depth look at creating applications
      with XML.</description>
   </book>
</catalog>
```

Let's find how genre is nested:
```
go run main.go -file=examples/books.xml -find=genre
```

Will output:
```
[*]Opening file: examples/books.xml
[*]Successfully opened: examples/books.xml
[*]The hierarchy for your element looks like:
catalog
book
genre
```

So now I would create my structs in my other project to Unmarshal:
```
type Catalog struct {
  book []Book
  ...
}

type Book struct {
  genre []Genre
  ...
}

type Genre struct {
  value string
  ...
}
```

## TODO
Mostly for me to learn...
- Provide the built Go struct
- Make code more idomatic & efficient
- Create functions for some of the things
- Make it a package ?
- tests (unit/coverage)
