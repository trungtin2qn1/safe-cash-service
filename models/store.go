package models

//Store ...
type Store struct {
	ID   string `json:"id,omitempty" form:"id,omitempty"`
	Name string `json:"name,omitempty" form:"name,omitempty"`
}
