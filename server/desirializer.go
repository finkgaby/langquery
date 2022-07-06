package server

import (
	"encoding/json"
	"langquery/commons"
)

func Deserialize(inData []byte) []byte {
	var text string
	err := json.Unmarshal(inData, &text)
	commons.CheckErr(err)
	text = "id='andy'"
	outData, _ := json.Marshal(text)
	return outData
}
