package itunes

import "net/http"

const (
	basePath = "https://itunes.apple.com/search"
)

// Service ...
type Service struct {
	Search SearchService

	basePath string
	client   *http.Client
}

// ServiceOpts ...
type ServiceOpts struct {
	BasePath string
	Client   *http.Client
}

// NewService ...
func NewService(options *ServiceOpts) *Service {
	if options == nil {
		options = &ServiceOpts{}
	}

	if options.BasePath == "" {
		options.BasePath = basePath
	}

	if options.Client == nil {
		options.Client = &http.Client{}
	}

	s := &Service{
		client:   options.Client,
		basePath: options.BasePath,
	}

	s.Search = NewSearchService(s)

	return s
}
