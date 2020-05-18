package hunter

import "net/http"

// IResponse interface {
//     Content() []byte
//     GetStatus() string
//     GetStatusCode() int
//     GetHeader() http.Header
//     GetCookie() []*http.Cookie

//     // 返回不同的自定义的Response, 也可以是其他定义的结构体如WebDriver
//     GetResponse() interface{}
// }

// HResponse Empty for easy create
type HResponse struct {
	Hcontent  []byte
	Hstatus   string
	Hcode     int
	Hheader   http.Header
	Hcookies  []*http.Cookie
	Hresponse interface{}
}

func (resp *HResponse) Content() []byte {
	return resp.Hcontent
}

func (resp *HResponse) GetStatus() string {
	return resp.Hstatus
}

func (resp *HResponse) GetStatusCode() int {
	return resp.Hcode
}

func (resp *HResponse) GetHeader() http.Header {
	return resp.Hheader
}

func (resp *HResponse) GetCookie() []*http.Cookie {
	return resp.Hcookies
}

func (resp *HResponse) GetResponse() interface{} {
	return resp.Hresponse
}
