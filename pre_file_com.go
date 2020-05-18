package hunter

import (
	"io/ioutil"
	"os"

	"github.com/474420502/requests"
)

// PreFile Task的 file  预处理组件
type PreFile string

func (u PreFile) Hunt() (requests.IResponse, error) {

	f, err := os.Open(string(u))
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	resp := &HResponse{}
	resp.Hcontent = data
	return resp, err
}
