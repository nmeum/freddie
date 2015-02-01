// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package atom

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

var atomFeed Feed

func TestFeed(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/feed.atom")
	if err != nil {
		t.Fatal(err)
	}

	if err := xml.Unmarshal(data, &atomFeed); err != nil {
		t.Fatal(err)
	}

	if atomFeed.Title != "Some Title" {
		t.Fatalf("Expected %q - got %q", "Some Title", atomFeed.Title)
	}

	testLink := Link{
		Type: "text/html",
		Href: "http://example.org/feed.atom",
		Rel:  "self",
	}

	if atomFeed.Links[0] != testLink {
		t.Fatalf("Expected %q - got %q", testLink, atomFeed.Links[0])
	}
}

func TestEntry(t *testing.T) {
	entry := atomFeed.Entries[0]

	if entry.Published != "2013-10-11T23:56:00Z" {
		t.Fatalf("Expected %q - got %q", "2013-10-11T23:56:00Z", entry.Published)
	}

	if entry.Title != "Test Post" {
		t.Fatalf("Expected %q - got %q", "Test Post", entry.Title)
	}

	testLink := Link{
		Type: "text/html",
		Href: "http://example.org/posts/test.html",
		Rel:  "alternate",
	}

	if entry.Links[0] != testLink {
		t.Fatalf("Expected %q - got %q", testLink, entry.Links[0])
	}
}

func TestFindLink(t *testing.T) {
	links := []Link{
		{"text/html", "http://example.com/my_link", ""},
		{"audio/ogg", "http://example.com/my_foo", "enclosure"},
	}

	link := findLink(links)
	if link != links[0] {
		t.Fatalf("Expected %q - got %q", link, links[0])
	}
}

func TestFindAttachment(t *testing.T) {
	links := []Link{
		{"text/html", "http://example.org/foo", "alternate"},
		{"image/png", "http://example.org/bar", "enclosure"},
		{"text/html", "http://example.org/baz", ""},
	}

	link := findAttachment(links)
	if link != links[1] {
		t.Fatalf("Expected %q - got %q", link, links[1])
	}
}
