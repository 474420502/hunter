package hunter

import (
	"encoding/json"
	"testing"
)

type WebPreUrl struct {
	PreGetUrl
}

func (web *WebPreUrl) Execute(cxt *TaskContext) {
	resp, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", string(resp.Content()))
}

func TestCasePreUrl(t *testing.T) {
	hunter := NewHunter()
	hunter.AddTask(&WebPreUrl{PreGetUrl: "http://httpbin.org/get"})
	hunter.Execute()

	data := make(map[string]interface{})
	content := hunter.GetShare("test").(string)
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		t.Error(err)
	}
	if iurl, ok := data["url"]; ok {
		if iurl.(string) != "http://httpbin.org/get" {
			t.Error(iurl)
		}
	}
}
