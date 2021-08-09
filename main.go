package main

import (
	jsonreader "Scraper/JSONReader"
	"Scraper/common"
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {
	dbpool := common.Init()
	defer dbpool.Close()

	fmt.Println(dbpool.Stat().MaxConns())
	ch := make(chan int, dbpool.Stat().MaxConns())
	var wg sync.WaitGroup

	// scrape minifigures
	rows, err := dbpool.Query(context.Background(), "Select item_id from minifigures")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}

	for rows.Next() {
		var item_id int
		err = rows.Scan(&item_id)
		if err != nil {
			return
		}
		go jsonreader.SaveMinifigListings(&wg, ch, strconv.Itoa(item_id))
		wg.Add(1)
	}

	//scrape parts (item_id, color_id)
	rows, err = dbpool.Query(context.Background(), "Select item_id, color_id from parts")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}

	for rows.Next() {
		var item_id int
		var color_id int
		err = rows.Scan(&item_id, &color_id)
		if err != nil {
			return
		}
		go jsonreader.SavePartListings(&wg, ch, strconv.Itoa(item_id), strconv.Itoa(color_id))
		wg.Add(1)
	}

	wg.Wait()
}
