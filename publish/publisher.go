package main

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"
	//для ввода в терминале
	// 	"bufio"
	// 	"os"
	"time"

	"github.com/nats-io/stan.go"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go createPublisher(&wg)
	wg.Wait()
}

func createPublisher(wg *sync.WaitGroup) {

	log.Println("pub started")
	// mark wait group done after createPublisher completes
	defer wg.Done()

	nc, err := stan.Connect("test-cluster", "client-124", stan.NatsURL("localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	//для ввода в терм
	//     for {
	for i := 0; i < 10; i++ {
		//для ввода в терм
		//         scanner := bufio.NewScanner(os.Stdin)
		//         scanner.Scan()
		//         strIn := []byte(scanner.Text())
		strIn := generateData(i)
		if errPub := nc.Publish("foo", strIn); errPub != nil {
			log.Fatal(errPub)
		}
	}

	log.Println("pub finish")
}

func generateData(i int) []byte {
	dateCreated, _ := time.Parse(time.RFC3339, "2021-11-26T06:22:19Z")
	var strI string = strconv.Itoa(i)

	s := ModelJson{
	    OrderUID:    "b563feb7b2b84b6test" + strI,
		Entry:       "WBIL" + strI,
		Delivery: DeliveryType{
			Name:    "Test Testov" + strI,
			Phone:   "+9720000000" + strI,
			Zip:     "2639809" + strI,
			City:    "Kiryat Mozkin" + strI,
			Address: "Ploshad Mira 15" + strI,
			Region:  "Kraiot" + strI,
			Email:   strI + "test@gmail.com",
		},
		Payment: PaymentType{
			Transaction:  "b563feb7b2b84b6test" + strI,
			RequestID:    "",
			Currency:     "USD" + strI,
			Provider:     "wbpay" + strI,
			Amount:       1817 + i,
			PaymentDt:    1637907727 + i,
			Bank:         "alpha" + strI,
			DeliveryCost: 1500 + i,
			GoodsTotal:   317 + i,
			CustomFee:    0 + i,
		},
		Items: ItemsType{
			Item{
				ChrtID:      9934930 + i,
				TrackNumber: "WBILMTESTTRACK" + strI,
				Price:       453 + i,
				Rid:         "ab4219087a764ae0btest" + strI,
				Name:        "Mascaras" + strI,
				Sale:        30 + i,
				Size:        "0" + strI,
				TotalPrice:  317 + i,
				NmID:        2389212 + i,
				Brand:       "Vivienne Sabo" + strI,
				Status:      202 + i,
			},
		},
		Locale:            "en" + strI,
		InternalSignature: "",
		CustomerID:        "test" + strI,
		DeliveryService:   "meest" + strI,
		Shardkey:          "9" + strI,
		SmID:              99 + i,
		DateCreated:       dateCreated,
		OofShard:          "1" + strI,
	}
	data, _ := json.Marshal(s)

	return data
}
