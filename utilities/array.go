package utilities

import (
  "fmt"
  "sort"
)

func SortAndRemoveDuplicates(arr []int) []int {
  sort.Ints(arr)

  var uniqueArr []int
  seen := make(map[int]bool)

  for _, item := range arr {
    if !seen[item] {
      seen[item] = true
      uniqueArr = append(uniqueArr, item)
    }
  }

  fmt.Printf("uniqueArr: %v", uniqueArr)

  return uniqueArr
}
