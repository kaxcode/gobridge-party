package main

import (
        "fmt"
        "io/ioutil"
        "os"
        "strings"
)

func main() {
	b, err := ioutil.ReadFile("proverbs.txt")
	if err != nil {
		os.Exit(1)
	}

	lines := strings.SplitAfter(string(b), ".")

	for _, str := range lines {
		m := map[string]int{}


		for _, r := range str {
			c := string(r)
			m[c]++
		}

		fmt.Println(str, m)

	}
}