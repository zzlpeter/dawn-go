package main

import (
	_ "github.com/zzlpeter/dawn-go/libs/mysql"
	_ "github.com/zzlpeter/dawn-go/libs/redis"
	_routers "github.com/zzlpeter/dawn-go/routers"
)

// @title Dawn-go API By Gin
// @version 1.0
// @description Dawn-go API By Golang Gin
// @termsOfService https://github.com/zzlpeter/dawn-go
func main() {
	routers := _routers.InitRouter()
	routers.Run(":8001")
}
