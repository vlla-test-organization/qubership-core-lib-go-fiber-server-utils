package pprofpoint

import (
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestEnablePprofOnPort(t *testing.T) {
	port, err := getFreePort()
	assert.Nil(t, err)
	portStr := strconv.Itoa(port)
	EnablePprofOnPort(portStr)
	time.Sleep(time.Second) // wait for starting pprof on localhost:6060
	resp, err := http.Get("http://localhost:" + portStr + "/debug/pprof/")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
