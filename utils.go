package hunter

import "reflect"

// NewTaskByContext new your create task struct pointer.
func NewTaskByContext(ctx *TaskContext) reflect.Value {
	return reflect.New(reflect.TypeOf(ctx.current.Task()).Elem())
}
