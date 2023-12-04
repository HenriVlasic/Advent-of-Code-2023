package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

var digits = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func main() {
  var finalSum int
  data,err := os.Open("./example")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer data.Close()

  scanner := bufio.NewScanner(data)
  for scanner.Scan() {
    line := scanner.Text()
    finalNumber := verifyStringForDigits(line)
    finalSum += finalNumber
  }

  if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

  fmt.Println("Final Sum:", finalSum)
}

func verifyStringForDigits(line string) int {
  var firstDigit rune
  var lastDigit rune

  firstIndex := -1

  for index, character := range line {
    for _, c := range digits {
      if character == c {
        if firstIndex < index && firstIndex == -1 {
          firstIndex = index
          firstDigit = character
        }
        lastDigit = character
      }
    }
  }
  joinedNumbers := fmt.Sprintf("%s%s", string(firstDigit), string(lastDigit))
  finalNumber, _ := strconv.Atoi(joinedNumbers)
  return finalNumber
}
