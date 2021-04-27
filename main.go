package main

import (
  "fmt"
  "os"
)

// Main function that takes in arguements from command line and returns phonetic conversion.
func main() {
  // Creating three variables. A list of arguements provided. A string for stdout.
  // A counter to trigger a stop the "while" loop. As well as the hashmap to convert numerical values.
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

  // While loop to go through arguements provided from command line. It will use hash table to convert each value
  // provided into the spelling of the number and then add it to r_val string.
  for counter < len(args) {
    item := []rune(args[counter])

    for i := 0; i < len(item); i++ {
      r_val = r_val + hashy[string(item[i])]
    }

    // Adds comma after each value in args but only if it is not the last item in list.
    if counter != len(args) - 1 {
      r_val = r_val + ","
    }

    counter ++;
  }
  // Returning expected value.
  fmt.Println(r_val)
}
