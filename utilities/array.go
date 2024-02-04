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

func SplitArrayToLowHighZone(arr []int) ([]int, []int) {
  var lowZone []int
  var highZone []int

  for _, num := range arr {
    if num >= 1 && num <= 3 {
      lowZone = append(lowZone, num)
    } else if num >= 4 && num <= 5 {
      highZone = append(highZone, num)
    }
  }

  return lowZone, highZone
}
