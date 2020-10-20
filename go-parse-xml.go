package main

import (
	"encoding/xml"
	"log"
)

type Entry struct {
	IpString         string `xml:"ip"`
	Vsys             string `xml:"vsys"`
	Type             string `xml:"type"`
	User             string `xml:"user"`
	IdeTimeoutString string `xml:"idle_timeout"`
	TimeoutString    string `xml:"timeout"`
}

type Result struct {
	Entries []Entry `xml:"entry"`
}

type PAXmlReseponse struct {
	XMLName xml.Name `xml:response`
	Status  string   `xml:"status,attr"`
	Result  Result   `xml:"result"`
}

func main() {
	testxml := "<response status=\"success\"><result><entry><ip>172.16.101.99</ip><vsys>vsys1</vsys><type>XMLAPI</type><user>pauser</user><idle_timeout>8747</idle_timeout><timeout>8747</timeout></entry>\n<entry><ip>172.16.101.101</ip><vsys>vsys1</vsys><type>XMLAPI</type><user>pauser</user><idle_timeout>10797</idle_timeout><timeout>10797</timeout></entry>\n<count>2</count>\n</result></response>"

	v := PAXmlReseponse{}

	err := xml.Unmarshal([]byte(testxml), &v)
	if err != nil {
		log.Printf("error %v\n", v.XMLName)
		return
	}

	log.Printf("status = %s\n", v.Status)
	if v.Status != "success" {
		return
	}

	log.Printf("entries: %d\n", len(v.Result.Entries))
	for idx, entry := range v.Result.Entries {
		log.Printf("  [%d] ip=%s\n", idx, entry.IpString)
		log.Printf("  [%d] user=%s\n", idx, entry.User)
		log.Printf("  [%d] timeout=%s\n", idx, entry.TimeoutString)
	}
}
