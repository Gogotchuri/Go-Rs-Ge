package responses

type CustomerWithUniqueID struct {
	UniqueID int    `xml:"Body>get_un_id_from_tinResponse>get_un_id_from_tinResult"`
	Name     string `xml:"Body>get_un_id_from_tinResponse>name"`
}

type GetUniqueIDFromTINResponse struct {
	XMLResponse
	CustomerWithUniqueID
}
