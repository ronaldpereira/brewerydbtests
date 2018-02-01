package main

import 	(
		"os"
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

func main () {
	url := "http://api.brewerydb.com/v2/beers/?key="
	url += os.Args[1]
	url += "&availableId=1&p="
	url += os.Args[2]

	respString := getBeers(url)

	os.Stdout.Write(respString)
}