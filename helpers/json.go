package helpers

import "encoding/json"

func ToByteSlice(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return b
}