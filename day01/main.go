package main

import "fmt"
import "sort"
import "bufio"
import "strings"
import "os"
import "strconv"

func main() {
  var leftList []int
  var rightList []int
  total := 0

  file, err := os.Open("./data.txt")
  if err != nil {
    fmt.Println("Error opening file:", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.Fields(line)

    if len(parts) != 2 {
      fmt.Println("Invalid line:", line)
      continue
    }


    left, err1 := strconv.Atoi(parts[0])
    right, err2 := strconv.Atoi(parts[1])

    if err1 != nil || err2 != nil {
      fmt.Println("Error converting line to integers:", line)
      continue
    }

    leftList = append(leftList, left)
    rightList = append(rightList, right)
  }


  fmt.Println("Left list:")
  fmt.Println(leftList)
  sort.Slice(leftList, func(i, j int) bool {
    return leftList[i] < leftList[j]
  })
  fmt.Println("Sorted Left list:")
  fmt.Println(leftList)
  fmt.Println("Right list:")
  fmt.Println(rightList)
  sort.Slice(rightList, func(i, j int) bool {
    return rightList[i] < rightList[j]
  })
  fmt.Println("Sorted Right list:")
  fmt.Println(rightList)

  if len(leftList) != len(rightList) {
    fmt.Println("List must have the same lenght")
    return
  }
  for i := 0; i < len(leftList); i++ {
    left := int(leftList[i])
    right := int(rightList[i])

    var result int
    if left > right {
      result = left - right
    } else {
      result = right - left
    }

    total += result
    // Print the result for each pair
    fmt.Printf("Result for index %d: %d\n", i, result)
  }

  fmt.Println(total)

}
