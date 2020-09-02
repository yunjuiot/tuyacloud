package device_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/yunjuiot/tuyacloud/device"
	"github.com/yunjuiot/tuyacloud/tests"
)

func TestQueryDevices(t *testing.T) {
	ctrl := gomock.NewController(t)
	client, httpClient := tests.NewMockClient(t, ctrl)
	mocked := tests.Mocked{}
	mocked.ResponseBody = `{
    "result": {
        "devices": [
            {
                "active_time": 1579397684,
                "biz_type": 0,
                "category": "ms",
                "create_time": 1579397684,
                "icon": "smart/icon/ay15602459704817nQ2d/b4bd2c6b5b4ee64f399feb7f12b5b3c4.png",
                "id": "funny",
                "ip": "",
                "local_key": "funny",
                "model": "停用",
                "name": "YEEUU K1 Smart Lock Box",
                "online": false,
                "owner_id": "funny",
                "product_id": "funny",
                "product_name": "Test-YEEUU K1 Smart Lock Box",
                "status": [
                    {
                        "code": "residual_electricity",
                        "value": 81
                    },
                    {
                        "code": "battery_state",
                        "value": "high"
                    },
                    {
                        "code": "child_lock",
                        "value": false
                    },
                    {
                        "code": "unlock_password",
                        "value": 1
                    },
                    {
                        "code": "unlock_dynamic",
                        "value": 1
                    },
                    {
                        "code": "door_opened",
                        "value": false
                    },
                    {
                        "code": "alarm_lock",
                        "value": "wrong_finger"
                    },
                    {
                        "code": "reverse_lock",
                        "value": false
                    },
                    {
                        "code": "closed_opened",
                        "value": "unknown"
                    }
                ],
                "sub": true,
                "time_zone": "+08:00",
                "uid": "funny",
                "update_time": 1586590558,
                "uuid": "funny"
            }
        ],
        "last_id": "1586590558",
        "total": 3
    },
    "success": true,
    "t": 1599023681809
}`
	httpClient.EXPECT().Do(gomock.Any()).Return(mocked.Response())
	req := &device.QueryDevicesRequest{
		Schema:    "",
		ProductID: "",
		DeviceIDs: "funny",
		PageNo:    1,
		PageSize:  1,
	}
	var resp device.QueryDevicesResponse
	err := client.DoAndParse(req, &resp)
	require.Nil(t, err)
	require.Greater(t, len(resp.Devices), 0)
}
