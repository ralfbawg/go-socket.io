// Package socketio
package socketio

import (
	"reflect"
	"sync"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

var testRedisBroadcast *redisBroadcast

type fakeRedisCon struct{}

func (conn *fakeRedisCon) Close() error { return nil }
func (conn *fakeRedisCon) Err() error   { return nil }
func (conn *fakeRedisCon) Do(string, ...interface{}) (reply interface{}, err error) {
	return nil, nil
}
func (conn *fakeRedisCon) Send(commandName string, args ...interface{}) error { return nil }
func (conn *fakeRedisCon) Flush() error                                       { return nil }
func (conn *fakeRedisCon) Receive() (reply interface{}, err error)            { return nil, nil }

func init() {
	testRedisBroadcast, _ = newRedisBroadcast(aliasRootNamespace, &RedisAdapterOptions{})
}

func Test_newRedisBroadcast(t *testing.T) {
	assert.NotNil(t, testRedisBroadcast)
}

func Test_redisBroadcast_AllRooms(t *testing.T) {
	a := &fakeRedisCon{}
	patches := ApplyFunc(redis.Dial, func(network, address string, options ...redis.DialOption) (redis.Conn,
		error) {
		return a, nil
	})
	ApplyMethod(reflect.TypeOf(a), "PSubscribe", func() {})
	ApplyMethod(reflect.TypeOf(a), "Subscribe", func() {})
	defer patches.Reset()
}

func Test_redisBroadcast_Clear(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_ForEach(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room string
		f    EachFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_Join(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room       string
		connection Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_Leave(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room       string
		connection Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_LeaveAll(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		connection Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_Len(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
			if got := bc.Len(tt.args.room); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisBroadcast_Rooms(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		connection Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
			if got := bc.Rooms(tt.args.connection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisBroadcast_Send(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room  string
		event string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_SendAll(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		event string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_allRooms(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
			if got := bc.allRooms(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisBroadcast_clear(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_getNumSub(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		channel string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
			got, err := bc.getNumSub(tt.args.channel)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNumSub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getNumSub() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisBroadcast_getRoomsByConn(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		connection Conn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
			if got := bc.getRoomsByConn(tt.args.connection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRoomsByConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisBroadcast_onMessage(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		channel string
		msg     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
			if err := bc.onMessage(tt.args.channel, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("onMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_redisBroadcast_onRequest(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_onResponse(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		msg []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_publish(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		channel string
		msg     interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_publishClear(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_publishMessage(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room  string
		event string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_send(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		room  string
		event string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}

func Test_redisBroadcast_sendAll(t *testing.T) {
	type fields struct {
		host       string
		port       string
		prefix     string
		pub        redis.PubSubConn
		sub        redis.PubSubConn
		nsp        string
		uid        string
		key        string
		reqChannel string
		resChannel string
		requests   map[string]interface{}
		rooms      map[string]map[string]Conn
		lock       sync.RWMutex
	}
	type args struct {
		event string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &redisBroadcast{
				host:       tt.fields.host,
				port:       tt.fields.port,
				prefix:     tt.fields.prefix,
				pub:        tt.fields.pub,
				sub:        tt.fields.sub,
				nsp:        tt.fields.nsp,
				uid:        tt.fields.uid,
				key:        tt.fields.key,
				reqChannel: tt.fields.reqChannel,
				resChannel: tt.fields.resChannel,
				requests:   tt.fields.requests,
				rooms:      tt.fields.rooms,
				lock:       tt.fields.lock,
			}
		})
	}
}
