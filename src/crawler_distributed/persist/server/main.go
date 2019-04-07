package main

import (
	"crawler_distributed/persist"
	"crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	err = rpcsupport.ServiceRpc(":1234", persist.ItemSaveServerService{
		Client: client,
		Index:  "qupu",
	})
}
