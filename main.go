package main

import (
	"fmt"
	"net/http"
)

const (
	url = "https://docs.github.com/en/apps/creating-github-apps/creating-github-apps/creating-a-github-app-using-url-parameters"
)

func ping(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	for i := 0; i < 50; i++ {
		// fmt.Printf("%d ping \n", i)
		err := ping(url)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("prosess done")
}
