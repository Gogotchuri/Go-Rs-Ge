package soap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

//GetXMLBytes does heavy lifing here
func GetXMLBytes(req XMLRequest) (*bytes.Buffer, error) {
	marshalledXML, err := xml.MarshalIndent(req, " ", "  ")
	if err != nil {
		fmt.Println("Marshalling error", err.Error())
		return nil, err
	}
	// Using the var getTemplate to construct request
	template, err := template.New("XMLRequest").Parse(RsGeXMLTemplate)
	if err != nil {
		fmt.Printf("error while parsing template %s ", err.Error())
		return nil, err
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, struct {
		OperationBody string
	}{
		OperationBody: string(marshalledXML),
	})
	if err != nil {
		fmt.Printf("template.Execute error %s ", err.Error())
		return nil, err
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Printf("encoder.Encode error %s ", err.Error())
		return nil, err
	}
	return bytes.NewBuffer(doc.Bytes()), nil
}

func GenerateSOAPRequest(req XMLRequest) (*http.Request, error) {
	buffer, err := req.GetXML()
	if err != nil {
		fmt.Printf("Error getting xml from request. %s \n", err.Error())
		return nil, err
	}
	r, err := http.NewRequest(http.MethodPost, req.GetEndpointURL(), buffer)
	if err != nil {
		fmt.Printf("Error making a request. %s ", err.Error())
		return nil, err
	}

	r.Header.Add("Content-Type", "text/xml; charset=UTF-8")
	r.Header.Add("accept", "*/*")

	return r, nil
}

//SoapCall make a soap call given formed http request and resp to unmarshall into
func SoapCall(req *http.Request, dest interface{}) (interface{}, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = xml.Unmarshal(body, dest)

	if err != nil {
		return nil, err
	}

	return dest, nil
}
