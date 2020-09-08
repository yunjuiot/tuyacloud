package tuyacloud_test

import (
	"errors"
	"fmt"

	"github.com/yunjuiot/tuyacloud"
)

var (
	client *tuyacloud.Client
	r tuyacloud.Request
	err error
)

func ExampleError() {
	m := map[string]interface{}{}
	err = client.DoAndParse(r, &m)
	if err != nil {
		var apiErr tuyacloud.Error
		if errors.As(err, &apiErr) {
			fmt.Println(apiErr.Code, apiErr.Msg)
		}
	}
}
