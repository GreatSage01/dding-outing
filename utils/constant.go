package utils

import (
	"bytes"
	"encoding/json"
)

//struct转json,并格式化
func Struct2json(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	var out bytes.Buffer
	json.Indent(&out, bs, "", "\t")
	return out.String()
}
