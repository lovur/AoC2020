package main

import (
	"bufio"
	"fmt"
	"os"
	// "reflect"
)

func main() {
    f, err := os.Open("data.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    s := bufio.NewScanner(f)

    var validCount int
    
    // req1 := map[string]bool {
    //     "byr": true,
    //     "cid": true,
    //     "ecl": true,
    //     "eyr": true,
    //     "hcl": true,
    //     "hgt": true,
    //     "iyr": true,
    //     "pid": true,
    // }
    // req2 := map[string]bool {
    //     "byr": true,
    //     "ecl": true,
    //     "eyr": true,
    //     "hcl": true,
    //     "hgt": true,
    //     "iyr": true,
    //     "pid": true,
    // }
    m := make(map[string]bool)

    for s.Scan() {
        line := s.Text()
        var code []byte
        if len(line) == 0 {
            // if reflect.DeepEqual(m, req1) || reflect.DeepEqual(m, req2) {
            //     validCount++
            // }
            _, ok := m["cid"]
            if (ok && len(m) == 7) || len(m) == 8 {
                validCount++
            }
            m = make(map[string]bool)
        } else {
            for i := 0; i < len(line); i++ {
                if line[i] == ':' {
                    m[string(code)] = true
                    code = nil
                } else {
                    if line[i] == ' ' {
                        code = nil
                    } else {
                        code = append(code, line[i])
                    }
                }
            }
        }
    }

    fmt.Printf("%d\n", validCount)
}
