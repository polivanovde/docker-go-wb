package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/cockroachdb/cockroach-go/v2/crdb"
	"github.com/nats-io/stan.go"
)

func initStore() (*sql.DB, error) {

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	openDB := func() error {
		db, err = sql.Open("postgres", pgConnString)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	// 	восстанавливаем в кеш последние 10 записей
	rows, err := db.Query("SELECT order_uid,message FROM messages ORDER BY id LIMIT 10")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	type cacheReload struct {
		order_uid string
		message   string
	}

	for rows.Next() {
		var chElem cacheReload
		if err := rows.Scan(&chElem.order_uid, &chElem.message); err != nil {
			log.Fatal(err)
		}
		cache.Set(chElem.order_uid, chElem.message, 5*time.Minute)
	}

	return db, nil
}

func saveHandler(db *sql.DB, id, mess string, wg *sync.WaitGroup, mu *sync.Mutex, m *stan.Msg) {
	defer wg.Done()
	mu.Lock()
	err := crdb.ExecuteTx(context.Background(), db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				`INSERT INTO messages (order_uid, message)VALUES ($1, $2)
					   ON CONFLICT (order_uid) DO UPDATE SET order_uid=$1,message=$2`,
				id,
				mess,
			)
			if err != nil {
				return err
			}
            m.Ack()
			return nil
		})

	if err != nil {
		log.Fatal(err)
	}
	mu.Unlock()
}

func selectMessageById(db *sql.DB, mess string) string {
	rows := db.QueryRow("SELECT message FROM messages WHERE order_uid = $1", mess)

	var result string
	if err := rows.Scan(&result); err != nil {
		log.Println(err)
	}

	return result
}
