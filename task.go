package hunter

import (
	pqueue "github.com/474420502/focus/priority_queue"
)

// IBefore 执行任务前处理
type IBefore interface {
	Before(cxt *TaskContext)
}

// // IExecute 执行任务
// type IExecute interface {
// 	Execute(ctx *TaskContext)
// }

// IAfter 执行任务后
type IAfter interface {
	After(cxt *TaskContext)
}

// ITask 任务接口
type ITask interface {
	Execute(ctx *TaskContext)
}

// IIdentity 身份id接口
type IIdentity interface {
	GetID() string
}

// ITaskNode 任务节点
type ITaskNode interface {
	Parent() ITaskNode
	SetParent(task ITaskNode)

	Children() *pqueue.PriorityQueue // ITaskNode
	SetChildren(children *pqueue.PriorityQueue)

	Task() ITask // ITaskNode
	SetTask(itask ITask)

	SetPath(path string)
	Path() string

	TaskID() string

	SetID(tid string)
}

// BaseTask 任务,必须包含子任务. 执行了第一个
type BaseTask struct {
	parent   ITaskNode
	children *pqueue.PriorityQueue // ITask类型
	path     string                // id 联合体, 禁用.命名
	task     ITask
	tid      string
}

// Parent 父
func (task *BaseTask) Parent() ITaskNode {
	return task.parent
}

// SetParent 设父
func (task *BaseTask) SetParent(itask ITaskNode) {
	task.parent = itask
}

// Children 孩子节点
func (task *BaseTask) Children() *pqueue.PriorityQueue {
	return task.children
}

// SetChildren 孩子节点
func (task *BaseTask) SetChildren(children *pqueue.PriorityQueue) {
	task.children = children
}

// Task 孩子节点
func (task *BaseTask) Task() ITask {
	return task.task
}

// SetTask 孩子节点
func (task *BaseTask) SetTask(itask ITask) {
	task.task = itask
}

// Path 路径 例如: web.subweb.subsubweb
func (task *BaseTask) Path() string {
	return task.path
}

// SetPath 路径 例如: web.subweb.subsubweb
func (task *BaseTask) SetPath(path string) {
	task.path = path
}

// TaskID 如果存在任务id就返回任务id, 否则生成自动生成的序列id
func (task *BaseTask) TaskID() string {
	if itid, ok := task.task.(IIdentity); ok {
		return itid.GetID()
	}
	return task.tid
}

// SetID 设置自动生成的序列id
func (task *BaseTask) SetID(tid string) {
	task.tid = tid
}

// // Task 孩子节点
// func (task *BaseTask) Task() ITask {
// 	return *task
// }
