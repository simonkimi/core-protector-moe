package httpUtil

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Response struct {
	Body       []byte
	StatusCode int
}

type HttpContext struct {
	Client   *http.Client
	AuthHead string
	AuthKey  string
}

type CookieJar struct {
	cookieStore map[string][]*http.Cookie
}

func (cookieJar *CookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	if strings.Contains(u.Host, "login") {
		cookieJar.cookieStore[u.Host] = cookies
	}
}

func (cookieJar *CookieJar) Cookies(u *url.URL) []*http.Cookie {
	if strings.Contains(u.Host, "passport") || strings.Contains(u.Host, "login") {
		return []*http.Cookie{}
	}
	return cookieJar.cookieStore[u.Host]
}

func makeClient() *http.Client {
	jar := &CookieJar{
		cookieStore: make(map[string][]*http.Cookie),
	}
	var transport *http.Transport
	if os.Getenv("PROXY") != "" {
		transport = &http.Transport{Proxy: func(_ *http.Request) (*url.URL, error) {
			return url.Parse(os.Getenv("PROXY"))
		}}
	} else {
		transport = &http.Transport{Proxy: nil}
	}
	client := &http.Client{
		Transport: transport,
		Jar:       jar,
	}
	return client
}

func MakeUtil(authHead string, authKey string) *HttpContext {
	client := makeClient()
	data := &HttpContext{Client: client, AuthHead: authHead, AuthKey: authKey}
	return data
}

func (context HttpContext) Get(httpUrl string) (*Response, error) {
	req, err := http.NewRequest("POST", httpUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := context.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{}
	response.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response.StatusCode = resp.StatusCode
	return response, nil
}

func (context HttpContext) Post(httpUrl string, body string) (*Response, error) {
	var reader *strings.Reader
	reader = strings.NewReader(body)
	request, err := http.NewRequest("POST", httpUrl, reader)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMaker(context, httpUrl) {
		request.Header.Set(k, v)
	}
	resp, err := context.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{}
	response.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response.StatusCode = resp.StatusCode
	return response, nil
}

func JsonToStrings(json map[string]string) string {
	var jsonList []string
	for k, v := range json {
		jsonList = append(jsonList, k+"="+v)
	}
	return strings.Join(jsonList, "&")
}

func headerMaker(util HttpContext, url string) map[string]string {
	if strings.Contains(url, "@self") {
		return util.loginHeaderBuilder(url)
	}
	return map[string]string{
		"Accept-Encoding": "identity",
		"Connection":      "Keep-Alive",
		"User-Agent":      "Dalvik/2.1.0 (Linux; U; Android 5.1.1; mi max Build/LMY48Z)",
	}
}

func (context HttpContext) loginHeaderBuilder(url string) map[string]string {
	gmtTime := time.Now().UTC().Format(http.TimeFormat)
	mac := hmac.New(sha1.New, []byte(context.AuthKey))
	uri := strings.SplitN(url, "/", 4)
	mac.Write([]byte("POST\n" + gmtTime + "\n/" + uri[3]))
	auth := context.AuthHead + ":" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
	var loginHeaders = map[string]string{
		//"Accept-Encoding": "gzip",
		"User-Agent":    "okhttp/3.12.1",
		"Content-Type":  "text/x-markdown; charset=utf-8",
		"Date":          gmtTime,
		"Authorization": auth,
	}
	return loginHeaders
}
