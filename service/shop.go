package service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rnikrozoft/himymonsters-master-backend/cont"
	"github.com/rnikrozoft/himymonsters-master-backend/model"
	"github.com/rnikrozoft/himymonsters-master-backend/repository"
)

type ShopService struct {
	ID    string `mapstructure:"id"`
	Sheet string `mapstructure:"sheet"`
	Range string `mapstructure:"range"`

	Repo repository.RepositoryIF
}

func NewShop(Repo repository.RepositoryIF, id, sheetName, sheetRange string) *ShopService {
	return &ShopService{
		ID:    id,
		Sheet: sheetName,
		Range: sheetRange,

		Repo: Repo,
	}
}

func (s *ShopService) AddItems() error {
	sheetRange := fmt.Sprintf("%s%s", s.Sheet, s.Range)
	data := ReadGoogleSheet(s.ID, sheetRange)
	var items []model.Item
	for i, v := range data {
		if i != 0 {
			items = append(items, model.Item{
				ID:     v[0].(string),
				Title:  v[1].(string),
				Price:  v[2].(string),
				Detail: v[3].(string),
				Open:   v[4].(string),
			})
		}
	}

	m, err := json.Marshal(items)
	if err != nil {
		log.Fatal("cannot marshal shop item from google sheet: ", err)
		return err
	}

	storageWrite := &model.StorageWrite{
		Collection:      string(cont.Shop),
		Key:             string(cont.Items),
		UserID:          cont.System,
		Value:           string(m),
		PermissionRead:  2,
		PermissionWrite: 1,
	}

	if err := s.Repo.AddItems(storageWrite); err != nil {
		log.Fatal("cannot add shop items to storage: ", err)
		return err
	}
	return nil
}
