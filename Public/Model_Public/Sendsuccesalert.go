package Model_Public

import (
	"encoding/json"
)

func Sendsuccessalert(message string) []byte{

	data1,err:=json.Marshal(message)
	CheckErr(err)
	return data1
}
