package main

import "fmt"

func isSet(board1 []int, board2 []int, board3 []int) bool {
  counter := 0
  for i:=0; i<4; i++ {
    total := board1[i] + board2[i] + board3[i]
    if (total == 0 || total == 3 || total == -3){
      counter = counter + 1
    }
  }

  if (counter == 4) {
    return true
  }

  return false
}

func identifySet(setsBoard [][]int){
  var identifySet bool;
  for i:=0; i<len(setsBoard); i++ {
    for j:=0; j<len(setsBoard); j++ {
      for k:=0; k<len(setsBoard); k++ {
        if (i != j && j != k && i != k){
          identifySet = isSet(setsBoard[i], setsBoard[j], setsBoard[k])
          if (identifySet == true){
            fmt.Println(setsBoard[i], setsBoard[j], setsBoard[k])
          }
        }
      }
    }
  }
}

func main() {
  setsBoard := [][]int{{0,1,0,0}, {0,1,0,0}, {0,1,0,0}, {0,0,0,0}, {0,0,0,0}, {0,0,0,0}, {0,0,0,0}, {0,0,0,0}, {0,0,0,0}, {0,0,0,0}, {0,0,0,0}, {0,0,0,0}}
  identifySet(setsBoard)
}
