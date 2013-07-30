package main

import (
  "fmt"
)

func getPivotIdx(list []int) int {
  return 0
}

func partition(list []int) []int {
  i := 1
  for j := 1; j < len(list); j++ {
    if list[j] < list[0] {
      // swap i and j
      tmp := list[i]
      list[i] = list[j]
      list[j] = tmp

      i++
    }
  }

  tmp := list[0]
  list[0] = list[i - 1]
  list[i - 1] = tmp

  return list
}

func quicksort(list []int) []int {
  pivotIdx := getPivotIdx(list)
  fmt.Printf("Pivot: list[%v] = %v\n", pivotIdx, list[pivotIdx])

  // Swap pivot for first position
  tmp := list[0]
  list[0] = list[pivotIdx]
  list[pivotIdx] = tmp

  partList := partition(list)

  return partList
}

func main() {
  fmt.Println("Starting")
  list := []int {4, 3, 2, 6, 1, 5}
  fmt.Println(quicksort(list))
  fmt.Println("Done")
}
