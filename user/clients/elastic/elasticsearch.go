package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
)

// NewClient returns a new client with the given configuration.
func NewClient(cfg elasticsearch.Config) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(cfg)
}
