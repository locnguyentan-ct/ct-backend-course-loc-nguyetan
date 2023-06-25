package main

import (
	"context"
	"duplicateExcercise/session1"
	adListing "duplicateExcercise/session2/clients/ad-listing"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	// Session 1
	//sessionCode1()

	// Session2
	sessionCode2()

}

func sessionCode2() {
	wg := sync.WaitGroup{}

	logger := log.Default()

	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	c := adListing.NewClient("https://gateway.chotot.com",
		adListing.WithHTTPClient(httpClient),
		adListing.WithRetryTimes(5),
		adListing.WithLogger(logger),
		adListing.WithLoggerToFile("logfile.txt"),
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		ads, err := c.GetAdByCate(context.TODO(), adListing.CatePty)
		if err != nil {
			panic("GetAdByCate " + err.Error())
		}
		fmt.Printf("Number of Ads: %v\n", ads.Total)
		for i := 0; i < len(ads.Ads); i++ {
			fmt.Printf("===============")
			fmt.Printf("\nSubject: %v\n", ads.Ads[i].Subject)
		}
	}()

	go func() {
		defer wg.Done()
		ads, err := c.GetAdByCate(context.TODO(), adListing.CateVeh)
		if err != nil {
			panic("GetAdByCate " + err.Error())
		}
		fmt.Printf("Number of Ads: %v\n", ads.Total)
		for i := 0; i < len(ads.Ads); i++ {
			fmt.Printf("===============")
			fmt.Printf("\nSubject: %v\n", ads.Ads[i].Subject)
		}
	}()

	wg.Wait()
}

func sessionCode1() {

	slicesInt := []int{1, 5, 2, 5, 3}

	fmt.Println("Check duplicate by sorting")
	if session1.ContainsDuplicateBySorting(slicesInt) {
		fmt.Println("Duplicate in list")
	} else {
		fmt.Println("List value is unique")
	}

	fmt.Println("Check duplicate by map")
	if session1.ContainsDuplicateByMap(slicesInt) {
		fmt.Println("Duplicate in list")
	} else {
		fmt.Println("List value is unique")
	}

	var inputString string
	fmt.Printf("Input string for validate: ")
	fmt.Scan(&inputString)
	fmt.Printf("Is valid input: %v\n", session1.IsValid(inputString))

	fmt.Println("===================Exerise 2===================")
	fmt.Println("Using sorting")
	fmt.Printf("Is anagram: %v\n", session1.IsAnagramUsingSorting("iloveyou", "youlovei"))

	fmt.Println("Using map")
	fmt.Printf("Is anagram: %v\n", session1.IsAnagramUsingMap("anagram", "managra"))

}
