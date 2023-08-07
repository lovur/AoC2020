package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func read_id(s string) int {
	// FBFBBFFRLR
	var hi, lo, row, col int
	hi = 127
	lo = 0
	for i := 0; i < 7; i++ {
		if s[i] == 'F' {
			hi = (hi + lo) / 2
			row = hi
		} else {
			lo = (hi+lo)/2 + 1
			row = lo
		}
		// fmt.Printf("%d thru %d r: %d\n", lo, hi, row)
	}

	hi = 7
	lo = 0
	for j := 7; j < 10; j++ {
		if s[j] == 'L' {
			hi = (hi + lo) / 2
			col = hi
		} else {
			lo = (hi+lo)/2 + 1
			col = lo
		}
		// fmt.Printf("%d thru %d c: %d\n", lo, hi, col)
	}
	return row*8 + col
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)

	// Part 1
	// var id, max int
	// for sc.Scan() {
	// 	line := sc.Text()
	// 	id = read_id(line)
	// 	if id > max {
	// 		max = id
	// 	}
	// }
	// fmt.Printf("%d", max)
	// Part 2
	var id int
	var seats []int
	for sc.Scan() {
		line := sc.Text()
		id = read_id(line)
		seats = append(seats, id)
	}

	sort.Ints(seats)
	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1]-seats[i] > 1 {
			fmt.Printf("%d", seats[i]+1)
		}
	}
}
