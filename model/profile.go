package model

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
