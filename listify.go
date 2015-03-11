package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

func notifyTime(db *sql.DB, channelName string) {
	for _ = range time.Tick(5 * time.Second) {
		_, err := db.Exec(fmt.Sprintf("NOTIFY %s, '%s'", channelName, time.Now()))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	var (
		user   = flag.String("db.user", "", "Database user")
		dbName = flag.String("db.name", "listify", "Database name")
	)
	flag.Parse()

	connInfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", *user, *dbName)
	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	eventCallback := func(event pq.ListenerEventType, err error) {
		if err != nil {
			log.Println(err)
		}
	}

	listener := pq.NewListener(connInfo, 5*time.Second, 30*time.Second, eventCallback)

	channelName := "time"
	if err := listener.Listen(channelName); err != nil {
		log.Fatal(err)
	}

	go notifyTime(db, channelName)

	for {
		select {
		case notification := <-listener.Notify:
			fmt.Printf("%s: %s\n", notification.Channel, notification.Extra)
		}
	}
}
