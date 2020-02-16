package downloader

import (
	"fmt"
	"goSpider/infra/errors"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
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
	fmt.Println(response.StatusCode)
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

// posthttpresponse
func PostHttpResponse(url string, body string, ok bool) ([]byte, error) {
	payload := strings.NewReader(body)
	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, errors.ErrorRequest
	}
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return nil, errors.ErrorResponse
	}
	fmt.Println("response.StatusCode:", response.StatusCode)
	defer response.Body.Close()
	if ok {
		utf8Content := transform.NewReader(response.Body, simplifiedchinese.GBK.NewDecoder())
		return ioutil.ReadAll(utf8Content)
	}
	return ioutil.ReadAll(response.Body)

}
