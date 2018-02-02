# Brewerydbtests


Development of tests for BreweryDB using Go http://www.brewerydb.com/developers/docs.

To run the first unit test, just type in your terminal:

> go build dataGetter.go
> ./dataGetter <page>
  
With <page> being a integer number of a valid response page.
  
Example:

> go build dataGetter.go
> ./dataGetter 2

To run the automatic test, just type in your terminal:

> go test -v
