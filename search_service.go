package itunes

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
)

const (
	// SearchParamTerm ...
	SearchParamTerm = "term"
	// SearchParamCountry ...
	SearchParamCountry = "country"
	// SearchParamMedia ...
	SearchParamMedia = "media"
	// SearchParamEntity ...
	SearchParamEntity = "entity"
	// SearchParamAttribute ...
	SearchParamAttribute = "attribute"
	// SearchParamCallback ...
	SearchParamCallback = "callback"
	// SearchParamLimit ...
	SearchParamLimit = "limit"
	// SearchParamLang ...
	SearchParamLang = "lang"
	// SearchParamVersion ...
	SearchParamVersion = "version"
	// SearchParamExplicit ...
	SearchParamExplicit = "explicit"
)

// SearchService ...
type SearchService struct {
	service *Service
}

// NewSearchService ...
func NewSearchService(service *Service) SearchService {
	return SearchService{
		service: service,
	}
}

// Do ...
func (ss *SearchService) Do(values *url.Values) (*SearchResults, error) {
	path := ss.service.basePath
	if values != nil {
		path += "?" + values.Encode()
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := ss.service.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("itunes: unknown status code " + strconv.FormatInt(int64(resp.StatusCode), 10))
	}

	var results *SearchResults
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	results.URL, err = url.Parse(path)
	if err != nil {
		return nil, err
	}

	return results, nil
}
