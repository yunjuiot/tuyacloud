package tuyacloud_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-log/log/log"
	"github.com/stretchr/testify/require"
	"github.com/yunjuiot/tuyacloud"
	"github.com/yunjuiot/tuyacloud/user"
)

func TestNewClient(t *testing.T) {
	accessID := os.Getenv("ACCESSID")
	accessKey := os.Getenv("ACCESSKEY")
	if accessID == "" || accessKey == "" {
		t.SkipNow()
		return
	}
	client := tuyacloud.NewClient(tuyacloud.APIEndpointCN, accessID, accessKey)
	token, err := client.Token()
	require.Nil(t, err)
	require.NotEmpty(t, token)
	fmt.Println(token)
}

func TestClient_TokenSign(t *testing.T) {
	type args struct {
		accessID  string
		accessKey string
		timestamp string
		token     string
	}
	tests := []struct {
		name     string
		args     args
		wantSign string
	}{
		{
			"normal",
			args{
				"1KAD46OrT9HafiKdsXeg",
				"4OHBOnWOqaEC1mWXOpVL3yV50s0qGSRC",
				"1588925778000",
				"3f4eda2bdec17232f67c0b188af3eec1",
			},
			"36C30E300F226B68ADD014DD1EF56A81EDB7B7A817840485769B9D6C96D0FAA1",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := tuyacloud.NewClient(tuyacloud.APIEndpointCN, tt.args.accessID, tt.args.accessKey)
				gotSign := c.TokenSign(tt.args.token, tt.args.timestamp)
				require.Equal(t, tt.wantSign, gotSign)
			},
		)
	}
}

func TestClient_PlainSign(t *testing.T) {
	type args struct {
		accessID  string
		accessKey string
		timestamp string
	}
	tests := []struct {
		name     string
		args     args
		wantSign string
	}{
		{
			"normal",
			args{
				"1KAD46OrT9HafiKdsXeg",
				"4OHBOnWOqaEC1mWXOpVL3yV50s0qGSRC",
				"1588925778000",
			},
			"CEAAFB5CCDC2F723A9FD3E91D3D2238EE0DD9A6D7C3C365DEB50FC2AF277AA83",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				c := tuyacloud.NewClient(tuyacloud.APIEndpointCN, tt.args.accessID, tt.args.accessKey)
				gotSign := c.PlainSign(tt.args.timestamp)
				require.Equal(t, tt.wantSign, gotSign)
			},
		)
	}
}

func ExampleClient() {
	client := tuyacloud.NewClient(
		tuyacloud.APIEndpointUS,
		"1KAD46OrT9HafiKdsXeg",
		"4OHBOnWOqaEC1mWXOpVL3yV50s0qGSRC",
		tuyacloud.WithLogger(log.New()),
	)
	req := &user.QueryUserInfoRequest{
		UID: "123456",
	}
	var info user.QueryUserInfoResponse
	err := client.DoAndParse(req, &info)
	if err != nil {
		panic(err)
	}
	// Blah Blah Blah...
}
