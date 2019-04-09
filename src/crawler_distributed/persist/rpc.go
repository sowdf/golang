package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"model"
	"qupu/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item model.Spectrum, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	log.Printf("Item %v Saved ", item)
	if err == nil {
		*result = "ok"
	}
	return err
}
