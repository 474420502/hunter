package hunter

import (
	pqueue "github.com/474420502/focus/priority_queue"
)

// IGobalBefore 全局任务执行之前
type IGobalBefore interface {
	GobalBefore()
}

// IGobalAfter 全局任务执行之后
type IGobalAfter interface {
	GobalAfter()
}

// Hunter 任务相关 必须有序
type Hunter struct {
	cxt         *TaskContext
	task        ITaskNode
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
	hunter.task = &BaseTask{}
	hunter.task.SetParent(nil)
	hunter.task.SetChildren(hunter.createQueue())

	hunter.cxt = NewContext()
	hunter.cxt.curNode = hunter.task
	hunter.cxt.share = make(map[string]interface{})
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

// Execute 执行任务
func (hunter *Hunter) Execute() {
	hunter.recursionTasks(hunter.task)
}

func (hunter *Hunter) recursionTasks(itask ITaskNode) {
	for children := itask.Children(); children != nil && children.Size() > 0; {
		if itask, ok := children.Pop(); ok {
			tasknode := itask.(ITaskNode)
			tasknode.Task().Execute(hunter.cxt)
			hunter.recursionTasks(tasknode)
		}
	}
}

// Stop 停止任务
func (hunter *Hunter) Stop() {

}

// AddTask 执行任务
func (hunter *Hunter) AddTask(task ITask) {
	hunter.cxt.AddTask(task)
}

// Execute 执行
// func (hunter *Hunter) Execute() {
// 	if itask, ok := hunter.task.Children().Top(); ok {
// 		task := itask.(ITask)
// 	}
// }
