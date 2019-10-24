package main

import (
	"flag"
	"fmt"
	"learning/crawler_distributed/config"
	"learning/crawler_distributed/rpcsupport"
	"learning/crawler_distributed/save"
	"log"
	"strconv"

	"gopkg.in/olivere/elastic.v5"
)

//itemsaver的rpc服务端

//命令行输入命令，输入具体的host
var host = flag.String("host", "localhost:"+strconv.Itoa(0), "the host for me to listen on")

func main() {
	//Parse parses the command-line flags from os.Args[1:].
	flag.Parse()
	if *host == "localhost:0" {
		fmt.Println("must specify a host")
		return
	}
	log.Fatal(serveRpc(
		*host, config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &save.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
