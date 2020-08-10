package main

import (
	"github.com/zzlpeter/dawn-go/libs/redis"
	"log"
)


//var c, _ = redis.Dial("tcp", ":6379")
func readData() {
	//c := redis.GetRedisPool("rt").Get()
	//defer c.Close()
	rst, err := redis.GetRedisPool("rt").GetVal("name")
	log.Println(rst, err)
}

func readData1() {
	rst, err := redis.GetRedisPool("local").GetVal("name")
	log.Println(rst, err)
}

func main()  {

	//fmt.Println(c, err)
	for i := 0; i < 1; i++ {
		readData()
		readData1()
	}
	//time.Sleep(time.Second * 60)
	//fmt.Println("done......")
	//time.Sleep(time.Second * 30)
}
