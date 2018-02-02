// Unicode codes available is this link: https://unicode-table.com/pt/#control-character

package main

import 	(
		"os"
		"strings"
		"reflect"
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

func matchPattern (pattern string) string {
	dictionary := 	map[string]string{ // Dictionary with the key and the expected type of variable of it
					"name":"string",
					"shortName":"string",
					"nameDisplay":"string",
					"description":"string",
					"foodPairings":"string",
					"abv":"float",
					"abvMin":"float",
					"abvMax":"float",
					"ibu":"float",
					"ibuMin":"float",
					"ibuMax":"float",
					"srmId":"int",
					"srmMin":"float",
					"srmMax":"float",
					"ogMin":"float",
					"fgMin":"float",
					"fgMax":"float",
					"glasswareId":"int",
					"availableId":"int",
					"categoryId":"int",
					"styleId":"int",
					"isOrganic":"string",
					"status":"string",
					"statusDisplay":"string",
					"servingTemperature":"string",
					"servingTemperatureDisplay":"string",
					"originalGravity":"float",
					"createDate":"date",
					"updateDate":"date",
					}

	return dictionary[pattern]
}

func testPattern (pattern, value string) bool {

	if pattern == "string" {
		return testString(value)
	} else if pattern == "int" {
		return testInt(value)
	} else if pattern == "float" {
		return testFloat(value)
	} else if pattern == "date" {
		return testDate(value)
	}
	println(value)
	return false
}

func testString (value string) bool {
	content := strings.Split(value, "\":\"")[1] // Split the key and the content with the '":"' delimiter
	content = strings.Split(content, "\"")[0] // Gets only the content between the '""'

	return reflect.TypeOf(content).String() == "string"
}

func testInt (value string) bool {
	content := strings.Split(value, "\":")[1] // Split the key and the content with the '":' delimiter

	for i := range content {
		if rune(content[i]) < 48 || rune(content[i]) > 57 { // If the specified char is unicode < 48 (0) or > 57 (9), it isn't a integer number, so returns false
			return false
		}
	}

	return true
}

func testFloat (value string) bool {
	content := strings.Split(value, "\":\"")[1] // Split the key and the content with the '":"' delimiter
	content = strings.Split(content, "\"")[0] // Gets only the content between the '""'

	for i := range content {
		if (rune(content[i]) < 48 || rune(content[i]) > 57) && rune(content[i]) != 46 { // If the specified char is unicode < 48 (0) or > 57 (9) and it isn't a unicode 46 (.), it isn't a float number, so returns false
			return false
		}
	}

	return true
}

func testDate (value string) bool {
	content := strings.Split(value, "\":\"")[1] // Split the key and the content with the '":"' delimiter
	content = strings.Split(content, "\"")[0] // Gets only the content between the '""'

	for i := range content {
		if (rune(content[i]) < 48 || rune(content[i]) > 57) && rune(content[i]) != 32 && rune(content[i]) != 45 && rune(content[i]) != 58 { // If the specified char is unicode < 48 (0) or > 57 (9) and it isn't a unicode 32 ( ) unicode 45 (-) and unicode 58 (:), it isn't a datetime format, so returns false
			println(content[i])
		}
	}

	return true
}

func tester (splitData []string) {
	key := ""
	pattern := ""
	passed := 0
	failed := 0

	for i := range splitData {
		key = strings.Split(splitData[i], "\":")[0]
		if key != "" {
			pattern = matchPattern(key)
			if testPattern(pattern, splitData[i]) {
				passed += 1
			} else {
				failed += 1
			}
		}
	}
	println("\n\nPATTERN AND TYPE TESTS:\nPASSED:",passed,"\nFAILED:",failed)
}

func main () {
	url := "http://api.brewerydb.com/v2/beers/?key="
	url += "1023032de7609989ab2ffdad23549ef5"
	url += "&availableId=1&p="
	url += os.Args[1]

	respString := getBeers(url)

	rawData := mapMount(respString)

	splitData := dataSplitter(rawData)

	tester(splitData)
}