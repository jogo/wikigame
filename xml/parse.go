package xml

import (
	"bufio"
	"bytes"
	"io"
)

// A wiki page.
type Page struct {
	Title []byte `title`
	Text  []byte `string`
}

// That which emits wiki pages.
type Parser interface {
	// Get the next page from the parser
	Next() (*Page, error)
}

type streamParser struct {
	x *bufio.Reader
}

// Get a wikipedia dump parser reading from the given reader.
func NewParser(r io.Reader) (Parser, error) {
	d := bufio.NewReader(r)

	return &streamParser{
		x: d,
	}, nil
}

func (p *streamParser) Next() (rv *Page, err error) {
	rv = new(Page)

	page := p.getPage()
	rv.Title = bytes.ToLower(p.getTag(page, []byte("\n    <title>"), []byte("</title>")))
	rv.Text = page
	//rv.Text = string(p.getTag(page, []byte("\n      <text "), []byte("</text>")))
	return
}

// <page> (<redirect title)? (text) <page/>
func (p *streamParser) getPage() []byte {
	buffer := make([]byte, 0)
	inPage := false

	for true {
		text, _ := p.x.ReadBytes(byte('>'))
		//use whitespace to reduce false hits
		startPage := []byte("\n  <page>")
		startIndex := bytes.Index(text, startPage)
		endPage := []byte("</page>")
		endIndex := bytes.Index(text, endPage)

		if startIndex != -1 {
			inPage = true
			buffer = text[startIndex+len(startPage):]
		} else if inPage && endIndex != -1 {
			inPage = false
			buffer = append(buffer, text[:endIndex]...)

			// don't index redirect pages
			if bytes.Index(buffer, []byte("<redirect ")) == -1 {
				return buffer
			}
			buffer = make([]byte, 0)
		} else if inPage == true {
			buffer = append(buffer, text...)
		}

	}

	return []byte("BAD")
}

func (p *streamParser) getTag(text, startTag, endTag []byte) []byte {
	startIndex := bytes.Index(text, startTag)
	endIndex := bytes.Index(text, endTag)
	if startIndex != -1 && endIndex != -1 {
		return text[startIndex+len(startTag) : endIndex]
	}
	return []byte("SOMETHING WENT WRONG, TURN INTO ERROR")

}
