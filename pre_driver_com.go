package hunter

import (
	"fmt"
	"log"

	"github.com/474420502/requests"
	"github.com/tebeka/selenium"
)

// https://github.com/tebeka/selenium

// PreDriverUrl Task的 curl bash 预处理组件
type PreDriverUrl string

func (u PreDriverUrl) Before(ctx *TaskContext) {
	service, err := selenium.NewChromeDriverService("chromedriver", 1030)
	if err != nil {
		log.Panic(err)
	}
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 1030))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	err = wd.Get(string(u))
	if err != nil {
		panic(err)
	}

	ele, err := wd.FindElement(selenium.ByXPATH, "//title")
	log.Println(ele.Text())
	log.Panicln(ele.TagName())
}

func (u PreDriverUrl) Hunt() (*requests.Response, error) {
	return nil, nil
}
