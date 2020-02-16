package currency

import (
	"fmt"
	"goSpider/infra/errors"
	"io/ioutil"
	"net/http"
)

func GetHttpResponse(url string, ok bool) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.ErrorRequest
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.ErrorResponse
	}
	defer response.Body.Close()
	fmt.Println("response.StatusCode:", response.StatusCode)
	if response.StatusCode >= 300 && response.StatusCode <= 500 {
		return nil, errors.ErrorStatusCode
	}
	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	} else {
		return ioutil.ReadAll(response.Body)
	}
}

func main() {
	fmt.Println("this test for an cueency spider for BOC")
}
