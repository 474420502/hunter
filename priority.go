package hunter

import pqueue "github.com/474420502/focus/priority_queue"

// IPriority 优先接口
type IPriority interface {
	Priority() float64
}

// PriorityInt Int优先级
type PriorityInt int

// Priority Get priority
func (pri PriorityInt) Priority() float64 {
	return (float64)(pri)
}

// PriorityInt32 Int优先级
type PriorityInt32 int32

// Priority Get priority
func (pri PriorityInt32) Priority() float64 {
	return (float64)(pri)
}

// PriorityInt64 Int优先级
type PriorityInt64 int64

// Priority Get priority
func (pri PriorityInt64) Priority() float64 {
	return (float64)(pri)
}

// PriorityFloat32 Int优先级
type PriorityFloat32 float32

// Priority Get priority
func (pri PriorityFloat32) Priority() float64 {
	return (float64)(pri)
}

// CreatePriorityMaxQueue 创建最大优先队列
func CreatePriorityMaxQueue() *pqueue.PriorityQueue {
	return pqueue.New(priorityMax)
}

// CreatePriorityMinQueue 创建最小优先队列
func CreatePriorityMinQueue() *pqueue.PriorityQueue {
	return pqueue.New(priorityMin)
}

// priorityMax 最大值优先
func priorityMax(k1, k2 interface{}) int {
	p1, p2 := 0.0, 0.0

	if priority, ok := k1.(IPriority); ok {
		p1 = priority.Priority()
	}

	if priority, ok := k2.(IPriority); ok {
		p2 = priority.Priority()
	}

	if p1 > p2 {
		return 1
	}
	return -1
}

// priorityMin 最小值优先
func priorityMin(k1, k2 interface{}) int {
	p1, p2 := 0.0, 0.0

	if priority, ok := k1.(IPriority); ok {
		p1 = priority.Priority()
	}

	if priority, ok := k2.(IPriority); ok {
		p2 = priority.Priority()
	}

	if p1 < p2 {
		return 1
	}
	return -1
}
