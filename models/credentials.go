package models

type DeclarationStatus int

type CredentialCheck struct {
	Result bool `xml:"chekResult"`
	UserID int  `xml:"user_id"`
	Sua    int  `xml:"sua"`
}
