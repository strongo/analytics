package analytics

import (
	"net/url"
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
	SetUserAgent(userAgent string) Pageview
}

func NewPageview(host, path string) Pageview {
	m := newMessage("$pageview")
	m.properties.Set("$host", host)
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	m.properties.Set("$pathname", path)
	if host == "telegram" && strings.HasPrefix(path, "/bot/") {
		m.properties.Set("$current_url", "tg://"+path[len("/bot/"):])
	}
	return &page{message: m, host: host, path: path}
}

type page struct {
	message
	host      string   `key:"ph"`
	path      string   `key:"pp"`
	url       *url.URL `key:"url"`
	title     string   `key:"pt"`
	userAgent string   `key:"ua"`
}

func (v *page) Host() string {
	return v.host
}
func (v *page) Path() string {
	return v.path
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

func (v *page) Validate() error {
	if err := v.message.Validate(); err != nil {
		return err
	}
	return nil
}
