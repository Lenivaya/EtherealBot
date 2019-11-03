package gifs

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetRandomGif(message string) (gif string, err error) {
	var searchWord *string
	var searchWordDefault = "cat"
	args := strings.SplitN(message, " ", 2)
	if len(args) != 2 {
		searchWord = &searchWordDefault
	} else {
		for i, v := range args {
			if i == 1 {
				searchWord = &v
				break
			}
		}
	}

	url := fmt.Sprintf("https://tenor.com/search/%s-gifs", *searchWord)

	resp, err := http.Get(url)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return "", err
	}

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)

		re := regexp.MustCompile("src=\"(https://media.tenor.com/images/\\w+\\/[\\w\\-\\.]+[\\-\\.][\\w\\-\\.]+)\"")
		matches := re.FindAllStringSubmatch(string(body), -1)

		Gifs := make([]string, len(matches))

		for index, match := range matches {
			Gifs[index] = match[1]
		}

		rand.Seed(time.Now().UnixNano())
		gif := Gifs[rand.Intn(len(Gifs))]
		return gif, nil
	}
	return "", nil
}
