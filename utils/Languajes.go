package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Language(languaje string, messageCode string) string {
	if len(languaje) <= 0 {
		languaje = os.Getenv("DEFAULT_LANGUAGE")
	}
	path := os.Getenv("PATH_LANGUAGE") + os.Getenv(languaje)
	resp, err := ioutil.ReadFile(path)
	var generic map[string]interface{}
	err = json.Unmarshal([]byte(resp), &generic)
	if err != nil {
		path := os.Getenv("PATH_LANGUAGE") + os.Getenv("DEFAULT_LANGUAGE")
		byteValue, _ := ioutil.ReadFile(path)
		var generic map[string]interface{}
		json.Unmarshal([]byte(byteValue), &generic)
		return generic[messageCode].(string)
	}
	return generic[messageCode].(string)
}
