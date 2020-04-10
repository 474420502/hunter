package hunter

import (
	"fmt"
	"log"
	"runtime"

	"github.com/474420502/requests"
	"github.com/tebeka/selenium"
)

// https://github.com/tebeka/selenium

// PreDriverUrl Task的 curl bash 预处理组件
type PreDriverUrl struct {
	url     string
	service *selenium.Service
	driver  selenium.WebDriver
}

func (u *PreDriverUrl) Before(ctx *TaskContext) {
	service, err := selenium.NewChromeDriverService("chromedriver", 1030)
	if err != nil {
		log.Panic(err)
	}
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	u.service = service

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 1030))
	if err != nil {
		panic(err)
	}
	u.driver = wd

	runtime.SetFinalizer(&[]interface{}{service, wd}, func(obj interface{}) {
		iobj := obj.([]interface{})
		service := iobj[0].(*selenium.Service)
		service.Stop()

		wd := iobj[1].(selenium.WebDriver)
		wd.Quit()
	})

	err = wd.Get(string(u.url))
	if err != nil {
		panic(err)
	}

	ele, err := wd.FindElement(selenium.ByXPATH, "//title")
	log.Println(ele.Text())
	log.Println(ele.TagName())
}

func (u *PreDriverUrl) Hunt() (requests.IResponse, error) {
	err := u.driver.Get(string(u.url))

	return nil, err
}
