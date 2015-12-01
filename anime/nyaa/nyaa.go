package nyaa

import (
	"encoding/xml"
	"io"

	"github.com/Wessie/unhtml"
)

type Item struct {
	Title       string `unhtml:"title"`
	Category    string `unhtml:"category"`
	Link        string `unhtml:"link"`
	Guid        string `unhtml:"guid"`
	Description string `unhtml:"description"`
	Date        string `unhtml:"pubDate"`
}

func Parse(r io.Reader) ([]Item, error) {
	for i, v := range xml.HTMLAutoClose {
		if v == "link" {
			xml.HTMLAutoClose[i] = "bogus"
			break
		}
	}

	d, err := unhtml.NewDecoder(r)
	if err != nil {
		return nil, err
	}

	var result []Item
	var startPath = "rss/channel/item"

	if err := d.UnmarshalRelative(startPath, &result); err != nil {
		return nil, err
	}

	return result, nil
}
