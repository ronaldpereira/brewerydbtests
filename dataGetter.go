package main

import 	(
		"os"
		"strings"
        "io/ioutil"
        "net/http"
        "encoding/json"
		)

func getBeers (url string) []byte{
	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	return body
}

func mapMount (respString []byte) string{
	var m map[string]json.RawMessage

	err := json.Unmarshal(respString, &m)

	if err != nil {
		panic(err)
	}

	return string(m["data"])
}

func dataSplitter (rawData string) []string {
	splitData := strings.Split(rawData, ",\"")

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