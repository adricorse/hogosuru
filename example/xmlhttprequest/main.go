package main

import (
	"fmt"
	"net/url"

	"github.com/realPy/hogosuru/formdata"
	"github.com/realPy/hogosuru/progressevent"
	"github.com/realPy/hogosuru/xmlhttprequest"
)

func main() {

	endpoint, _ := url.Parse("http://localhost:9090/static.json")
	if xhr, err := xmlhttprequest.New(); err == nil {

		xhr.Open("GET", endpoint)
		xhr.SetOnload(func(x xmlhttprequest.XMLHTTPRequest) {

			fmt.Printf("XML HTTPRequest Loaded\n")

			if text, err := x.ResponseText(); err == nil {
				fmt.Printf("Resultat: %s\n", text)
			}

			if header, err := x.GetResponseHeader("Content-Type"); err == nil {
				fmt.Printf("Resultat: %s\n", header)
			}

		})

		xhr.SetOnProgress(func(x xmlhttprequest.XMLHTTPRequest, p progressevent.ProgressEvent) {
			println("onload")
			loaddata, err := p.Loaded()
			totaldata, err2 := p.Total()

			if err == nil && err2 == nil {
				fmt.Printf("On progress :%d / %d \n", loaddata, totaldata)
			}

		})
		xhr.Send()

	}

	if xhr, err := xmlhttprequest.New(); err == nil {

		xhr.Open("POST", endpoint)
		f, _ := formdata.New()
		f.AppendString("data", "pouet")
		xhr.SetOnload(func(x xmlhttprequest.XMLHTTPRequest) {

			fmt.Printf("XML HTTPRequest Loaded\n")

			if text, err := x.ResponseText(); err == nil {
				fmt.Printf("Resultat: %s\n", text)
			}

			if header, err := x.GetResponseHeader("Content-Type"); err == nil {
				fmt.Printf("Resultat: %s\n", header)
			}

		})
		xhr.SendForm(f)
	}
	ch := make(chan struct{})
	<-ch

}
