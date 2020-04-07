package hunter

import (
	"testing"
)

type Web struct {
}

func (web Web) Execute(cxt *TaskContext) {
	cxt.SetShare("123", 123)
}

func TestCase1(t *testing.T) {
	hunter := NewHunter()
	hunter.AddTask(&Web{})
	hunter.Execute()
	t.Error(hunter.cxt.GetShare("123"))
}
