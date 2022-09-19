package auth

import (
	"fmt"
	"testing"
)

func TestNewJwt(t *testing.T) {
	type args struct {
		openid string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1",
			args: args{openid: "1234567890"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := NewJwt(tt.args.openid)
			if err != nil {
				panic(err)
			}
			fmt.Println(token)
			_, err = ParseJwt(token)
			if err != nil {
				panic(err)
			}
		})
	}
}
