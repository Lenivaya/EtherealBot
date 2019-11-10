package gifs

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"time"
)

func GetRandomGif(message string) (gif string, err error) {
	url := fmt.Sprintf("https://tenor.com/search/%s-gifs", message)

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
