package jsonreader

import (
	"Scraper/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Fetch the json data by using http get method
func getListings(url string) *models.Listings {
	listings := new(models.Listings)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return listings
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:88.0) Gecko/20100101 Firefox/88.0")

	fmt.Println(req.Header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return listings
	}

	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	json.NewDecoder(resp.Body).Decode(listings)

	return listings
}

// Insert records into the listings table
func SaveMinifigListings(wg *sync.WaitGroup, ch chan int, itemId string) error {
	ch <- 1
	fmt.Println(itemId)
	url := "https://www.bricklink.com/ajax/clone/catalogifs.ajax?itemid=" + itemId + "&rpp=500&pi="

	err := parseJsonRecords(url, itemId)
	<-ch
	wg.Done()
	if err != nil {
		return err
	}
	return nil
}

// Insert records into the listings table
func SavePartListings(wg *sync.WaitGroup, ch chan int, itemId string, colorId string) error {
	ch <- 1
	fmt.Println(itemId)
	url := "https://www.bricklink.com/ajax/clone/catalogifs.ajax?itemid=" + itemId + "&color=" + colorId + "&rpp=500&pi="

	err := parseJsonRecords(url, itemId)
	<-ch
	wg.Done()
	if err != nil {
		return err
	}
	return nil
}

func parseJsonRecords(url string, itemId string) error {
	pageIndex := 1

	for {
		listings := getListings(url + strconv.Itoa(pageIndex))
		fmt.Println(len(listings.Records))
		for _, rec := range listings.Records {
			err := models.InsertRecord(itemId, rec)
			if err != nil {
				return err
			}
			//fmt.Println(rec.MDisplaySalePrice)
		}
		fmt.Println(url + strconv.Itoa(pageIndex))
		if len(listings.Records) != 500 {
			break
		}

		pageIndex++
	}

	return nil
}

// Return the number of pages for a particular item
func PagesOfListings(itemId string) int {

	pageIndex := 1

	url := "https://www.bricklink.com/ajax/clone/catalogifs.ajax?itemid=" + itemId + "&rpp=500&pi="

	for {
		listings := getListings(url + strconv.Itoa(pageIndex))

		if len(listings.Records) != 500 {
			break
		}

		pageIndex++
	}

	return pageIndex
}
