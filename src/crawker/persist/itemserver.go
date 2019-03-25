package persist

import (
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item server: go item #%d,%v\n", itemCount, item)
			itemCount++

			_, err := save(item)

			if err != nil {
				log.Printf("Save Item Error : %v,Item : %v", err, item)
			}
		}
	}()
	return out
}

func save(item interface{}) (id string, err error) {
	client, e := elastic.NewClient(
		// in docker  sniff turn false
		elastic.SetSniff(false))

	if e != nil {
		return "", e
	}

	response, e := client.Index().Index("data_profile").Type("zhenai").BodyJson(item).Do(context.Background())

	if e != nil {
		return "", e
	}

	return response.Id, nil

}
