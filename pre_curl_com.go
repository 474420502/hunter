package hunter

import (
	gcurl "github.com/474420502/gurl"
)

// PreCurlUrl Task的 curl bash 预处理组件
type PreCurlUrl string

func (h PreCurlUrl) Before(ctx *TaskContext) {
	gurl := gcurl.ParseRawCURL(string(h))
	ctx.GetHunter().SetSession(gurl.CreateSession())
	ctx.SetWorkflow(gurl.CreateWorkflow(ctx.Session()))
}
