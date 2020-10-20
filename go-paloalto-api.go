package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func GetIPUserMapping(host string, key string) ([]Entry, error) {
	scheme := "https://"
	urlsuffix := "/api/?type=op"
	xpath := "<show><user><ip-user-mapping><all></all></ip-user-mapping></user></show>"

	url := fmt.Sprintf("%s%s%s&key=%s&cmd=%s", scheme, host, urlsuffix, key, xpath)

	// disable tls

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	v := PAXmlReseponse{}
	err = xml.Unmarshal([]byte(body), &v)
	if err != nil {
		log.Printf("error %v\n", v.XMLName)
		return nil, fmt.Errorf("failed to parse XML response")
	}

	if v.Status != "success" {
		return nil, fmt.Errorf("response status is error")
	}

	entries := v.Result.Entries

	return entries, nil
}

const (
	kLogInOutXMLTemplate = `<uid-message>
	<version>1.0</version>
	<type>update</type>
	<payload>
		<{{.Method}}>
			<entry name="{{.UserName}}" ip="{{.IPAddress}}" timeout="{{.Timeout}}"></entry>
		</{{.Method}}>
	</payload>
</uid-message>`
)

type LogInOutTemplateValue struct {
	Method    string
	UserName  string
	IPAddress string
	Timeout   int
}

func generateXMLTemplateString(liot LogInOutTemplateValue) (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("temp").Parse(kLogInOutXMLTemplate))
	err := t.Execute(&buf, liot)
	if err != nil {
		return "", err
	}
	str := html.UnescapeString(buf.String())

	return str, nil
}

func generateXMLAPIURL(host string) string {
	return fmt.Sprintf("https://%s/api/?type=user-id", host)
}

func LoginUserForIP(username string, ipaddr string, timeoutMin int, host string, key string) error {
	liot := LogInOutTemplateValue{
		Method:    "login",
		UserName:  username,
		IPAddress: ipaddr,
		Timeout:   timeoutMin,
	}

	str, err := generateXMLTemplateString(liot)
	if err != nil {
		log.Fatal(err)
		return err
	}

	uri := generateXMLAPIURL(host)

	v := url.Values{}
	v.Set("key", key)
	v.Set("cmd", str)

	resp, err := http.PostForm(uri, v)
	if err != nil {
		log.Fatal(err)
		return err
	}
	respStr, err := ioutil.ReadAll(resp.Body)

	log.Printf("resp => %s\n", respStr)

	return err
}

func LogoutUserForIP(username string, ipaddr string, host string, key string) error {
	liot := LogInOutTemplateValue{
		Method:    "logout",
		UserName:  username,
		IPAddress: ipaddr,
		Timeout:   0,
	}

	str, err := generateXMLTemplateString(liot)
	if err != nil {
		log.Fatal(err)
		return err
	}

	uri := generateXMLAPIURL(host)

	v := url.Values{}
	v.Set("key", key)
	v.Set("cmd", str)

	resp, err := http.PostForm(uri, v)
	if err != nil {
		log.Fatal(err)
		return err
	}
	respStr, err := ioutil.ReadAll(resp.Body)

	log.Printf("resp => %s\n", respStr)

	return nil
}

func DisableTLSVerification(yes bool) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: yes}
}
