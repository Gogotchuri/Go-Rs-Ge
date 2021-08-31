package requests

type Operation string

type BaseRequest struct {
	ServiceURL string    `xml:"-"`
	Operation  Operation `xml:"-"`
	XMLns      string    `xml:"xmlns,attr"`
}

func (br *BaseRequest) GetEndpointURL() string {
	return br.ServiceURL + string(br.Operation)
}

func (br *BaseRequest) Initialize(op Operation) {
	br.XMLns = "http://tempuri.org/"
	br.Operation = op
	br.ServiceURL = ServiceAddress
}

//BaseAuthorizedRequest contains base properties to be authorized by rs.ge
type BaseAuthorizedRequest struct {
	UserID          int    `xml:"user_id"`
	UniqueID        int    `xml:"un_id"`
	ServiceUser     string `xml:"su"`
	ServicePassword string `xml:"sp"`
}