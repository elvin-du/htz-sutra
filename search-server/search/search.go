package search

import (
	"encoding/json"
	"github.com/go-ego/riot"
	"github.com/go-ego/riot/types"
	"htz/sutra/search-server/config"
)

var (
	// searcher 是协程安全的
	searcher = riot.Engine{}

	opts = types.EngineOpts{
		Using:         1,
		GseDict:       "./dict/dictionary.txt",
		StopTokenFile: "./dict/stop_tokens.txt",
		IndexerOpts: &types.IndexerOpts{
			IndexType: types.LocsIndex,
		},
		UseStore: true,
		//默认存储到当前目录的文件夹名称，如果config没有配置的话
		StoreFolder: "htz_sutra_search_engine_db",
	}
)

type SutraItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description "`
	Original    string `json:"original"`
	Explanation string `json:"explanation"`
	PlayedCount int64  `json:"played_count"` //TODO 为了以后作为排序自定义评分的要素之一
}

type SearchResult struct {
	Items  []*SutraItem `json:"items"`
	Tokens []string     `json:"tokens"` //搜索到的关键词
	// 搜索到的文档个数。注意这是全部文档中满足条件的个数，可能比返回的文档数要大
	NumDocs int `json:"num_docs"`
}

func init() {
	InitEngine()
}

func InitEngine() {
	if "" != config.DefaultConfig.SearchEngineDBPath {
		opts.StoreFolder = config.DefaultConfig.SearchEngineDBPath
	}
	// 初始化
	searcher.Init(opts)
	//defer searcher.Close()

	// 等待索引刷新完毕
	searcher.Flush()
}

func CloseEngine() {
	searcher.Close()
}

//labels: 用户设置的标签
//id: 必须唯一，且不能是"0"
func Index(content *SutraItem, labels ...string) {
	bin, err := json.Marshal(content)
	if nil != err {
		panic(err)
	}

	//添加拼音搜索
	pyTokens := searcher.PinYin(string(bin))
	tokens := make([]types.TokenData, 0, len(pyTokens))
	for _, t := range pyTokens {
		tokens = append(tokens, types.TokenData{Text: t})
	}

	docData := types.DocData{Tokens: tokens, Labels: labels, Content: string(bin)}
	searcher.Index(content.ID, docData, true)
	searcher.Flush() //添加后立刻生效
}

func Search(key string, outputOffset, maxOutputs int, labels ...string) *SearchResult {
	output := searcher.SearchDoc(types.SearchReq{
		Text:   key,
		Labels: labels,
		RankOpts: &types.RankOpts{
			//ScoringCriteria: &WeiboScoringCriteria{}, TODO 添加搜索的自定义评分
			OutputOffset: outputOffset,
			MaxOutputs:   maxOutputs,
		},
	})

	items := make([]*SutraItem, 0, len(output.Docs))
	for _, doc := range output.Docs {
		item := SutraItem{}
		err := json.Unmarshal([]byte(doc.Content), &item)
		if nil != err {
			panic(err)
		}
		items = append(items, &item)
	}

	return &SearchResult{Items: items, Tokens: output.Tokens, NumDocs: output.NumDocs}
}

func Remove(ID string) {
	searcher.RemoveDoc(ID, true)
	searcher.Flush() //添加后立刻生效
}
