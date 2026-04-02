package schemaorg

import "encoding/json"

// ImageObject is a subset of https://schema.org/ImageObject (Person.image as object).
type ImageObject struct {
	URL         string `json:"url,omitempty"`
	ContentURL  string `json:"contentUrl,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Caption     string `json:"caption,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type ImageObject.
func (i ImageObject) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "ImageObject"}
	if i.URL != "" {
		m["url"] = i.URL
	}
	if i.ContentURL != "" {
		m["contentUrl"] = i.ContentURL
	}
	if i.Name != "" {
		m["name"] = i.Name
	}
	if i.Description != "" {
		m["description"] = i.Description
	}
	if i.Caption != "" {
		m["caption"] = i.Caption
	}
	return json.Marshal(m)
}
