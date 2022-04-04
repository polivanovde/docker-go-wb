package main

import (
	"encoding/json"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/nats-io/stan.go"
)

func createSubscriber() {
	log.Println("sub started")

	nc, err := stan.Connect("test-cluster", "client-123", stan.NatsURL("nats:4222"))
	if err != nil {
		log.Fatal(err)
	}
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	nc.QueueSubscribe("foo", "messages", func(m *stan.Msg) {
		if json.Valid(m.Data) {
			if err := json.Unmarshal(m.Data, &id); err != nil {
				log.Println(err)
				return
			}
			uid := id.OrderUID
			msg := string(m.Data)
			//TODO: хранить в оригинальных байтах будет менее ресурснозатратно, но в текущей задаче визуально удобнее

			if uid != "" && msg != "" {
				cache.Set(uid, msg, 5*time.Minute)
				wg.Add(1)
				go saveHandler(db, uid, msg, wg, mu)
				log.Printf("Received a message: %v\n", uid)
			} else {
				log.Println("некорректное сообщение")
			}
		}
	})
	wg.Wait()
	runtime.Goexit()

}
