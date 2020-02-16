package dongqiudi

/*
 懂球帝app 数据抓取分析
 fifa排行： 国家队、俱乐部、球员
 返回格式 json
*/

var (
	FiFaRankingURL = "https://sport-data.dongqiudi.com/soccer/biz/data/fifa_ranking?type=nation&&app=dqd&version=162&platform=android"
	ClubRankingURL = "https://sport-data.dongqiudi.com/soccer/biz/data/fifa_ranking?type=club&&app=dqd&version=162&platform=android"
	MarketValueURL = "https://sport-data.dongqiudi.com/soccer/biz/data/market_value_ranking?&app=dqd&version=162&platform=android"
	UrlList        = []string{FiFaRankingURL, ClubRankingURL, MarketValueURL}
)

func dongqiudiGetDoc(url string) (*goquery.Document, error) {
	req := gorequest.New()
}
