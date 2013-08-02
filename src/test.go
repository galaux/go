package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
  //"os"
  "time"
)

var Max_depth = 10
var Depth     =  0

//func getPivotIdx(list []int, start int, end int) int {
//  return 0
//}

func getPivotIdx(list *[]int, start int, end int) int {
  return start
}

// start: index of list to start partitionning (included)
// end  : index of list to end partitionning   (included)
func partition(list *[]int, start int, end int) int {
  i := start + 1
  if end != start {
    for j := start + 1; j <= end; j++ {
      if (*list)[j] < (*list)[start] {
        // swap i and j
        tmp       := (*list)[i]
        (*list)[i] = (*list)[j]
        (*list)[j] = tmp

        i++
      }
    }
  }

  newPivotIdx := i - 1
  tmp := (*list)[start]
  (*list)[start] = (*list)[newPivotIdx]
  (*list)[newPivotIdx] = tmp

  return newPivotIdx
}

// start: index of list on which to start quicksort (included)
// end  : index of list on which to end quicksort   (included)
func quicksort(list *[]int, start int, end int) {
  if start >= end {
    return
  }

  pivotIdx := getPivotIdx(list, start, end)

  // Swap pivot for first position
  tmp := (*list)[start]
  (*list)[start] = (*list)[pivotIdx]
  (*list)[pivotIdx] = tmp

  newPivotIdx := partition(list, start, end)
  quicksort(list, start, newPivotIdx - 1)
  quicksort(list, newPivotIdx + 1, end)
}

func readFile(fname string) (nums []int, err error) {
  b, err := ioutil.ReadFile(fname)
  lines := strings.Split(string(b), "\r\n")
  // Assign cap to avoid resize on every append.
  nums = make([]int, 0, len(lines))

  for _, l := range lines {
    // Empty line occurs at the end of the file when we use Split.
    if len(l) == 0 { continue }
    // Atoi better suits the job when we know exactly what we're dealing
    // with. Scanf is the more general option.
    n, err := strconv.Atoi(l)
            if err != nil { return nil, err }
    nums = append(nums, n)
  }

  return nums, nil
}

func main() {
  fmt.Println("Starting")

  list, err := readFile("./QuickSort.txt")
  if err != nil { panic(err) }
  //list := []int{3, 5, 4, 1, 6, 7, 2}
  fmt.Println("Found", len(list), "integers")

  t0 := time.Now()
  quicksort(&list, 0, len(list) - 1)
  t1 := time.Now()
  fmt.Println("Done. FYI, the resulting array has size", len(list))
  fmt.Println(list)
  fmt.Printf("The call took %v to run.\n", t1.Sub(t0))


  fmt.Println("Done")
}
