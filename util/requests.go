package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response struct {
	Body       []byte
	StatusCode int
}

type ResponseError struct {
	httpType string
	url      string
	locate   string
	err      error
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("请求发生错误, 方法%s, url: %s, err: %s", re.httpType, re.url, re.err)
}

func Get(url string, params map[string]string, headers map[string]string, cookies map[string]string) (*Response, error) {
	response := &Response{}
	req, err := http.Get(url)
	if err != nil {
		return nil, &ResponseError{
			httpType: "GET",
			url:      url,
			locate:   "创建GET请求",
			err:      err,
		}
	}
	defer req.Body.Close()
	// 加入参数
	if params != nil {
		url += "?"
		var paramsList []string
		for k, v := range params {
			paramsList = append(paramsList, k+"="+v)
		}
		url += strings.Join(paramsList, "&")
	}
	// 设置请求头
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	// 设置Cookies
	if cookies != nil {
		var cookiesList []string
		for k, v := range cookies {
			cookiesList = append(cookiesList, k+"="+v)
		}
		req.Header.Set("Cookie", strings.Join(cookiesList, ";"))
	}
	// 解析请求头
	response.Body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, &ResponseError{
			httpType: "GET",
			url:      "url",
			locate:   "解析请求头",
			err:      err,
		}
	}
	response.StatusCode = req.StatusCode
	return response, nil
}

func Post(
	url string, params map[string]string, body string, json map[string]string, headers map[string]string, cookies map[string]string) (*Response, error) {
	// 处理 body
	response := &Response{}
	client := &http.Client{}
	var reader *strings.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}
	if json != nil {
		var jsonList []string
		for k, v := range json {
			jsonList = append(jsonList, k+"="+v)
		}
		reader = strings.NewReader(strings.Join(jsonList, "&"))
	}
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, &ResponseError{
			httpType: "POST",
			url:      url,
			locate:   "创建http",
			err:      err,
		}
	}

	// 加入参数
	if params != nil {
		url += "?"
		var paramsList []string
		for k, v := range params {
			paramsList = append(paramsList, k+"="+v)
		}
		url += strings.Join(paramsList, "&")
	}
	// 设置请求头
	if headers != nil {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}
	// 设置Cookies
	if cookies != nil {
		var cookiesList []string
		for k, v := range cookies {
			cookiesList = append(cookiesList, k+"="+v)
		}
		request.Header.Set("Cookie", strings.Join(cookiesList, ";"))
	}
	// 解析请求头
	req, err := client.Do(request)
	if err != nil {
		return nil, &ResponseError{
			httpType: "POST",
			url:      url,
			locate:   "解析请求头",
			err:      err,
		}
	}
	defer req.Body.Close()
	response.Body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, &ResponseError{
			httpType: "POST",
			url:      url,
			locate:   "解析body",
			err:      err,
		}
	}

	response.StatusCode = req.StatusCode
	return response, nil
}
