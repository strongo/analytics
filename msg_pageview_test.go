package analytics

import (
	"reflect"
	"testing"
)

const expectedHost = "unit-test"
const expectedPathWithoutLeadingSlash = "test/path"
const expectedPathWithLeadingSlash = "/" + expectedPathWithoutLeadingSlash

func Test_page_SetURL(t *testing.T) {
	const expectedUrl = "tg://TestBot/some/path"

	pageView := NewPageview(expectedHost, expectedPathWithLeadingSlash).
		SetURL(expectedUrl)

	if url := pageView.URL(); url != expectedUrl {
		t.Errorf("expected %s, got %s", expectedUrl, url)
	}
}

func Test_page_SetTitle(t *testing.T) {
	const expectedTitle = "Unit Test"

	pageView := NewPageview(expectedHost, expectedPathWithLeadingSlash).
		SetTitle(expectedTitle).
		SetUserAgent("Mozilla/5.0")

	if title := pageView.Title(); title != expectedTitle {
		t.Errorf("expected %s, got %s", expectedTitle, title)
	}

	if host := pageView.Host(); host != expectedHost {
		t.Errorf("expected %s, got %s", expectedHost, host)
	}

	if path := pageView.Path(); path != expectedPathWithLeadingSlash {
		t.Errorf("expected %s, got %s", expectedPathWithLeadingSlash, path)
	}

	if err := pageView.Validate(); err != nil {
		t.Errorf("expected valid pageview, got %v", err)
	}

	invalidPage := &page{message: message{event: ""}}
	if err := invalidPage.Validate(); err == nil {
		t.Error("expected validation error for empty event")
	}
}

func TestNewPageview(t *testing.T) {
	type args struct {
		host string
		path string
	}
	tests := []struct {
		name string
		args args
		want Pageview
	}{
		{
			name: "path_with_leading_slash",
			args: args{
				host: expectedHost,
				path: expectedPathWithLeadingSlash,
			},
			want: &page{
				host: expectedHost,
				path: expectedPathWithLeadingSlash,
				message: message{
					event:      "$pageview",
					properties: make(Properties, defaultPropertiesCapacity),
				},
			},
		},
		{
			name: "path_without_leading_slash",
			args: args{
				host: expectedHost,
				path: expectedPathWithoutLeadingSlash,
			},
			want: &page{
				host: expectedHost,
				path: expectedPathWithLeadingSlash,
				message: message{
					event:      "$pageview",
					properties: make(Properties, defaultPropertiesCapacity),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPageview(tt.args.host, tt.args.path)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPageview() =\n\t%+v \nwant\n\t%+v", got, tt.want)
			}
		})
	}
}
