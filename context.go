package hunter

// TaskContext 上下文
type TaskContext struct {
	share   map[string]interface{}
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
	bt := &BaseTask{}
	bt.SetTask(itask)
	cxt.curNode.Children().Push(bt)
}

// AddParentTask 添加到当前任务队列
func (cxt *TaskContext) AddParentTask(itask ITask) {
	bt := &BaseTask{}
	bt.SetTask(itask)
	cxt.curNode.Parent().Children().Push(bt)
}

// GetShare 获取share的数据, 存储用的
func (cxt *TaskContext) GetShare(key string) interface{} {
	if v, ok := cxt.share[key]; ok {
		return v
	}
	return nil
}

// SetShare 设置share的数据, 存储用的
func (cxt *TaskContext) SetShare(key string, value interface{}) {
	cxt.share[key] = value
}
