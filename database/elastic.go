package database

import (
	"github.com/acger/pair-svc/internal/config"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/zeromicro/go-zero/core/logx"
)

func NewElasticsearch(conf *config.Config) *es.Client {
	client, err := es.NewClient(es.Config{
		Addresses: conf.Elasticsearch.Addresses,
		Username:  conf.Elasticsearch.Username,
		Password:  conf.Elasticsearch.Password,
	})

	if err != nil {
		logx.Error("elasticsearch connect fail.")
		return nil
	}

	InitIndex(client, "acger_pair_0", AcgerPairBody)

	return client
}
