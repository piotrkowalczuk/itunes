package itunes

import (
	"time"

	"net/url"
)

const (
	// WrapperTypeTrack ...
	WrapperTypeTrack = "track"
	// WrapperTypeCollection ...
	WrapperTypeCollection = "collection"
	// WrapperTypeArtist ...
	WrapperTypeArtist = "artist"

	// MediaTypeMovie ...
	MediaTypeMovie = "movie"
	// MediaTypePodcast ...
	MediaTypePodcast = "podcast"
	// MediaTypeMusic ...
	MediaTypeMusic = "music"
	// MediaTypeMusicVideo ...
	MediaTypeMusicVideo = "musicVideo"
	// MediaTypeAudiobook ...
	MediaTypeAudiobook = "audiobook"
	// MediaTypeShortFilm ...
	MediaTypeShortFilm = "shortFilm"
	// MediaTypeTVShow ...
	MediaTypeTVShow = "tvShow"
	// MediaTypeSoftware ...
	MediaTypeSoftware = "software"
	// MediaTypeEBook ...
	MediaTypeEBook = "ebook"
	// MediaTypeAll ...
	MediaTypeAll = "all"

	// KindBook ...
	KindBook = "book"
	// KindAlbum ...
	KindAlbum = "album"
	// KindCoachedAudio ...
	KindCoachedAudio = "coached-audio"
	// KindFeatureMovie ...
	KindFeatureMovie = "feature-movie"
	// KindInteractiveBooklet ...
	KindInteractiveBooklet = "interactive-booklet"
	// KindMusicVideo ...
	KindMusicVideo = "music-video"
	// KindPDF ...
	KindPDF = "pdf"
	// KindPodcast ...
	KindPodcast = "podcast"
	// KindPodcastEpisode ...
	KindPodcastEpisode = "podcast-episode"
	// KindSoftwarePackage ...
	KindSoftwarePackage = "software-package"
	// KindSong ...
	KindSong = "song"
	// KindTVEpisode ...
	KindTVEpisode = "tv-episode"
	// KindArtist ...
	KindArtist = "artist"
)

// SearchResult ...
type SearchResult struct {
	// WrapperType is the name of the object returned by the search request.
	WrapperType string `json:"wrapperType"`
	// PreviewURL is a URL referencing the 30-second preview file for the content associated with the returned media type. .
	PreviewURL string `json:"previewUrl"`
	// ReleaseDate ...
	ReleaseDate      *time.Time `json:"releaseDate"`
	Country          string     `json:"country"`
	Currency         string     `json:"currency"`
	PrimaryGenreName string     `json:"primaryGenreName"`
	RadioStationURL  string     `json:"radioStationUrl"`
	IsStreamable     bool       `json:"isStreamable"`
	DiscCount        int        `json:"discCount"`
	DiscNumber       int        `json:"discNumber"`
	Kind             string     `json:"kind"`
	// ArtworkURL100 is a URL for the artwork associated with the returned media type, sized to 100x100 pixels.
	ArtworkURL100 string `json:"artworkUrl100"`
	// ArtworkURL60 is a URL for the artwork associated with the returned media type, sized to 30x30 pixels.
	ArtworkURL60 string `json:"artworkUrl60"`
	// ArtworkURL30 is a URL for the artwork associated with the returned media type, sized to 30x30 pixels.
	ArtworkURL30 string `json:"artworkUrl30"`

	// TrackID ...
	TrackID int `json:"trackId"`
	// TrackName is the name of the track, song, video, TV episode, and so on returned by the search request.
	TrackName string `json:"trackName"`
	// TrackExplicitness is the Recording Industry Association of America (RIAA) parental advisory for the content returned by the search request.
	TrackExplicitness string `json:"trackExplicitness"`
	// TrackViewURL is a URL for the track. You can click the URL to view the content in the iTunes Store
	TrackViewURL string `json:"trackViewUrl"`
	// TrackViewURL is a URL for the track. You can click the URL to view the content in the iTunes Store
	TrackPrice        float32 `json:"trackPrice"`
	TrackCensoredName string  `json:"trackCensoredName"`
	// TrackTimeMillis is the returned track's time in milliseconds.
	TrackTimeMillis int64 `json:"trackTimeMillis"`
	TrackCount      int   `json:"trackCount"`
	TrackNumber     int   `json:"trackNumber"`

	// ArtistID ...
	ArtistID int `json:"artistId"`
	// ArtistName is the name of the artist returned by the search request.
	ArtistName string `json:"artistName"`
	ArtistType string `json:"artistType"`
	// ArtistExplicitness is the Recording Industry Association of America (RIAA) parental advisory for the content returned by the search request.
	ArtistExplicitness string `json:"artistExplicitness"`
	// ArtistViewURL is a URL for the artist. You can click the URL to view the content in the iTunes Store
	ArtistViewURL string `json:"artistViewUrl"`

	// CollectionID ...
	CollectionID int `json:"collectionId"`
	// CollectionExplicitness is the Recording Industry Association of America (RIAA) parental advisory for the content returned by the search request.
	CollectionExplicitness string `json:"collectionExplicitness"`
	// Kind of content returned by the search request.
	// CollectionName is the name of the album, TV season, audiobook, and so on returned by the search request, with objectionable words *'d out.
	CollectionName string `json:"collectionName"`
	// CollectionViewURL is a URL for the collection. You can click the URL to view the content in the iTunes Store
	CollectionViewURL      string `json:"collectionViewUrl"`
	CollectionCensoredName string `json:"collectionCensoredName"`
	// CollectionViewURL is a URL for the collection. You can click the URL to view the content in the iTunes Store
	CollectionPrice float32 `json:"collectionPrice"`
}

// Explicitness returns explicitness for current wrapper type.
func (sr *SearchResult) Explicitness() string {
	switch sr.WrapperType {
	case WrapperTypeArtist:
		return sr.ArtistExplicitness
	case WrapperTypeCollection:
		return sr.CollectionExplicitness
	case WrapperTypeTrack:
		return sr.TrackExplicitness
	default:
		return ""
	}
}

// Name returns name of an object for current wrapper type.
func (sr *SearchResult) Name() string {
	switch sr.WrapperType {
	case WrapperTypeArtist:
		return sr.ArtistName
	case WrapperTypeCollection:
		return sr.CollectionName
	case WrapperTypeTrack:
		return sr.TrackName
	default:
		return ""
	}
}

// ViewURL returns view name of an object for current wrapper type.
func (sr *SearchResult) ViewURL() string {
	switch sr.WrapperType {
	case WrapperTypeArtist:
		return sr.ArtistViewURL
	case WrapperTypeCollection:
		return sr.CollectionViewURL
	case WrapperTypeTrack:
		return sr.TrackViewURL
	default:
		return ""
	}
}

// TrackDuration returns track duration in milliseconds.
func (sr *SearchResult) TrackDuration() time.Duration {
	return time.Duration(sr.TrackTimeMillis) * time.Millisecond
}

// SearchResults ...
type SearchResults struct {
	Results []*SearchResult `json:"results"`
	Count   int             `json:"resultCount"`
	URL     *url.URL        `json:"-"`
}
