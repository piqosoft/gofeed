package shared

import (
	"strings"
	"testing"

	xpp "github.com/mmcdole/goxpp/v2"

	ext "github.com/piqosoft/gofeed/extensions"
)

// parserOn returns a NewXMLParser positioned on the first StartTag with the
// given local name.
func parserOn(t *testing.T, doc, name string) *xpp.Parser {
	t.Helper()
	p := NewXMLParser(strings.NewReader(doc))
	for {
		tok, err := p.NextToken()
		if err != nil {
			t.Fatalf("positioning on <%s>: %v", name, err)
		}
		if tok == xpp.StartTag && p.Name() == name {
			return p
		}
		if tok == xpp.EndDocument {
			t.Fatalf("no <%s> element found", name)
		}
	}
}

func TestIsExtension(t *testing.T) {
	doc := `<rss xmlns:itunes="http://www.itunes.com/DTDs/PodCast-1.0.dtd" xmlns:content="http://purl.org/rss/1.0/modules/content/">
		<channel>
			<title>t</title>
			<itunes:author>a</itunes:author>
			<content:encoded>c</content:encoded>
		</channel>
	</rss>`

	cases := []struct {
		element string
		want    bool
	}{
		{"title", false},  // no namespace
		{"author", true},  // iTunes namespace
		{"encoded", true}, // content namespace
	}
	for _, c := range cases {
		p := parserOn(t, doc, c.element)
		if got := IsExtension(p, ""); got != c.want {
			t.Errorf("IsExtension(<%s>) = %v, want %v", c.element, got, c.want)
		}
	}
}

func TestIsExtensionUsesResolvedNamespace(t *testing.T) {
	doc := `<rss xmlns:a="http://www.w3.org/2005/Atom"><channel><a:link/></channel></rss>`
	p := parserOn(t, doc, "link")

	if IsExtension(p, Atom10Namespace) {
		t.Error("Atom namespace should be native when the Atom parser supplies it")
	}
	if !IsExtension(p, "") {
		t.Error("Atom namespace should be an extension when the RSS parser supplies only its native namespace")
	}
}

func TestParseExtension(t *testing.T) {
	doc := `<rss xmlns:itunes="http://www.itunes.com/DTDs/PodCast-1.0.dtd">
		<channel>
			<itunes:owner>
				<itunes:name>Alice</itunes:name>
				<itunes:email>a@example.org</itunes:email>
			</itunes:owner>
			<itunes:category text="Tech"/>
			<itunes:category text="News"/>
		</channel>
	</rss>`

	p := NewXMLParser(strings.NewReader(doc))
	extensions := ext.Extensions{}
	for {
		tok, err := p.NextToken()
		if err != nil {
			t.Fatal(err)
		}
		if tok == xpp.EndDocument {
			break
		}
		if tok == xpp.StartTag && IsExtension(p, "") {
			extensions, err = ParseExtension(extensions, p)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	owner := extensions["itunes"]["owner"]
	if len(owner) != 1 {
		t.Fatalf("owner entries = %d, want 1", len(owner))
	}
	if got := owner[0].Children["name"][0].Value; got != "Alice" {
		t.Errorf("owner name = %q, want Alice", got)
	}
	if got := owner[0].Children["email"][0].Value; got != "a@example.org" {
		t.Errorf("owner email = %q, want a@example.org", got)
	}

	cats := extensions["itunes"]["category"]
	if len(cats) != 2 {
		t.Fatalf("category entries = %d, want 2 (same-name elements must append)", len(cats))
	}
	if cats[0].Attrs["text"] != "Tech" || cats[1].Attrs["text"] != "News" {
		t.Errorf("category attrs = %v, %v", cats[0].Attrs, cats[1].Attrs)
	}
	if cats[0].Name != "category" {
		t.Errorf("extension Name = %q, want category", cats[0].Name)
	}
}

func TestParseExtensionPreservesResolvedNamespaceOnEveryNode(t *testing.T) {
	const (
		podcastNamespace = "https://podcastindex.org/namespace/1.0"
		itunesNamespace  = "http://www.itunes.com/DTDs/PodCast-1.0.dtd"
	)
	doc := `<rss xmlns:p20="` + podcastNamespace + `" xmlns:itunes="` + itunesNamespace + `">
		<channel>
			<p20:liveItem status="pending">
				<itunes:season>2</itunes:season>
				<p20:season name="Bonus">3</p20:season>
				<title>Live now</title>
			</p20:liveItem>
		</channel>
	</rss>`

	p := parserOn(t, doc, "liveItem")
	extensions, err := ParseExtension(ext.Extensions{}, p)
	if err != nil {
		t.Fatal(err)
	}

	liveItems := extensions["podcast"]["liveItem"]
	if len(liveItems) != 1 {
		t.Fatalf("liveItem entries = %d, want 1", len(liveItems))
	}
	if got := liveItems[0].NamespaceURI; got != podcastNamespace {
		t.Errorf("liveItem namespace = %q, want %q", got, podcastNamespace)
	}
	seasons := liveItems[0].Children["season"]
	if len(seasons) != 2 {
		t.Fatalf("season entries = %d, want 2", len(seasons))
	}
	if got := seasons[0].NamespaceURI; got != itunesNamespace {
		t.Errorf("iTunes season namespace = %q, want %q", got, itunesNamespace)
	}
	if got := seasons[1].NamespaceURI; got != podcastNamespace {
		t.Errorf("Podcast season namespace = %q, want %q", got, podcastNamespace)
	}
	if got := liveItems[0].Children["title"][0].NamespaceURI; got != "" {
		t.Errorf("RSS title namespace = %q, want empty", got)
	}
}

func TestParseExtensionCanonicalizesPodcastNamespace(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		uri    string
	}{
		{name: "preferred URI literal prefix", prefix: "podcast", uri: PodcastNamespace},
		{name: "preferred URI alternate prefix", prefix: "p20", uri: PodcastNamespace},
		{name: "legacy URI literal prefix", prefix: "podcast", uri: PodcastLegacyNamespace},
		{name: "legacy URI alternate prefix", prefix: "legacy", uri: PodcastLegacyNamespace},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := `<rss xmlns:` + tt.prefix + `="` + tt.uri + `"><channel><` + tt.prefix + `:funding>help</` + tt.prefix + `:funding></channel></rss>`
			got, err := ParseExtension(ext.Extensions{}, parserOn(t, doc, "funding"))
			if err != nil {
				t.Fatal(err)
			}
			funding := got["podcast"]["funding"]
			if len(funding) != 1 || funding[0].NamespaceURI != tt.uri {
				t.Fatalf("podcast extension = %#v, want observed URI %q", funding, tt.uri)
			}
		})
	}

	t.Run("unrelated URI cannot claim literal prefix", func(t *testing.T) {
		const unrelated = "urn:not-podcast"
		doc := `<rss xmlns:podcast="` + unrelated + `"><channel><podcast:funding>nope</podcast:funding></channel></rss>`
		got, err := ParseExtension(ext.Extensions{}, parserOn(t, doc, "funding"))
		if err != nil {
			t.Fatal(err)
		}
		if _, claimed := got["podcast"]; claimed {
			t.Fatalf("unrelated extension claimed canonical podcast key: %#v", got)
		}
		if funding := got[unrelated]["funding"]; len(funding) != 1 || funding[0].NamespaceURI != unrelated {
			t.Fatalf("unrelated extension = %#v", got)
		}
	})
}

func TestPrefixForNamespace(t *testing.T) {
	const (
		podcastNamespace       = "https://podcastindex.org/namespace/1.0"
		legacyPodcastNamespace = "https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md"
	)
	doc := `<rss xmlns:pod="http://www.itunes.com/DTDs/PodCast-1.0.dtd" xmlns:custom="http://example.org/ns" xmlns:a="http://www.w3.org/2005/Atom" xmlns:legacy="http://purl.org/atom/ns#" xmlns:p20="` + podcastNamespace + `" xmlns:oldp20="` + legacyPodcastNamespace + `">
		<channel><pod:author>a</pod:author><a:link/><legacy:link/><p20:funding/><oldp20:funding/></channel>
	</rss>`
	p := parserOn(t, doc, "author")

	// A canonical namespace URI maps to its canonical prefix, not the
	// prefix the feed used.
	if got := PrefixForNamespace("http://www.itunes.com/DTDs/PodCast-1.0.dtd", p); got != "itunes" {
		t.Errorf("canonical = %q, want itunes", got)
	}
	// Whitespace-padded URIs map the same way.
	if got := PrefixForNamespace(" http://www.itunes.com/DTDs/PodCast-1.0.dtd ", p); got != "itunes" {
		t.Errorf("padded canonical = %q, want itunes", got)
	}
	// A feed-declared namespace maps to the feed's prefix.
	if got := PrefixForNamespace("http://example.org/ns", p); got != "custom" {
		t.Errorf("feed-declared = %q, want custom", got)
	}
	// An undeclared prefix comes through encoding/xml as the raw prefix in
	// Space; it maps to itself.
	if got := PrefixForNamespace("media", p); got != "media" {
		t.Errorf("undeclared = %q, want media", got)
	}
	atom := parserOn(t, doc, "link")
	if got := PrefixForNamespace(Atom10Namespace, atom); got != "atom" {
		t.Errorf("Atom canonical = %q, want atom", got)
	}
	if got := PrefixForNamespace(Atom03Namespace, atom); got != "atom03" {
		t.Errorf("Atom 0.3 canonical = %q, want atom03", got)
	}
	if got := PrefixForNamespace(podcastNamespace, parserOn(t, doc, "funding")); got != "podcast" {
		t.Errorf("Podcast canonical = %q, want podcast", got)
	}
	if got := PrefixForNamespace(legacyPodcastNamespace, parserOn(t, doc, "funding")); got != "podcast" {
		t.Errorf("legacy Podcast canonical = %q, want podcast", got)
	}
	wrongPodcast := parserOn(t, `<rss xmlns:podcast="urn:not-podcast"><channel><podcast:funding/></channel></rss>`, "funding")
	if got := PrefixForNamespace("urn:not-podcast", wrongPodcast); got == "podcast" {
		t.Errorf("unrelated namespace claimed reserved key %q", got)
	}
	defaultOnly := parserOn(t, `<root><value xmlns="urn:example"/></root>`, "value")
	if got := PrefixForNamespace("urn:example", defaultOnly); got != "urn:example" {
		t.Errorf("default namespace = %q, want urn:example", got)
	}
}

// NewXMLParser must be non-strict (real feeds carry bare ampersands) and must
// convert declared non-UTF-8 encodings.
func TestNewXMLParserLeniencyAndCharset(t *testing.T) {
	p := parserOn(t, `<root>Fish & Chips</root>`, "root")
	text, err := p.NextText()
	if err != nil {
		t.Fatalf("bare ampersand should tokenize in non-strict mode: %v", err)
	}
	if !strings.Contains(text, "&") {
		t.Errorf("text = %q, want the ampersand preserved", text)
	}

	// 0xE9 is é in ISO-8859-1 and invalid UTF-8; parsing it proves the
	// charset reader is wired up.
	latin1 := `<?xml version="1.0" encoding="ISO-8859-1"?><root>caf` + "\xe9" + `</root>`
	p2 := parserOn(t, latin1, "root")
	text, err = p2.NextText()
	if err != nil {
		t.Fatal(err)
	}
	if text != "café" {
		t.Errorf("text = %q, want café", text)
	}
}
