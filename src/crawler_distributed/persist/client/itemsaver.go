package client

import (
	"crawler_distributed/config"
	"crawler_distributed/rpcsupport"
	"log"
	"model"
)

func ItemSaver(host string) (chan model.Spectrum, error) {
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		return nil, err
	}

	out := make(chan model.Spectrum)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			var result string
			// Call Rpc
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Save Item Error : %v,Item : %v", err, item)
			}
		}
	}()
	return out, nil
}
