package main

import (
  "fmt"
  "math/rand"
)

func pickRandomCard(setsBoard [][4]int) [4]int {
  r := rand.Intn(81)
  fmt.Println("random number", r)

  return setsBoard[r]
}

func cardsAreEqual(card1 [4]int, card2 [4]int) bool {
  return card1 == card2
}

func addCard(setsBoard [][4]int, card [4]int) [][4]int {
  return append(setsBoard, card)
}

func generateSetsBoard() [][4]int {
  setsBoard, helperBoard := [][4]int{{0,0,0,0}}, [][4]int{}

  for j:=0; j<4; j++ {
    for i:=0; i<len(setsBoard); i++ {
      card1 := setsBoard[i]
      card1[j] = 1
      helperBoard = addCard(helperBoard, card1)

      card2 := setsBoard[i]
      card2[j] = -1
      helperBoard = addCard(helperBoard, card2)
    }

    setsBoard = append(setsBoard, helperBoard...)
    helperBoard = [][4]int{}
  }

  return setsBoard
}

func removeCard(setsBoard [][4]int, card [4]int) (newSetsBoard [][4]int, removedCard [4]int) {
  for i:=0; i<12; i++ {
    if (setsBoard[i] == card){
      setsBoard[i] = setsBoard[len(setsBoard) - 1]
      return setsBoard[:len(setsBoard) - 1], card
    }
  }

  return setsBoard, card
}

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

  board1 := generateSetsBoard()
  fmt.Println(board1)
  fmt.Println(len(board1))

  newBoard, cardRemoved := removeCard(board1, [4]int{0,0,0,0})
  fmt.Println("newBoard", newBoard)
  fmt.Println("cardRemoved", cardRemoved)

  randomCard := pickRandomCard(newBoard)
  fmt.Println("random card", randomCard)
}
