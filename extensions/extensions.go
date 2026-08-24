package ext

// Extensions is the generic extension map for Feeds and Items. The first map
// key is a canonical namespace key, a publisher-declared prefix, or a namespace
// URI fallback when no safe prefix is available. The second key is the element
// name (e.g., author).
type Extensions map[string]map[string][]Extension

// Extension represents a generic XML element retained from an extension tree.
// Nested unnamespaced children have an empty NamespaceURI.
type Extension struct {
	NamespaceURI string                 `json:"namespace_uri"`
	Name         string                 `json:"name"`
	Value        string                 `json:"value"`
	Attrs        map[string]string      `json:"attrs"`
	Children     map[string][]Extension `json:"children"`
}

func parseTextExtension(name string, extensions map[string][]Extension) (value string) {
	if extensions == nil {
		return
	}

	matches, ok := extensions[name]
	if !ok || len(matches) == 0 {
		return
	}

	match := matches[0]
	return match.Value
}

func parseTextArrayExtension(name string, extensions map[string][]Extension) (values []string) {
	if extensions == nil {
		return
	}

	matches, ok := extensions[name]
	if !ok || len(matches) == 0 {
		return
	}

	values = []string{}
	for _, m := range matches {
		values = append(values, m.Value)
	}
	return
}
