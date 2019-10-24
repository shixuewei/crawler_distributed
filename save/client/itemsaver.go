package client

import (
	"learning/crawler_distributed/config"
	"learning/crawler_distributed/rpcsupport"
	"learning/crawler_goroutine/engine"
	"log"
)

//itemsaver客户端

func ItemSaver(host string) (chan engine.Item, error) {
	//客户端对象
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	//输出
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver : got item #%d : %v ", itemCount, item)
			itemCount++
			//Call RPC to save item
			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item saver : error"+"saving item %v : %v", item, err)
			}
		}
	}()
	return out, nil
}
