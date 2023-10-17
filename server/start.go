package main

import "s-url/core"

func init() {
	core.Init() // 初始化框架
}

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
func main() {

}
