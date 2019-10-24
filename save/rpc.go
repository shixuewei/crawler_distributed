package save

import (
	"learning/crawler_goroutine/engine"
	"learning/crawler_goroutine/save"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

//itemsaver的rpc调用方法、

//定义rpc服务的方法
type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

//保存
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := save.Save(s.Client, item, s.Index)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "OK"
	} else {
		log.Printf("Error saving item %v:%v", item, err)
	}
	return err
}
