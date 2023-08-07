package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	// m := make(map[rune]bool)
	m := make(map[rune]int)

	var ret int
	var n int
	for s.Scan() {
		// Part 1
		// if s.Text() == "" {
		// 	ret += len(m)
		// 	fmt.Printf("%d\n", ret)
		// 	n = 0
		// 	m = make(map[rune]bool)
		// }
		//
		// line := s.Text()
		// n++
		//
		// for _, c := range line {
		// 	m[c] = true
		// }
		// Part 2
		line := s.Text()
		if line == "" {
			for c, v := range m {
				fmt.Printf("v: %d c: %c n: %d", v, c, n)
				if v >= n {
					fmt.Printf(" adding\n")
					ret++
				}
			}
			m = make(map[rune]int)
			fmt.Printf("ret: %d\n\n", ret)
			n = 0
		} else {
			n++
			for _, c := range line {
				m[c]++
			}
		}

	}
	for c, v := range m {
		fmt.Printf("v: %d c: %c n: %d", v, c, n)
		if v >= n {
			fmt.Printf("adding\n")
			ret++
		}
	}
	fmt.Printf("%d\n", ret)
}
