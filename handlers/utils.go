package handlers

import "encoding/json"

func parseBody(j json.RawMessage) (map[string]any, error) {
	var body map[string]any
	if err := json.Unmarshal(j, &body); err != nil {
		return nil, err
	}
	return body, nil
}
