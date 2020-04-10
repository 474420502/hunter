package hunter

import (
	"regexp"
	"testing"
)

type WebGurl struct {
	PreCurlUrl
}

func (web *WebGurl) Execute(cxt *TaskContext) {
	resp, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", resp.Content())
}

func TestCurlCom(t *testing.T) {
	curlBash := "curl 'http://httpbin.org/' -H 'Connection: keep-alive' -H 'Cache-Control: max-age=0' -H 'Upgrade-Insecure-Requests: 1' -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36' -H 'Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9' -H 'Accept-Encoding: gzip, deflate' -H 'Accept-Language: zh-CN,zh;q=0.9' --compressed --insecure"
	hunter := NewHunter(&WebGurl{PreCurlUrl(curlBash)}) // first params PreCurlUrl
	hunter.Execute()

	content := hunter.GetShare("test").(string)
	isMatchContent := regexp.MustCompile("<title>httpbin.org</title>").MatchString(content)
	if !isMatchContent {
		t.Error(content)
	}
}
