package maouyan

import (
	"fmt"
	"goSpider/infra/url"
	"testing"
)

func TestGetMaoYanRankInfo(t *testing.T) {
	tests := []struct {
		url string
	}{
		{
			url: url.GetMaoYanURL(2019),
		},
	}

	for _, test := range tests {
		fmt.Println(GetMaoYanRankInfo(test.url))
	}

}
