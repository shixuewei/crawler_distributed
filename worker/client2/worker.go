package client2

import (
	"learning/crawler_distributed/config"
	"learning/crawler_distributed/worker"
	"learning/crawler_goroutine/engine"
	"net/rpc"
)

//worker的rpc客户端
func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(req engine.Request) (engine.ParseResult, error) {
		//将请求序列化
		sReq := worker.SerializeRequest(req)
		//定义序列化的sResult
		var sResult worker.ParseResult
		//从客户队列传送客户给c
		c := <-clientChan
		//调用rpc服务
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		//将返回结果反序列化，得到真实的结果
		return worker.DeserializeResult(sResult), nil
	}
}
