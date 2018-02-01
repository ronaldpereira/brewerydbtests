package main

import 	(
		"os"
		"strings"
        "io/ioutil"
        "net/http"
        "encoding/json"
		)

func getBeers (url string) []byte{
	resp, err := http.Get(url) // Gets the response from the brewerydb API

	defer resp.Body.Close() // Will close the response connection right after function returns

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body) // Reads the response body

	return body
}

func mapMount (respString []byte) string{
	var m map[string]json.RawMessage

	err := json.Unmarshal(respString, &m) // Converts a Json byte object to a map object

	if err != nil {
		panic(err)
	}

	return string(m["data"])
}

func dataSplitter (rawData string) []string {
	splitData := strings.Split(rawData, ",\"") // Split the data string into separate fields by the ," marker

	return splitData
}

func main () {
	url := "http://api.brewerydb.com/v2/beers/?key="
	url += "1023032de7609989ab2ffdad23549ef5"
	url += "&availableId=1&p="
	url += os.Args[1]

	respString := getBeers(url)

	rawData := mapMount(respString)

	splitData := dataSplitter(rawData)

	println(splitData[0])
}