package hunter

import (
	"strconv"

	pqueue "github.com/474420502/focus/priority_queue"
	"github.com/474420502/requests"
)

// IGobalBefore 全局任务执行之前
type IGobalBefore interface {
	GobalBefore(cxt *TaskContext)
}

// IGobalAfter 全局任务执行之后
type IGobalAfter interface {
	GobalAfter(cxt *TaskContext)
}

// Hunter 任务相关 必须有序
type Hunter struct {
	share map[string]interface{}

	session *requests.Session

	tasks       []ITask
	createQueue func() *pqueue.PriorityQueue
}

// NewHunter 默认最大优先
func NewHunter() *Hunter {
	return NewPriorityMaxHunter()
}

// NewPriorityHunter 自定义优先处理队列
func NewPriorityHunter(queueCreator func() *pqueue.PriorityQueue) *Hunter {
	hunter := &Hunter{}
	hunter.createQueue = queueCreator

	// hunter.task = &BaseTask{}
	// hunter.task.SetParent(nil)
	// hunter.task.SetChildren(hunter.createQueue())

	// hunter.cxt = NewContext()
	// hunter.cxt.curNode = hunter.task

	hunter.share = make(map[string]interface{})
	return hunter
}

// NewPriorityMaxHunter 最大优先
func NewPriorityMaxHunter() *Hunter {
	return NewPriorityHunter(CreatePriorityMaxQueue)
}

// NewPriorityMinHunter 最小优先
func NewPriorityMinHunter() *Hunter {
	return NewPriorityHunter(CreatePriorityMinQueue)
}

// Session Get session *requests.Session
func (hunter *Hunter) Session() *requests.Session {
	return hunter.session
}

// SetSession Set session *requests.Session
func (hunter *Hunter) SetSession(session *requests.Session) {
	hunter.session = session
}

// GetShare 获取share的数据, 存储用的
func (hunter *Hunter) GetShare(key string) interface{} {
	if v, ok := hunter.share[key]; ok {
		return v
	}
	return nil
}

// SetShare 设置share的数据, 存储用的
func (hunter *Hunter) SetShare(key string, value interface{}) {
	hunter.share[key] = value
}

// Execute 执行任务
func (hunter *Hunter) Execute() {
	for _, task := range hunter.tasks {
		hunter.execute(task)
	}
}

// Execute 执行任务
func (hunter *Hunter) execute(task ITask) {
	cxt := NewContext()

	btask := &BaseTask{}
	// btask.SetTask(task)
	btask.SetParent(nil)
	btask.SetChildren(hunter.createQueue())

	cxt.current = btask
	cxt.hunter = hunter
	cxt.AddTask(task)

	hunter.recursionTasks(cxt)
}

func (hunter *Hunter) recursionTasks(cxt *TaskContext) {

	autoid := 0

	for children := cxt.current.Children(); children != nil && children.Size() > 0; {
		if itask, ok := children.Pop(); ok {

			ncxt := NewContext()

			tasknode := itask.(ITaskNode)
			task := tasknode.Task()

			ncxt.curPath = cxt.current.Path()
			if itid, ok := task.(IIdentity); ok {
				ncxt.curTaskID = itid.GetID()
			} else {
				ncxt.curTaskID = strconv.Itoa(autoid)
				autoid++
			}

			if before, ok := task.(IBefore); ok {
				before.Before(cxt)
			}

			tasknode.Task().Execute(cxt)

			if after, ok := task.(IAfter); ok {
				after.After(cxt)
			}

			tasknode.SetPath(cxt.current.Path() + "." + ncxt.curTaskID)
			ncxt.parent = cxt.current
			ncxt.current = tasknode
			hunter.recursionTasks(ncxt)
		}
	}

}

// Stop 停止任务
func (hunter *Hunter) Stop() {

}

// AddTask 执行任务
func (hunter *Hunter) AddTask(task ITask) {
	hunter.tasks = append(hunter.tasks, task)
}

// Execute 执行
// func (hunter *Hunter) Execute() {
// 	if itask, ok := hunter.task.Children().Top(); ok {
// 		task := itask.(ITask)
// 	}
// }
