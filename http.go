package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

var defaultHttpTestServer HttpTestServer

func init() {
	defaultHttpTestServer = HttpTestServer{}
}

// 与 http 相关的 testutil
type HttpTestServer struct {
	server *httptest.Server
}

// Server 返回默认的 HttpTest Server
func Server() *httptest.Server {
	return defaultHttpTestServer.Server()
}

// Server 返回默认的 HttpTest Server
func (hs HttpTestServer) Server() *httptest.Server {
	return hs.server
}

// SetupServer 设置默认的 http server. 类似 echo/gin 或者标准的 net/http 实现 http.Handler 接口
// 可以直接通过包级别的默认处理, 也可以通过 HttpTestServer 实例进行处理
func SetupServer(handler http.Handler) *httptest.Server {
	return defaultHttpTestServer.SetupServer(handler)
}

// SetupServer 设置默认的 http server. 类似 echo/gin 或者标准的 net/http 实现 http.Handler 接口
func (hs *HttpTestServer) SetupServer(handler http.Handler) *httptest.Server {
	hs.server = httptest.NewServer(handler)
	return hs.Server()
}

// CheckServer 检查 Server 是否初始化, 否则 panic
func CheckServer() {
	defaultHttpTestServer.CheckServer()
}

// CheckServer 检查 Server 是否初始化, 否则 panic
func (hs HttpTestServer) CheckServer() {
	if hs.server == nil {
		panic("TestSever 还未初始化")
	}
}

// ServerPath 返回使用对应 server 的 api url
func ServerPath(path string, server ...*httptest.Server) string {
	fmt.Println(defaultHttpTestServer)
	return defaultHttpTestServer.ServerPath(path, server...)
}

// ServerPath 返回使用对应 server 的 api url
func (hs HttpTestServer) ServerPath(path string, srv ...*httptest.Server) string {
	hs.CheckServer()
	if len(srv) > 1 {
		panic("only one httptest server allowed")
	}

	// 自动判断添加 "/"
	p := path
	if !strings.HasPrefix(path, "/") {
		p = "/" + path
	}
	return fmt.Sprintf("%s%s", hs.Server().URL, p)
}
