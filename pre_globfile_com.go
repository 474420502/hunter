package hunter

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/474420502/requests"
)

// PreGlobFile Task的 file  预处理组件
type PreGlobFile string

func (u PreGlobFile) Before(ctx *TaskContext) {
	if strings.Count(ctx.Path(), "/") > 2 {
		return
	}

	m, err := filepath.Glob(string(u))
	if err != nil {
		panic(err)
	}

	for _, mfile := range m {
		itask := NewTaskByContext(ctx)
		itask.Elem().FieldByName("PreGlobFile").SetString(mfile)
		ctx.AddTask(itask.Interface().(ITask))
	}

	ctx.SetCancelNext(true) // cancel first Execute
}

func (u PreGlobFile) Hunt() (requests.IResponse, error) {

	f, err := os.Open(string(u))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	resp := &HResponse{}
	resp.Hcontent = data
	return resp, err
}
