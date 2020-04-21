package hunter

import (
	"encoding/json"
	"log"
	"testing"
)

func init() {
	log.Println("recommend: docker run -p 80:80 kennethreitz/httpbin")
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
	cxt.SetShare("test", string(resp.Content()))
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
	json.Unmarshal(resp.Content(), &data)
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
	cxt.SetShare("test", cxt.Path()+"/"+cxt.TaskID())
}

func TestCaseWebSub(t *testing.T) {

	hunter := NewHunter()
	hunter.AddTask(&WebSub{"http://httpbin.org/post"})
	hunter.Execute()

	content := hunter.GetShare("test").(string)
	if content != "/0/1" {
		t.Error(content)
	}

}

type WebSavePoint struct {
	PrePostUrl
}

func (web *WebSavePoint) Execute(cxt *TaskContext) {
	wf := cxt.Workflow()
	wf.SetBodyAuto(`{"a": "1","url":["http://httpbin.org/post","http://httpbin.org/get"]}`)
	resp, err := wf.Execute()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", resp.Content())

	data := make(map[string]interface{})
	json.Unmarshal(resp.Content(), &data)
	if urlList, ok := data["json"].(map[string]interface{})["url"].([]interface{}); ok {
		for _, is := range urlList {
			s := is.(string)
			cxt.AddTask(&WebSavePoint1{PrePostUrl(s)})
		}
	}
}

type WebSavePoint1 struct {
	PrePostUrl
}

func (web *WebSavePoint1) Execute(cxt *TaskContext) {
	cxt.SetShare("test", cxt.Path()+"/"+cxt.TaskID())
	cxt.hunter.savePoint()
}

func TestSavePoint(t *testing.T) {

	hunter := NewHunter()
	hunter.AddTask(&WebSavePoint{"http://httpbin.org/post"})
	hunter.Execute()

	content := hunter.GetShare("test").(string)
	if content != "/0/1" {
		t.Error(content)
	}

}
