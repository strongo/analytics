package analytics

var _ Message = (*page)(nil)
var _ Pageview = (*page)(nil)

type Pageview interface {
	Message
	Host() string
	Path() string
	Title() string
	SetTitle(title string) Pageview
}

func NewPageview(host, path string) Pageview {
	return &page{host: host, path: path}
}

type page struct {
	message
	host  string `key:"ph"`
	path  string `key:"pp"`
	title string `key:"pt"`
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

func (v *page) Validate() error {
	if err := v.message.Validate(); err != nil {
		return err
	}
	return nil
}
