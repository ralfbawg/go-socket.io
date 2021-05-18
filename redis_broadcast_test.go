// Package socketio
package socketio

import (
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

var testRedisBroadcast *redisBroadcast

type fakeCon struct{}

func (f fakeCon) Close() error {
	return nil
}

func (f fakeCon) Context() interface{} {
	return nil
}

func (f fakeCon) SetContext(ctx interface{}) {

}

func (f fakeCon) Namespace() string {
	return ""
}

func (f fakeCon) Emit(eventName string, v ...interface{}) {

}

func (f fakeCon) Join(room string) {

}

func (f fakeCon) Leave(room string) {

}

func (f fakeCon) LeaveAll() {

}

func (f fakeCon) Rooms() []string {
	return []string{}
}

func (f fakeCon) ID() string {
	return ""
}

func (f fakeCon) URL() url.URL {
	return url.URL{}
}

func (f fakeCon) LocalAddr() net.Addr {
	return nil
}

func (f fakeCon) RemoteAddr() net.Addr {
	return nil
}

func (f fakeCon) RemoteHeader() http.Header {
	return nil
}

var redisLocalServer *miniredis.Miniredis

var testConn *fakeCon

func init() {
	var err error
	redisLocalServer, err = miniredis.Run()
	if err != nil {
		panic(err)
	}
	testConn = &fakeCon{}
	tmpAddr := strings.Split(redisLocalServer.Addr(), ":")
	redisOpt := &RedisAdapterOptions{Host: tmpAddr[0], Port: tmpAddr[1]}
	testRedisBroadcast, _ = newRedisBroadcast(aliasRootNamespace, redisOpt)
}

func Test_newRedisBroadcast(t *testing.T) {
	defer redisLocalServer.Close()
	assert.NotNil(t, testRedisBroadcast)
}

func Test_redisBroadcast_AllRooms(t *testing.T) {

}

func Test_redisBroadcast_Clear(t *testing.T) {

}

func Test_redisBroadcast_ForEach(t *testing.T) {

}

func Test_redisBroadcast_Join(t *testing.T) {
	testRedisBroadcast.Join("/", testConn)
}

func Test_redisBroadcast_Leave(t *testing.T) {
	testRedisBroadcast.Leave("/", testConn)
}

func Test_redisBroadcast_LeaveAll(t *testing.T) {
	testRedisBroadcast.LeaveAll(testConn)
}

func Test_redisBroadcast_Len(t *testing.T) {
	assert.Equal(t, 1, testRedisBroadcast.Len("/"))
}

func Test_redisBroadcast_Rooms(t *testing.T) {

}

func Test_redisBroadcast_Send(t *testing.T) {

}

func Test_redisBroadcast_SendAll(t *testing.T) {

}

func Test_redisBroadcast_allRooms(t *testing.T) {

}

func Test_redisBroadcast_clear(t *testing.T) {

}

func Test_redisBroadcast_getNumSub(t *testing.T) {

}

func Test_redisBroadcast_getRoomsByConn(t *testing.T) {

}

func Test_redisBroadcast_onMessage(t *testing.T) {

}

func Test_redisBroadcast_onRequest(t *testing.T) {

}

func Test_redisBroadcast_onResponse(t *testing.T) {

}

func Test_redisBroadcast_publish(t *testing.T) {

}

func Test_redisBroadcast_publishClear(t *testing.T) {

}

func Test_redisBroadcast_publishMessage(t *testing.T) {

}

func Test_redisBroadcast_send(t *testing.T) {

}

func Test_redisBroadcast_sendAll(t *testing.T) {

}
