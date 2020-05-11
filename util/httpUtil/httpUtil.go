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

type Util struct {
	Client   *http.Client
	AuthHead string
	AuthKey  string
}

type Jar struct {
	cookieStore map[string][]*http.Cookie
}

func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	if strings.Contains(u.Host, "login") {
		jar.cookieStore[u.Host] = cookies
	}
}

func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	if strings.Contains(u.Host, "passport") || strings.Contains(u.Host, "login") {
		return []*http.Cookie{}
	}
	return jar.cookieStore[u.Host]
}

func makeClient() *http.Client {
	jar := &Jar{
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

func MakeUtil(authHead string, authKey string) *Util {
	client := makeClient()
	data := &Util{Client: client, AuthHead: authHead, AuthKey: authKey}
	return data
}

func (util Util) Get(httpUrl string) (*Response, error) {
	req, err := http.NewRequest("POST", httpUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := util.Client.Do(req)
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

func (util Util) Post(httpUrl string, body string) (*Response, error) {
	var reader *strings.Reader
	reader = strings.NewReader(body)
	request, err := http.NewRequest("POST", httpUrl, reader)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMaker(util, httpUrl) {
		request.Header.Set(k, v)
	}
	resp, err := util.Client.Do(request)
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

func headerMaker(util Util, url string) map[string]string {
	if strings.Contains(url, "@self") {
		return util.loginHeaderBuilder(url)
	}
	return map[string]string{
		"Accept-Encoding": "identity",
		"Connection":      "Keep-Alive",
		"User-Agent":      "Dalvik/2.1.0 (Linux; U; Android 5.1.1; mi max Build/LMY48Z)",
	}
}

func (util Util) loginHeaderBuilder(url string) map[string]string {
	gmtTime := time.Now().UTC().Format(http.TimeFormat)
	mac := hmac.New(sha1.New, []byte(util.AuthKey))
	uri := strings.SplitN(url, "/", 4)
	mac.Write([]byte("POST\n" + gmtTime + "\n/" + uri[3]))
	auth := util.AuthHead + ":" + base64.StdEncoding.EncodeToString(mac.Sum(nil))
	var loginHeaders = map[string]string{
		//"Accept-Encoding": "gzip",
		"User-Agent":    "okhttp/3.12.1",
		"Content-Type":  "text/x-markdown; charset=utf-8",
		"Date":          gmtTime,
		"Authorization": auth,
	}
	return loginHeaders
}
