package factoryBrowser

import "github.com/helmutkemper/webassembly/browser/html"

func NewTagImg() (ref *html.TagImg) {
	ref = &html.TagImg{}
	ref.Init()

	return ref
}
