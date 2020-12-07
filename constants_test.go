package tuyacloud_test

import (
	"errors"
	"fmt"

	"github.com/yunjuiot/tuyacloud"
)

func ExampleError() {
	var (
		client *tuyacloud.Client
		r      tuyacloud.Request
		err    error
	)

	m := map[string]interface{}{}
	err = client.DoAndParse(r, &m)
	if err != nil {
		var apiErr tuyacloud.Error
		if errors.As(err, &apiErr) {
			fmt.Println(apiErr.Code, apiErr.Msg)
		}
	}
}
