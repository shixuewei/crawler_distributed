package worker

import (
	"errors"
	"fmt"
	"learning/crawler_distributed/config"
	"learning/crawler_goroutine/book/parser"
	"learning/crawler_goroutine/engine"
	"log"
)

//序列化后的Parser（提取函数名和函数的参数）
type SerializedParser struct {
	//函数名字
	Name string
	//该函数参数
	Args interface{}
}

//构建是为了将Request在网上传递
type Request struct {
	Url    string
	//此时在网上传递的是函数名和参数
	Parser SerializedParser
}

//构建是为了将Request在网上传递，
type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

//序列化Request
func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

//序列化Result
func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

//反序列化Request
func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

//反序列化Result
func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializing"+"request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)

	}
	return result
}

//反序列化Parser
func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseBookSort:
		return engine.NewFuncParser(parser.ParseBookSort, config.ParseBookSort), nil
	case config.ParseBookPage:
		return engine.NewFuncParser(parser.ParseBookPage, config.ParseBookPage), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseBook:
		return engine.NewFuncParser(parser.ParseBook, config.ParseBook), nil
	case config.ParseBookMes:
		if bookName, ok := p.Args.(string); ok {
			return parser.NewBookParser(bookName), nil
		} else {
			return nil, fmt.Errorf("invalid"+"arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}
