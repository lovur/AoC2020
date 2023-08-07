package main

import (
	"bufio"
	"fmt"
	"os"
)

func slope_checker(lines []string, right int, down int) int {
    ret := 0
    pos := 0

    var i int = 0
    for x := 0; x < down; x++ { i++ } // Go down for first iter

    for i < len(lines) {
        size := len(lines[i])

        pos = pos + right

        rel_pos := pos % size
        if lines[i][rel_pos] == '#' {
            ret++
            // fmt.Printf("%s%s%s\n", line[:rel_pos], "X", line[rel_pos+1:])
        // } else {
        //     fmt.Printf("%s%s%s\n", line[:rel_pos], "O", line[rel_pos+1:])
        }
        for x := 1; x < down; x++ { i++ } 
        i++
    }
    return ret
}

func main() {
    f, err := os.Open("data.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    // s := bufio.NewScanner(f)
    // if err != nil {
    //     panic(err)
    // }

    // Solution 1

    // pos := 0
    // var trees int
    // s.Scan() // Extra at beginning to get rid of first row
    // for s.Scan() {
    //     line := s.Text()
    //     size := len(line)
    //
    //     pos = pos + 3
    //
    //     rel_pos := pos % size
    //     if line[rel_pos] == '#' {
    //         trees++
    //     }
    //     // fmt.Printf("%s%s%s\n", line[:rel_pos], "O", line[rel_pos+1:])
    // }


    // Solution 2
    s := bufio.NewScanner(f)
    if err != nil {
        panic(err)
    }
    s.Split(bufio.ScanLines)
    var lines []string

    for s.Scan() {
        lines = append(lines, s.Text())
    }

    var trees int

    trees = 1
    trees *= slope_checker(lines, 1, 1)
    trees *= slope_checker(lines, 3, 1)
    trees *= slope_checker(lines, 5, 1)
    trees *= slope_checker(lines, 7, 1)
    trees *= slope_checker(lines, 1, 2)

    fmt.Printf("\n\nt: %d\n", trees)
}
