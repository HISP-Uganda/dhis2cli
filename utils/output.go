package utils

import (
	"encoding/json"
)

func PrintResponse(responseMap map[string]any, pretty bool) (string, error) {
	if pretty {
		prettyJSON, err := json.MarshalIndent(responseMap, "", "  ")
		if err != nil {
			return "", err
		}
		return string(prettyJSON), nil
	} else {
		retJson, err := json.Marshal(responseMap)
		if err != nil {
			return "", err
		}
		return string(retJson), nil
	}
}
