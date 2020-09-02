package smartlock

import (
	"testing"
)

func TestCreateTempPwd(t *testing.T) {
	type fields struct {
		DeviceID      string
		Name          string
		Password      string
		EffectiveTime int64
		InvalidTime   int64
		Phone         string
	}
	type mocked struct {
		MockedResponseBody string
		MockedError        error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "All",
			fields: fields{
				DeviceID:      "",
				Name:          "",
				Password:      "",
				EffectiveTime: 0,
				InvalidTime:   0,
				Phone:         "123456789",
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {

			},
		)
	}
}
