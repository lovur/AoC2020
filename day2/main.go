package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    f, err := os.Open("data.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    s := bufio.NewScanner(f)
    if err != nil {
        panic(err)
    }

    var validPasswordCount int

    for s.Scan() {
        line := s.Text()
        var minStr string
        var maxStr string
        var min, max int
        var letter rune
        var password string
        c := 0

        // get min 
        for {
            if line[c] == '-' {
                min, err = strconv.Atoi(minStr)
                if err != nil {
                    panic(err)
                }
                c++
                break
            } else {
                minStr += string(line[c])
                c++
            }
        }

        // get max 
        for {
            if line[c] == ' ' {
                max, err = strconv.Atoi(maxStr)
                if err != nil {
                    panic(err)
                }
                c++
                break
            } else {
                maxStr += string(line[c])
                c++
            }
        }

        // get letter
        letter = rune(line[c])
        // fmt.Printf("letter: %q", letter)
        c+=2

        // get password
        password = line[c:]
        var count int
        for _, c :=  range password {
            if c == letter {
                count++
            } 
        }
        // Solution 1
        // if count >= min && count <= max {
        //     validPasswordCount++
        // }
        // fmt.Printf("valid: %t letter: %q mn:%d a:%d mx%d\n", count >= min && count <= max, letter, min, count, max)

        // Solution 2
        isValid := (password[min] == byte(letter) && password[max] != byte(letter)) ||  (password[min] != byte(letter) && password[max] == byte(letter))
        if isValid {
            validPasswordCount++
        }
        fmt.Printf("%t c: %q f:%q l:%q\n", isValid, letter, password[min], password[max])
    }
    fmt.Printf("\nValid password count: %d\n", validPasswordCount)
}
