package hunter

import (
	pqueue "github.com/474420502/focus/priority_queue"
)

// IBefore 执行任务前处理
type IBefore interface {
	Before()
}

// // IExecute 执行任务
// type IExecute interface {
// 	Execute(ctx *TaskContext)
// }

// IAfter 执行任务后
type IAfter interface {
	After()
}

// ITask 任务接口
type ITask interface {
	Execute(ctx *TaskContext)
}

// ITaskNode 任务节点
type ITaskNode interface {
	Parent() ITaskNode
	SetParent(task ITaskNode)

	Children() *pqueue.PriorityQueue // ITaskNode
	SetChildren(children *pqueue.PriorityQueue)

	Task() ITask // ITaskNode
	SetTask(itask ITask)
}

// BaseTask 任务,必须包含子任务. 执行了第一个
type BaseTask struct {
	parent   ITaskNode
	children *pqueue.PriorityQueue // ITask类型
	task     ITask
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

// // Task 孩子节点
// func (task *BaseTask) Task() ITask {
// 	return *task
// }
