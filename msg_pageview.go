package analytics

import (
	"strings"
)

//const PageViewEventPrefix = "page:"

var _ Message = (*page)(nil)
var _ Pageview = (*page)(nil)

type Pageview interface {
	Message
	Host() string
	Path() string
	Title() string
	SetTitle(title string) Pageview
	URL() string
	SetURL(url string) Pageview
	SetUserAgent(userAgent string) Pageview
}

func NewPageview(host, path string) Pageview {
	m := newMessage("$pageview")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return &page{message: m, host: host, path: path}
}

type page struct {
	message
	host      string `key:"ph"`
	path      string `key:"pp"`
	url       string `key:"pu"`
	title     string `key:"pt"`
	userAgent string `key:"ua"`
}

func (v *page) Host() string {
	return v.host
}

func (v *page) Path() string {
	return v.path
}

func (v *page) URL() string {
	return v.url
}

func (v *page) Title() string {
	return v.title
}

func (v *page) SetTitle(title string) Pageview {
	v.title = title
	return v
}

func (v *page) SetUserAgent(userAgent string) Pageview {
	v.userAgent = userAgent
	return v
}

func (v *page) SetURL(url string) Pageview {
	v.url = url
	return v
}

func (v *page) Validate() error {
	if err := v.message.Validate(); err != nil {
		return err
	}
	return nil
}
