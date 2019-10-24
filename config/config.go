package config

const (
	//ElasticSearch
	ElasticIndex = "dating_book"
	//RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	//Parser names
	ParseBookSort = "ParseBookSort"
	ParseBookPage = "ParseBookPage"
	ParseBook     = "ParseBook"
	ParseBookMes  = "BookParser"
	NilParser     = "NilParser"
)
