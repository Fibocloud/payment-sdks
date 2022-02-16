package utils

import "encoding/json"

// JSONMarshal Interface to json
func JSONMarshal(input interface{}) (result []byte, err error) {
	result, err = json.Marshal(input)
	if err != nil {
		return
	}
	return
}

// JSONUnmarshal Json to interface
func JSONUnmarshal(input []byte) (result interface{}, err error) {
	err = json.Unmarshal(input, &result)
	if err != nil {
		return
	}
	return
}
