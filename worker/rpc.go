package worker

import "learning/crawler_goroutine/engine"

//worker的rpc所需要实现的业务

//定义一个爬虫结构
type CrawlService struct {
}

//具体的爬取过程
func (CrawlService) Process(req Request, result *ParseResult) error {
	//将序列化的Request反序列化为原来的Request
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	//将得到的结果序列化为可传输的结果
	*result = SerializeResult(engineResult)
	return nil

}
