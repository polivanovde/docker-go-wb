package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/nats-io/stan.go"
)

func submitter() {
	var wg sync.WaitGroup
	wg.Add(1)
	go createSubscriber(&wg)
	wg.Wait()
}

func createSubscriber(wg *sync.WaitGroup) {
	log.Println("sub started")
	defer wg.Done()

	nc, err := stan.Connect("test-cluster", "client-123", stan.NatsURL("nats:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("foo", func(m *stan.Msg) {
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
				saveHandler(db, uid, msg)
				//log.Println("Received a message: %v\n", cache)
			} else {
				log.Println("некорректное сообщение")
			}
		}
	})
	for { //не придумал ничего лучше чтобы получатель всегда ждал сообщений =(
	}
}
