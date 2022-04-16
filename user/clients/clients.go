package clients

import (
	"context"

	"github.com/HotPotatoC/twitter-clone/user/clients/elastic"
	"github.com/HotPotatoC/twitter-clone/user/clients/postgres"
	"github.com/HotPotatoC/twitter-clone/user/config"
	"github.com/elastic/go-elasticsearch/v8"
	"golang.org/x/sync/errgroup"
)

type Clients struct {
	DB postgres.Client
	ES *elasticsearch.Client
}

func Init(ctx context.Context, cfg config.Clients) (Clients, error) {
	var errg errgroup.Group

	c := Clients{}

	errg.Go(func() error {
		var err error
		c.DB, err = postgres.New(ctx, cfg.Postgres.GetDSN())
		if err != nil {
			return err
		}

		return nil
	})

	errg.Go(func() error {
		var err error
		c.ES, err = elastic.NewClient(elasticsearch.Config{
			Addresses: cfg.ElasticSearch.Addresses,
			Username:  cfg.ElasticSearch.Username,
			Password:  cfg.ElasticSearch.Password,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err := errg.Wait(); err != nil {
		return Clients{}, err
	}

	return c, nil
}
