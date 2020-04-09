package hunter

import (
	"encoding/json"
	"log"
	"testing"
)

func init() {
	log.Println("测试最好使用 docker run -p 80:80 kennethreitz/httpbin")
}

type WebGet struct {
	PreGetUrl
}

func (web *WebGet) Execute(cxt *TaskContext) {
	resp, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", resp.Content())
}

func TestCasePreUrl(t *testing.T) {
	hunter := NewHunter()
	hunter.AddTask(&WebGet{PreGetUrl: "http://httpbin.org/get"})
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

type WebPost struct {
	PrePostUrl
}

func (web *WebPost) Execute(cxt *TaskContext) {
	wf := cxt.Workflow()
	wf.SetBodyAuto("param=hello form")
	resp, err := wf.Execute()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", resp.Content())
}

func TestCasePostForm(t *testing.T) {
	hunter := NewHunter()
	hunter.AddTask(&WebPost{PrePostUrl: "http://httpbin.org/post"})
	hunter.Execute()

	data := make(map[string]interface{})
	content := hunter.GetShare("test").(string)
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		t.Error(err)
	}

	if iform, ok := data["form"]; ok {
		form := iform.(map[string]interface{})
		if form["param"].(string) != "hello form" {
			t.Error(iform)
		}
	}
}

type WebSub struct {
	PrePostUrl
}

func (web *WebSub) Execute(cxt *TaskContext) {
	wf := cxt.Workflow()
	wf.SetBodyAuto(`{"a": "1","url":["http://httpbin.org/post","http://httpbin.org/get"]}`)
	resp, err := wf.Execute()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", resp.Content())

	data := make(map[string]interface{})
	json.Unmarshal([]byte(resp.ContentBytes()), &data)
	if urlList, ok := data["json"].(map[string]interface{})["url"].([]interface{}); ok {
		for _, is := range urlList {
			s := is.(string)
			cxt.AddTask(&WebSub1{PrePostUrl(s)})
		}
	}
}

type WebSub1 struct {
	PrePostUrl
}

func (web *WebSub1) Execute(cxt *TaskContext) {
	log.Panic(cxt.Path() + "." + cxt.TaskID())
}

func TestCaseWebSub(t *testing.T) {
	hunter := NewHunter()
	hunter.AddTask(&WebSub{"http://httpbin.org/post"})
	hunter.Execute()

	data := make(map[string]interface{})
	content := hunter.GetShare("test").(string)
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		t.Error(err)
	}

	if ijson, ok := data["json"]; ok {
		if j, ok := ijson.(map[string]interface{}); ok {
			if ia, ok := j["a"]; ok {
				if ia.(string) != "1" {
					t.Error(ia)
				}
			} else {
				t.Error(ia)
			}

		} else {
			t.Error(j)
		}

	}
}
