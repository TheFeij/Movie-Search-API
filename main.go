package main

import (
	"Movie_Search_API/api"
	"Movie_Search_API/redis"
)

func main() {
	redis.InitializeRedisClient()
	api.StartServer()
}
