package schemaorg

import "encoding/json"

// InteractionCounter is a subset of https://schema.org/InteractionCounter (Person.interactionStatistic).
type InteractionCounter struct {
	InteractionType      string `json:"interactionType,omitempty"`
	UserInteractionCount int    `json:"userInteractionCount,omitempty"`
}

// MarshalJSON emits a JSON-LD node with @type InteractionCounter.
func (i InteractionCounter) MarshalJSON() ([]byte, error) {
	m := map[string]any{"@type": "InteractionCounter"}
	if i.InteractionType != "" {
		m["interactionType"] = i.InteractionType
	}
	if i.UserInteractionCount != 0 {
		m["userInteractionCount"] = i.UserInteractionCount
	}
	return json.Marshal(m)
}
