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

// Return random url of bad quality image from google
func GetRandomShittyImage(message string) (image string, err error) {
	var searchWord *string
	var searchWordDefault = "cat"
	if len(strings.SplitN(message, " ", 2)) != 2 {
		searchWord = &searchWordDefault
	} else {
		for i, v := range strings.SplitN((message), " ", 2) {
			if i == 1 {
				searchWord = &v
				break
			}
		}
	}

	url := fmt.Sprintf("http://www.google.com/search?q=%s&tbm=isch", *searchWord)
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	if resp != nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		re := regexp.MustCompile("src=\"(http[^\"]+)\"")
		matches := re.FindAllStringSubmatch(string(body), -1)

		Images := make([]string, len(matches))

		for index, match := range matches {
			Images[index] = match[1]
		}

		rand.Seed(time.Now().UnixNano())

		image = fmt.Sprint(Images[rand.Intn(len(Images))])
		return image, nil
	}

	return "", err
}

// Gets random wallpaper page from wallhaven
func GetRandomWallFromWallhaven() (wallpage string, err error) {
	url := "https://wallhaven.cc/random"

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	if resp != nil && resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		re := regexp.MustCompile("href=\"(https://wallhaven.cc/w/\\w+)\"")
		matches := re.FindAllStringSubmatch(string(body), -1)

		Walls := make([]string, len(matches))

		for index, match := range matches {
			Walls[index] = match[1]
		}

		rand.Seed(time.Now().UnixNano())
		wall := fmt.Sprint(Walls[rand.Intn(len(Walls))])
		return wall, err
	}

	return "", err
}

// Gets a link of the wallpaper itself
func GetWallFromWallhaven() (wallpaper string, err error) {
	wallpage, err := GetRandomWallFromWallhaven()
	if err != nil {
		return "", err
	}

	resp, err := http.Get(wallpage)

	if err != nil {
		fmt.Println(err)
	}

	if resp != nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		re := regexp.MustCompile("src=\"(https://w.wallhaven.cc/full/\\w+\\/[\\w\\-\\.]+[\\-\\.][\\w\\-\\.]+)\"")
		matches := re.FindAllStringSubmatch(string(body), -1)

		return matches[0][1], err
	}

	return "", err
}
