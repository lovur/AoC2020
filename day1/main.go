package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    f, err := os.Open("data.txt")

    if err != nil {
        panic(err)
    }
    defer f.Close()

    s := bufio.NewScanner(f)
    s.Split(bufio.ScanLines) 
    
    var nums []int
    for s.Scan() {
        val, err := strconv.Atoi(s.Text())
        if err != nil {
            panic(err)
        }

        nums = append(nums, val) 
    }

    // Part 1
    // for i := 0; i < len(nums); i++ {
    //     for j := i+1; j < len(nums); j++ {
    //         if nums[i] + nums[j] == 2020 { 
    //             fmt.Printf("%d\n", nums[i] * nums[j])
    //         }
    //     }
    // }
    
    // Part 2 
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            for k := j+1; k < len(nums); k++ {
                if nums[i] + nums[j] + nums[k] == 2020 { 
                    fmt.Printf("%d\n", nums[i] * nums[j] * nums[k])
                }
            }
        }
    }
}
