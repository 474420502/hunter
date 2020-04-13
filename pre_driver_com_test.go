package hunter

import (
	"log"
	"testing"

	"github.com/lestrrat-go/libxml2"
	"github.com/tebeka/selenium"
)

type WebPreDriverUrl struct {
	PreChromeUrl
}

func (web *WebPreDriverUrl) Execute(cxt *TaskContext) {
	resp, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", resp.Content())
	wd := resp.GetResponse().(selenium.WebDriver)
	ele, err := wd.FindElement(selenium.ByXPATH, "//title")
	if err != nil {
		log.Panic(err)
	}
	title, err := ele.GetAttribute("text")
	if err != nil {
		log.Panic(err)
	}
	cxt.SetShare("driver-title", title)
}

func TestDriver(t *testing.T) {
	preurl := &WebPreDriverUrl{}
	preurl.PreUrl = "http://httpbin.org"

	hunter := NewHunter(preurl) // first params PreCurlUrl
	hunter.Execute()
	defer hunter.Stop()

	content := hunter.GetShare("test").([]byte)
	if content != nil {
		doc, err := libxml2.ParseHTML(content)
		if err != nil {
			t.Error(err)
		} else {
			if result, err := doc.Find("//title"); err == nil {
				iter := result.NodeIter()
				if iter.Next() {
					n := iter.Node()
					if n.TextContent() != "httpbin.org" {
						t.Error(n.TextContent())
					}
				} else {
					t.Error("can't xpath title")
				}
			} else {
				t.Error(err)
			}
		}

	}

	title := hunter.GetShare("driver-title").(string)
	if title != "httpbin.org" {
		t.Error(title)
	}
}
