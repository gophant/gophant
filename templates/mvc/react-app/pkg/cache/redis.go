package cache

import "log"

func NewRedis(url string) {
	log.Printf("(template) connect to redis: %s", url)
}
