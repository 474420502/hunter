package hunter

import "github.com/474420502/requests"

type IHunt interface {
	Hunt() (requests.IResponse, error)
}

// TaskContext 上下文
type TaskContext struct {
	hunter    *Hunter
	temporary *requests.Temporary

	parent  ITaskNode
	current ITaskNode
	autoid  int

	cancel bool
}

// NewContext 任务上下文
func NewContext() *TaskContext {
	return &TaskContext{}
}

// AddTask 添加到当前子任务队列
func (cxt *TaskContext) AddTask(itask ITask) {
	if children := cxt.current.Children(); children == nil {
		cxt.current.SetChildren(cxt.hunter.createQueue())
	}
	bt := &BaseTask{task: itask}
	cxt.current.Children().Push(bt)
}

// AddParentTask 添加到当前任务队列
func (cxt *TaskContext) AddParentTask(itask ITask) {

	bt := &BaseTask{task: itask}
	cxt.current.Parent().Children().Push(bt)
}

// GetShare 获取share的数据, 存储用的
func (cxt *TaskContext) GetShare(key string) interface{} {
	if v, ok := cxt.hunter.share[key]; ok {
		return v
	}
	return nil
}

// SetShare 设置share的数据, 存储用的
func (cxt *TaskContext) SetShare(key string, value interface{}) {
	cxt.hunter.share[key] = value
}

// Session Get return session *requests.Session
func (cxt *TaskContext) Session() *requests.Session {
	if cxt.hunter.Session() == nil {
		cxt.hunter.SetSession(requests.NewSession())
	}
	return cxt.hunter.Session()
}

// Temporary Get return Temporary *requests.Temporary. not exists, return nil
func (cxt *TaskContext) Temporary() *requests.Temporary {
	return cxt.temporary
}

// SetTemporary Set Temporary *requests.Temporary
func (cxt *TaskContext) SetTemporary(temporary *requests.Temporary) {
	cxt.temporary = temporary
}

// TaskID Get Task ID
func (cxt *TaskContext) TaskID() string {
	return cxt.current.TaskID()
}

// Path curren  Task tree path.
func (cxt *TaskContext) Path() string {
	return cxt.current.Path()
}

// GetHunter 获取share的数据, 存储用的
func (cxt *TaskContext) GetHunter() *Hunter {
	return cxt.hunter
}

// SetCancelNext set cancel next func(execute after ...)
func (cxt *TaskContext) SetCancelNext(is bool) {
	cxt.cancel = is
}

// Hunt Hunt() = cxt.Temporary().Execute()
func (cxt *TaskContext) Hunt() (requests.IResponse, error) {
	if ihunt, ok := cxt.current.Task().(IHunt); ok {
		return ihunt.Hunt()
	}
	return cxt.temporary.Execute()
}
