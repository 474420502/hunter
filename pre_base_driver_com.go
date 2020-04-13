package hunter

import (
	"log"
	"net/http"

	"github.com/474420502/requests"
	"github.com/tebeka/selenium"
)

// https://github.com/tebeka/selenium

// PreBaseDriverUrl Task的 curl bash 预处理组件
type PreBaseDriverUrl struct {
	PreUrl  string
	Port    int
	service *selenium.Service
	driver  selenium.WebDriver
}

// Close 如果需要在最后执行销毁操作, 继承覆盖该方法
func (u *PreBaseDriverUrl) Close() error {

	if u.service != nil {
		// 直接退出, 所有销毁 直接忽略webdriver.Quit(). // Delete Session
		if err := u.service.Stop(); err != nil {
			return err
		}
	}
	return nil
}

// IResponse interface {
//     Content() []byte
//     GetStatus() string
//     GetStatusCode() int
//     GetHeader() http.Header
//     GetCookie() []*http.Cookie

//     // 返回不同的自定义的Response, 也可以是其他定义的结构体如WebDriver
//     GetResponse() interface{}
// }

// Content 内容
func (u *PreBaseDriverUrl) Content() []byte {
	content, err := u.driver.PageSource()
	if err != nil {
		log.Println(err)
	}
	return []byte(content)
}

// GetStatusCode 暂时为空
func (u *PreBaseDriverUrl) GetStatusCode() int {
	return 0
}

// GetStatus 内容 暂时为空
func (u *PreBaseDriverUrl) GetStatus() string {
	return ""
}

// GetHeader 暂时为空
func (u *PreBaseDriverUrl) GetHeader() http.Header {
	return nil
}

// GetCookie 暂时为空
func (u *PreBaseDriverUrl) GetCookie() []*http.Cookie {
	return nil
}

// GetResponse 返回 webdriver
func (u *PreBaseDriverUrl) GetResponse() interface{} {
	return u.driver
}

func (u *PreBaseDriverUrl) Hunt() (requests.IResponse, error) {
	err := u.driver.Get(string(u.PreUrl))
	return u, err
}
