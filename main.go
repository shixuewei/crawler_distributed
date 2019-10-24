package main

import (
	"flag"
	"learning/crawler_distributed/config"
	"learning/crawler_distributed/rpcsupport"
	"learning/crawler_distributed/save/client"
	"learning/crawler_distributed/worker/client2"
	"learning/crawler_goroutine/book/parser"
	"learning/crawler_goroutine/engine"
	"learning/crawler_goroutine/scheduler"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts(comma separated)")
)

//分布式版爬虫
func main() {
	flag.Parse()
	//得到Item(调用itemsaver保存)
	itemChan, err := client.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}
	//获取连接池中的client（不同的client以逗号隔开）
	pool := createClientPool(
		strings.Split(*workerHosts, ","))
	//爬取页面的处理器
	processor := client2.CreateProcessor(pool)

	//运行
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "https://www.biikan.com",
		Parser: engine.NewFuncParser(parser.ParseBookSort, config.ParseBookSort),
	})
}

//创建worker的Client连接池（让不同的worker去调用不同的client）
func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connection to %s: %v", h, err)
		}
	}

	//分发client
	out := make(chan *rpc.Client)
	go func() {
		//循环轮流发
		for {
			//一轮
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
