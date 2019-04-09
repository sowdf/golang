package main

import (
	"crawler_distributed/config"
	"crawler_distributed/persist"
	"crawler_distributed/rpcsupport"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	err := serviceRpc(fmt.Sprintf(":%d", config.ItemSaverProt), config.ElasticSearchIndex)
	if err != nil {
		panic(err)
	}

}

func serviceRpc(host string, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	return rpcsupport.ServiceRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
