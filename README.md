# testutil
用于 golang 的测试相关的库

`go get github.com/wppurking/testutil`

## 功能介绍
* file: 提供 test 中用于读取 testdata 中的问题
* fixutres: 使用 [testfixtures](github.com/go-testfixtures/testfixtures/v3) 来解决数据库 fixtures 的问题, 固定文件路径
* http: 提供了默认的 HttpTestServer 初始化, 用于测试兼容 net/http 的 web 服务请求
