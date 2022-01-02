package funny

import (
	"errors"
	"testing"
	"time"
)

func TestGoFns(t *testing.T) {
	type args struct {
		fns []func() error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: testing.CoverMode(), args: struct{ fns []func() error }{fns: []func() error{
			func() error {
				t.Log("服务一开始执行")
				time.Sleep(time.Second * 6)
				t.Log("服务执行完毕，但这一行由于提前退出了，不会被打出,")
				return nil
			},
			func() error {
				time.Sleep(time.Second * 2)
				return errors.New("服务二错误了")
			},
			func() error {
				time.Sleep(time.Second * 6)
				return errors.New("服务三错误了")
			},
		}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GoFns(tt.args.fns...); (err != nil) != tt.wantErr {
				t.Errorf("GoFns() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
