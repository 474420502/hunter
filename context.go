package hunter

import "github.com/474420502/requests"

// TaskContext 上下文
type TaskContext struct {
	share    map[string]interface{}
	hunter   *Hunter
	curNode  ITaskNode
	session  *requests.Session
	workflow *requests.Workflow
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

	bt := &BaseTask{task: itask}
	cxt.curNode.Children().Push(bt)
}

// AddParentTask 添加到当前任务队列
func (cxt *TaskContext) AddParentTask(itask ITask) {

	bt := &BaseTask{task: itask}
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

// Session Get return session *requests.Session
func (cxt *TaskContext) Session() *requests.Session {
	if cxt.session == nil {
		cxt.session = requests.NewSession()
	}
	return cxt.session
}

// SetSession Set session *requests.Session
func (cxt *TaskContext) SetSession(session *requests.Session) {
	cxt.session = session
}

// Workflow Get return Workflow *requests.Workflow. not exists, return nil
func (cxt *TaskContext) Workflow() *requests.Workflow {
	return cxt.workflow
}

// SetWorkflow Set Workflow *requests.Workflow
func (cxt *TaskContext) SetWorkflow(workflow *requests.Workflow) {
	cxt.workflow = workflow
}

// Hunt Hunt() = cxt.Workflow().Execute()
func (cxt *TaskContext) Hunt() (*requests.Response, error) {
	return cxt.workflow.Execute()
}
