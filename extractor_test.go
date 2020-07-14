package hunter

import (
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type AreaCode struct {
	PreGlobFile
}

func (a *AreaCode) Execute(cxt *TaskContext) {
	r, err := cxt.Hunt()
	if err != nil {
		panic(err)
	}

	cxt.SetShare("cookies", r.GetCookie())
	cxt.SetShare("header", r.GetHeader())
	cxt.SetShare("status", r.GetStatus())
	cxt.SetShare("code", r.GetStatusCode())

	t := cxt.GetShare("test").(*testing.T)

	etor := NewExtractor(r.Content())
	xp, err := etor.XPaths("//div[@class='ip']")
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
	areastr1 := spew.Sdump(area)
	if len(area) != 345 {
		t.Error(len(area), area)
	}
	city, _ := xpli.ForEachString("./ul/li//text()")

	if len(city) != 3131 {
		t.Error(len(city))
	}

	names, errlist := xp.ForEachName("./ul//li")
	if len(errlist) != 0 {
		t.Error("error ForEachName")
	}
	if len(names) != 3131+345 || names[0] != "li" || names[len(names)-1] != "li" {
		log.Println(names, len(names))
	}

	area, _ = xpli.ForEachText("./h5")
	areastr2 := spew.Sdump(area)
	if len(area) != 345 || areastr2 != areastr1 {
		t.Error(len(area))
		return
	}

	area, _ = xpli.ForEachValue("./h5")
	areastr3 := spew.Sdump(area)
	if len(area) != 345 || areastr3 != areastr1 {
		t.Error(len(area), areastr1)
		return
	}

	h5values, _ := xpli.ForEachAttrValue("./h5", "value")
	areastr4 := spew.Sdump(h5values)
	if len(h5values) != 345 || h5values[0] != "hello h5" || h5values[len(h5values)-1] != "hello h5" {
		t.Error(len(h5values), areastr4)
		return
	}

	keyslist, _ := xpli.ForEachAttrKeys("./h5")
	if len(keyslist) != 345 {
		t.Error(len(keyslist))
		return
	}

	for _, keys := range keyslist {
		if keys[0] != "value" {
			t.Error("all h5 attribute, key = value")
			return
		}
	}

	attrs, _ := xpli.ForEachAttr("./h5")
	if len(attrs) != 345 {
		t.Error(len(attrs))
		return
	}

	types, _ := xpli.ForEachType("./h5")
	if len(types) != 345 {
		t.Error(len(types))
		return
	}

}

func TestExtractor(t *testing.T) {
	ht := NewHunter()
	ht.SetShare("test", t)
	ht.AddTask(&AreaCode{"./testfile/*.html"})
	ht.Execute()
}
