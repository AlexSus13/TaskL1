package main

import (
        "fmt"
)

func main() {
        fmt.Println("result:", set([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func set(sl []string) []string {

	if len(sl) == 0 || len(sl) == 1 {
		return sl
	}

        m := make(map[string]int)

        var newSl []string

        for _, s := range sl {
                m[s]++
        }

        for k, _ := range m {
                newSl = append(newSl, k)
        }

        return newSl
}
