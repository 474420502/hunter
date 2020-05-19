package hunter

import (
	"testing"
)

type AreaCode struct {
	PreGlobFile
}

func (a *AreaCode) Execute(cxt *TaskContext) {
	r, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}

	t := cxt.GetShare("test").(*testing.T)

	etor := NewExtractor(r.Content())
	xp, err := etor.XPath("//div[@class='ip']")
	if err != nil {
		panic(err)
	}

	pri, errl := xp.ForEachText("./h4")
	if len(errl) != 0 {
		t.Error(errl)
	}

	if len(pri) != 31 {
		t.Error(pri)
	}

	xpli, errlist := xp.ForEach("./ul//li")
	if len(errlist) != 0 {
		t.Error(err, xpli)
	}
	area, _ := xpli.ForEachString("./h5/text()")
	if len(area) != 345 {
		t.Error(len(area), area)
	}
	city, _ := xpli.ForEachString("./ul/li//text()")

	if len(city) != 3131 {
		t.Error(len(city))
	}
}

func TestExtractor(t *testing.T) {
	ht := NewHunter()
	ht.SetShare("test", t)
	ht.AddTask(&AreaCode{"./testfile/*.html"})
	ht.Execute()
}
