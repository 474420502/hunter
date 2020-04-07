package hunter

// TaskContext 上下文
type TaskContext struct {
	hunter  *Hunter
	curNode ITaskNode
}

// NewContext 任务上下文
func NewContext() *TaskContext {
	return &TaskContext{}
}

// AddTask 添加到当前子任务队列
func (cxt *TaskContext) AddTask(itask ITask) {
	if children := cxt.curNode.Children(); children == nil {
		cxt.curNode.SetChildren(cxt.hunter.createQueue())
	}
	cxt.curNode.Children().Push(itask)
}

// AddParentTask 添加到当前任务队列
func (cxt *TaskContext) AddParentTask(itask ITask) {
	cxt.curNode.Parent().Children().Push(itask)
}
