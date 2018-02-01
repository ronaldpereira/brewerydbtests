package main

import 	(
		"testing"
		"reflect"
		)

func TestDataGetterPageOne (t *testing.T) { // Tests the first page of the response
	url := "http://api.brewerydb.com/v2/beers/?key=1023032de7609989ab2ffdad23549ef5&availableId=1&p=1"
    respString := getBeers(url)

    if respString == nil {
        t.Errorf("Error from getBeers function: Cannot get beers from API.")
    } else if reflect.TypeOf(respString).String() != "[]uint8" {
		t.Errorf("Error from getBeers function: respString type is incorrect. Got %T. Expected []uint8", respString)
	}

	rawData := mapMount(respString)

	if reflect.TypeOf(rawData).String() != "string" {
		t.Errorf("Error from mapMount function: rawData type is incorrect. Got %T. Expected string", rawData)
	} else if rawData == "" {
		t.Errorf("Error from mapMount function: rawData string is empty.")
	}

	splitData := dataSplitter(rawData)

	if splitData[0] == "" {
		t.Errorf("Error from dataSplitter function: splitData string is empty.")
	}
}

func TestDataGetterPageTwo (t *testing.T) { // Tests the second page of the response
	url := "http://api.brewerydb.com/v2/beers/?key=1023032de7609989ab2ffdad23549ef5&availableId=1&p=2"
	respString := getBeers(url)

    if respString == nil {
        t.Errorf("Error from getBeers function: Cannot get beers from API.")
    } else if reflect.TypeOf(respString).String() != "[]uint8" {
		t.Errorf("Error from getBeers function: respString type is incorrect. Got %T. Expected string", respString)
	}

	rawData := mapMount(respString)

	if reflect.TypeOf(rawData).String() != "string" {
		t.Errorf("Error from mapMount function: RawData type is incorrect. Got %T. Expected string", rawData)
	} else if rawData == "" {
		t.Errorf("Error from mapMount function: RawData string is empty.")
	}

	splitData := dataSplitter(rawData)

	if splitData[0] == "" {
		t.Errorf("Error from dataSplitter function: splitData string is empty.")
	}
}