package main

import (
	"flag"
	"fmt"
	"learning/crawler_distributed/rpcsupport"
	"learning/crawler_distributed/worker"
	"log"
	"strconv"
)

//worker的rpc服务端
var host = flag.String("host", "localhost:"+strconv.Itoa(0), "the host for me to listen on")

func main() {
	flag.Parse()
	if *host == "localhost:0" {
		fmt.Println("must specify a host")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		*host, worker.CrawlService{}))
}
