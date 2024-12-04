package main

import "fmt"
import "bufio"
import "strings"
import "os"
import "strconv"

func main() {
  var leftList []int
  var rightList []int
  var total int
  total = 0

  // file, err := os.Open("./sample.txt")
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
  fmt.Println("Right list:")
  fmt.Println(rightList)

  for _, left := range leftList{

    count := 0
    for _, right := range rightList{
      if left == right {
        count += 1
      }
    }
    factorTotal := left * count
    total += factorTotal
    fmt.Printf("Factor %d is found %d times for a similarity score of: %d",left, count, factorTotal)
    fmt.Println("")
    fmt.Printf("Current similarity factor: %d", total)
    fmt.Println("")

  }

}
