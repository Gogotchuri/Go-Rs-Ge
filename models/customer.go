package models

type Customer struct {
	TIN  string `xml:"Body>get_tin_from_un_idResponse>get_tin_from_un_idResult"`
	Name string `xml:"Body>get_tin_from_un_idResponse>name"`
}

type CustomerM struct {
	TIN      string
	Name     string
	UniqueID int
}
