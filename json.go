package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//PrintJSON func
func PrintJSON(data interface{}) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

//PrintPrettyJSON func
func PrintPrettyJSON(data interface{}) error {

	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

//PrintPrettyJSON func
func PrintPrettyJSONFromStruct(data interface{}) error {

	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

//PrintPrettyJSON func
func PrintPrettyJSONFromJSON(data []byte) error {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(prettyJSON.Bytes()))
	return nil
}
