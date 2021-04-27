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

  for counter < len(args) {
    item := []rune(args[counter])

    for i := 0; i < len(item); i++ {
      r_val = r_val + hashy[string(item[i])]
    }

    if counter != len(args) - 1 {
      r_val = r_val + ","
    }

    counter ++;
  }
  fmt.Println(r_val)
}
