# iTunes
[![Build Status](https://travis-ci.org/piotrkowalczuk/itunes.svg?branch=master)](https://travis-ci.org/piotrkowalczuk/itunes)
[![](https://godoc.org/github.com/piotrkowalczuk/itunes?status.svg)](http://godoc.org/github.com/piotrkowalczuk/itunes)

Go library that wraps apple [iTunes Store Web Service](https://www.apple.com/itunes/affiliates/resources/documentation/itunes-store-web-service-search-api.html).

## Usage

```go
package main

import (
    "github.com/piotrkowalczuk/itunes"
)

func main() {
	service := itunes.NewService(nil)
	
	results, err := service.Search.Do(&url.Values{
		itunes.SearchParamTerm:      []string{"Above & Beyond: Group Therapy"},
		itunes.SearchParamMedia:     []string{"podcast"},
		itunes.SearchParamEntity:    []string{"podcast"},
		itunes.SearchParamExplicit:  []string{"Yes"},
		itunes.SearchParamAttribute: []string{"titleTerm"},
	})
	
	for _, r := range results.Results {
		log.Println(r.Name())
	}
}
```