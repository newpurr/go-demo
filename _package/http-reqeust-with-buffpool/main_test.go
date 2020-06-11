package main

import (
	"reflect"
	"sync"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestAdapter_Request(t *testing.T) {
	type fields struct {
		pool sync.Pool
	}
	type args struct {
		r *Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			"test buffer and pool",
			fields{sync.Pool{}},
			args{&Request{StatusCode: 1234}},
			&Response{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := New()
			got, err := api.Request(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			spew.Dump(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Request() got = %v, want %v", got, tt.want)
			}
		})
	}
}
