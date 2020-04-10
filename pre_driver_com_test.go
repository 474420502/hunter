package hunter

import "testing"

type WebPreDriverUrl struct {
	PreDriverUrl
}

func (web *WebPreDriverUrl) Execute(cxt *TaskContext) {
	resp, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}
	cxt.SetShare("test", resp.Content())
}

func TestDriver(t *testing.T) {
	hunter := NewHunter(&WebPreDriverUrl{PreDriverUrl("http://httpbin.org")}) // first params PreCurlUrl
	hunter.Execute()
}
