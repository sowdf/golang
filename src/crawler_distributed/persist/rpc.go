package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"model"
	"qupu/persist"
)

type ItemSaveServerService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveServerService) Save(item model.Spectrum, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
