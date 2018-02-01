package main

import "os"
		
func main () {
	url := "http://api.brewerydb.com/v2/beers/?key="
	url += os.Args[1]
	url += "&availableId=1&p="
	url += os.Args[2]

	println(url)
}