package xml

//From
//https://github.com/dustin/go-wikiparse

import (
	"bytes"
	"regexp"
)

var nowikiRE, commentRE, languageRE *regexp.Regexp
var ignoreList []string

func init() {
	nowikiRE = regexp.MustCompile(`(<nowiki>.*</nowiki>)`)
	languageRE = regexp.MustCompile(`(\[\[[a-z][a-z]:.*\]\])`)
	commentRE = regexp.MustCompile(`(<!--.*-->)`)
	//ignoreList = []string{"File:", "wikt:", "Wiktionary:", "als:", "simple:", "rue:", "ckb:", "mwl:", "zh-yue:", "pdc:", "mwl:", "arz:", "pnb:", "sah:"}
}

// Find all the links from within an article body.
func FindLinks(text []byte) []string {
	// http://www.mediawiki.org/wiki/Help:Links
	//TODO add unit tests <-----
	//TODO remove xml parser as a separate branch. Is this really needed or was it the parser all along
	rv := make([]string, 0)
	cleanedText := cleanText(text)

	for stop := false; stop == false; {
		startTag := []byte("[[")
		startIndex := bytes.Index(cleanedText, startTag)
		if startIndex == -1 {
			break
		}

		endTag := []byte("]]")
		endIndex := startIndex + len(startTag) +
			bytes.Index(cleanedText[startIndex+len(startTag):], endTag)

		if startIndex != -1 && endIndex != -1 && startIndex+len(startTag) < endIndex {
			linkText := cleanedText[startIndex+len(startTag) : endIndex]
			splitIndex := bytes.IndexAny(linkText, "#|")
			if splitIndex != -1 {
				linkText = linkText[:splitIndex]
			}
			//TODO: Add in ignore list support
			rv = append(rv, string(bytes.ToLower(linkText)))
			cleanedText = cleanedText[endIndex+len(endTag):]
		} else {
			stop = true
		}
	}

	return rv
}

func cleanText(text []byte) []byte {
	return text
	//return languageRE.ReplaceAllString(nowikiRE.ReplaceAllString(commentRE.ReplaceAllString(text, ""), ""), "")
}
