package hunter

import (
	"log"
	"testing"
)

type AreaCode struct {
	PreFile
}

func (a *AreaCode) Execute(cxt *TaskContext) {
	r, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}
	etor := NewExtractor(r.Content())
	xp, err := etor.XPath("//div[@class='ip']")
	if err != nil {
		panic(err)
	}

	log.Println(xp.ForEachText("./h4"))
	xpli, errlist := xp.ForEach("./h4/ul//li")
	if len(errlist) != 0 {
		panic(err)
	}
	log.Println(xpli.ForEachString("./h5/text()"))
}

func TestExtractor(t *testing.T) {
	ht := NewHunter()
	ht.AddTask(&AreaCode{"./testfile/area.html"})
	ht.Execute()
	t.Error()
}
