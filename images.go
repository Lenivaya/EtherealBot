package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// Return random url of bad quality image
func GetRandomShittyImage(message string) (image string, err error) {
	var searchWord *string
	var searchWordDefault = "cat"

	if len(strings.SplitN(message, " ", 2)) != 2 {
		searchWord = &searchWordDefault
	} else {
		for i, v := range strings.SplitN((message), " ", 2) {
			fmt.Println(i, v)
			if i == 1 {
				searchWord = &v
				break
			}
		}
	}

	url := fmt.Sprintf("http://www.google.com/search?q=%s&tbm=isch", *searchWord)
	// url := "http://www.google.com/search?q=cat&tbm=isch"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err == nil && resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		re := regexp.MustCompile("src=\"(http[^\"]+)\"")
		matches := re.FindAllStringSubmatch(string(body), -1)

		MatchesSlice := make([]string, len(matches))

		for index, match := range matches {
			MatchesSlice[index] = match[1]
		}

		rand.Seed(time.Now().UnixNano())

		image = fmt.Sprint(MatchesSlice[rand.Intn(len(MatchesSlice))])
		return image, nil
	}

	return "", err
}
