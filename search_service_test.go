package itunes_test

import (
	"net/url"
	"testing"

	"github.com/piotrkowalczuk/itunes"
	"github.com/stretchr/testify/assert"
)

func TestSearchService_Do(t *testing.T) {
	service := itunes.NewService(nil)

	testSearchService_DoEmptyRequest(t, service)
	testSearchService_DoBasicRequest(t, service)
	testSearchService_DoReachRequest(t, service)
}

func testSearchService_DoEmptyRequest(t *testing.T, s *itunes.Service) {
	results, err := s.Search.Do(nil)

	if assert.NoError(t, err) {
		assert.Equal(t, 0, results.Count)
		assert.Len(t, results.Results, 0)
	}

}

func testSearchService_DoBasicRequest(t *testing.T, s *itunes.Service) {
	results, err := s.Search.Do(&url.Values{
		itunes.SearchParamTerm: []string{"marron"},
	})

	if assert.NoError(t, err) {
		assert.NotEqual(t, 0, results.Count)
		assert.True(t, len(results.Results) > 0)
	}
}

func testSearchService_DoReachRequest(t *testing.T, s *itunes.Service) {
	title := "Above & Beyond: Group Therapy"

	results, err := s.Search.Do(&url.Values{
		itunes.SearchParamTerm:      []string{title},
		itunes.SearchParamMedia:     []string{"podcast"},
		itunes.SearchParamEntity:    []string{"podcast"},
		itunes.SearchParamExplicit:  []string{"Yes"},
		itunes.SearchParamAttribute: []string{"titleTerm"},
	})

	if assert.NoError(t, err) {
		assert.Equal(t, 1, results.Count)
		if assert.Len(t, results.Results, 1) {
			result := results.Results[0]

			assert.Equal(t, title, result.CollectionName)
			assert.Equal(t, title, result.TrackName)
			assert.Equal(t, itunes.WrapperTypeTrack, result.WrapperType)
			assert.Equal(t, itunes.KindPodcast, result.Kind)
			assert.Equal(t, "Music", result.PrimaryGenreName)
		}
	}
}
