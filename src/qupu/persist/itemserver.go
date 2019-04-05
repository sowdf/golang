package persist

import (
	"errors"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"model"
)

func ItemServer(index string) (chan model.Spectrum, error) {
	client, e := elastic.NewClient(
		// in docker  sniff turn false
		elastic.SetSniff(false))

	if e != nil {
		return nil, e
	}

	out := make(chan model.Spectrum)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item server: go item #%d,%v\n", itemCount, item)
			itemCount++

			err := save(client, item, index)

			if err != nil {
				log.Printf("Save Item Error : %v,Item : %v", err, item)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item model.Spectrum, index string) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	service := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		service.Id("qupu123" + item.Id)
	}
	_, e := service.Do(context.Background())

	if e != nil {
		return e
	}

	return nil

}
