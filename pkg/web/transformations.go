package web

import "encoding/json"

func UnmarshalJSON(body []byte, object interface{}) error {
	return json.Unmarshal(body, object)
}
