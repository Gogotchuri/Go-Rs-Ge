package models

type Customer struct {
	TIN int `xml:"Body>get_tin_from_un_idResponse>get_tin_from_un_idResult"`
	Name string `xml:"Body>get_tin_from_un_idResponse>name"`
}