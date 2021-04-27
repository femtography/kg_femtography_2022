package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args[1:]
  r_val := ""
  counter := 0;
  hashy := map[string]string {
    "0" : "Zero",
    "1" : "One",
    "2" : "Two",
    "3" : "Three",
    "4" : "Four",
    "5" : "Five",
    "6" : "Six",
    "7" : "Seven",
    "8" : "Eight",
    "9" : "Nine",
  }
  fmt.Println(args)
}
