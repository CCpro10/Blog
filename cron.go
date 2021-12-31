package main

import (
	"time"

	"log"

	"github.com/robfig/cron"

	"Blog/models"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("*/5 * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	//只是为了阻塞main函数
	//可以使用select{}阻塞
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}

}
