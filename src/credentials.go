package src

import (
	"encoding/json"
	"os"
)
type Credential struct {
	User string `json:"user"`
	Passwd string `json:"passwd"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type Credentials struct {
	Users  []Credential `json:"users"`
}

func ReadJSONUsers (jsonPath string) Credentials {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		panic("Error reading doc "+ err.Error())
	}
	local := Credentials{}
	err = json.Unmarshal(data, &local)
	if err != nil {
		panic("Error reading doc "+ err.Error())
	}
	return local
}