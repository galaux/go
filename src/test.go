package main

import (
  "fmt"
//  "os"
)

var Max_depth = 10
var Depth     =  0

func getPivotIdx(list []int, start int, end int) int {
  return 0
}

// start: index of list to start partitionning (included)
// end  : index of list to end partitionning   (excluded)
func partition(list []int, start int, end int) ([]int, int) {
  if end - start == 1 {
    return list, start
  }

  i := start
  for j := 1; j < end; j++ {
    if list[j] < list[0] {
      // swap i and j
      tmp := list[i]
      list[i] = list[j]
      list[j] = tmp

      i++
    }
  }

  newPivotIdx := i - 1
  tmp := list[0]
  list[0] = list[newPivotIdx]
  list[newPivotIdx] = tmp

  return list, newPivotIdx
}

// start: index of list on which to start quicksort (included)
// end  : index of list on which to end quicksort   (included)
func quicksort(list []int, start int, end int) []int {
/*
  Depth++
  if Depth > Max_depth {
    os.Exit(1)
  }
  */
  fmt.Println("quicksort", list, ",", start, ",", end)
  pivotIdx := getPivotIdx(list, start, end)
  fmt.Printf("Pivot: list[%v] = %v\n", pivotIdx, list[pivotIdx])

  // Swap pivot for first position
  tmp := list[start]
  list[start] = list[pivotIdx]
  list[pivotIdx] = tmp

  partList, newPivotIdx := partition(list, 1, len(list))
  if start < newPivotIdx {
    quicksort(list, start, newPivotIdx - 1)
  }
  if newPivotIdx < end {
    quicksort(list, newPivotIdx + 1, end)
  }

  return partList
}

func main() {
  fmt.Println("Starting")
  list := []int {4, 3, 2, 6, 1, 5}
  fmt.Println(quicksort(list, 0, len(list) - 1))
  fmt.Println("Done")
}
