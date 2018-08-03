package model

import "encoding/json"

type Profile struct {
	Name          string
	Age           int
	Height        string
	Income        string
	Status        string
	Edu           string
	Occupation    string
	Constellation string
	Hukou         string
	Place         string
}
type Item struct {
	URL      string
	ID       string
	TYPE     string
	Playload interface{}
}

func MarshJson(data interface{}) (Profile, error) {
	var profile Profile
	middle_result, err := json.Marshal(data)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(middle_result, &profile)
	return profile, err
}