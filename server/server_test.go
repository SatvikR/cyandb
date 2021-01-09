//
//    Copyright 2020 Satvik Reddy
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
//

package server

import (
	"testing"
)

const (
	testDBLocation = "../test.dat"
)

func TestServer_Set(t *testing.T) {
	type fields struct {
		Location string
	}
	type args struct {
		key string
		val string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "hello world",
			fields: fields{Location: testDBLocation},
			args: args{
				key: "hello",
				val: "world",
			},
			want: "world",
		},
		{
			name:   "long key/value",
			fields: fields{Location: testDBLocation},
			args: args{
				key: "KeyLongerThanNineCharacters",
				val: "ValueLongerThanNineCharacters",
			},
			want: "ValueLongerThanNineCharacters",
		},
		{
			name:   "re-writing same key",
			fields: fields{Location: testDBLocation},
			args: args{
				key: "hello",
				val: "new world",
			},
			want: "new world",
		},
	}
	server := &Server{
		Location: testDBLocation,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := server.Set(tt.args.key, tt.args.val); got != tt.want {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Get(t *testing.T) {
	type fields struct {
		Location string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "hello world",
			fields:  fields{Location: testDBLocation},
			args:    args{key: "hello"},
			want:    "new world",
			wantErr: false,
		},
		{
			name:    "long key/value",
			fields:  fields{Location: testDBLocation},
			args:    args{key: "KeyLongerThanNineCharacters"},
			want:    "ValueLongerThanNineCharacters",
			wantErr: false,
		},
		{
			name:    "non-existent key",
			fields:  fields{Location: testDBLocation},
			args:    args{key: "nonExistentKey"},
			want:    "",
			wantErr: true,
		},
	}
	server := &Server{
		Location: testDBLocation,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err, _ := server.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
