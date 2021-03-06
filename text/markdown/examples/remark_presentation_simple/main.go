package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/grokify/simplego/text/markdown"
)

func main() {
	slides := markdown.PresentationData{
		Slides: []markdown.RemarkSlideData{
			{
				Layout:   "middle, center, inverse",
				Class:    "false",
				Markdown: "# Test Slide\n\nTest Remark Slide",
			},
			{
				Markdown: "# Test Slide\n\nTest Remark Slide",
			},
		},
	}
	html := markdown.RemarkHTML(slides)
	fmt.Println(html)

	err := ioutil.WriteFile("test_slides.html", []byte(html), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
