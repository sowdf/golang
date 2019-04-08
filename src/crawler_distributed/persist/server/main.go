package main

import (
	"crawler_distributed/persist"
	"crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	err := serviceRpc(":1234", "qupu")
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
	return rpcsupport.ServiceRpc(host, &persist.ItemSaveServerService{
		Client: client,
		Index:  index,
	})
}
