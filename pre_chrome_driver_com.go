package hunter

import (
	"fmt"
	"log"

	"github.com/Pallinder/go-randomdata"
	"github.com/tebeka/selenium"
)

// PreChromeUrl Chrome的url预处理
type PreChromeUrl struct {
	PreBaseDriverUrl
}

// Before 驱动的预处理
func (u *PreChromeUrl) Before(ctx *TaskContext) {

	var err error
	var service *selenium.Service

	if u.service == nil {
		for i := 0; i < 50; i++ {
			if u.Port == 0 {
				u.Port = randomdata.Number(10000, 50000)
			}
			service, err = selenium.NewChromeDriverService("chromedriver", u.Port)
			if err != nil {
				log.Println(i, err)
			} else {
				break
			}
		}

		u.service = service
	}

	if u.driver == nil {
		caps := selenium.Capabilities{"browserName": "chrome"}
		wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", u.Port))
		if err != nil {
			panic(err)
		}
		u.driver = wd
	}
}

// AddCookie 设置cookie
func (u *PreChromeUrl) AddCookie(cookie *selenium.Cookie) error {
	u.Hunt()
	return u.driver.AddCookie(cookie)
}

// DeleteCookie  删除cookie
func (u *PreChromeUrl) DeleteCookie(name string) error {
	return u.driver.DeleteCookie(name)
}

// DeleteAllCookies  删除所有cookie
func (u *PreChromeUrl) DeleteAllCookies() error {
	return u.driver.DeleteAllCookies()
}

// GetCookie  获取cookie
func (u *PreChromeUrl) GetCookie(name string) (selenium.Cookie, error) {
	return u.driver.GetCookie(name)
}

// GetCookies  获取所有cookie
func (u *PreChromeUrl) GetCookies() ([]selenium.Cookie, error) {
	return u.driver.GetCookies()
}
