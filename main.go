package main

import (
  "fmt"
  "math/rand"
)

func pickRandomCard(setsBoard [][4]int) ([4]int, int) {
  r := rand.Intn(81)
  fmt.Println("random number", r)

  return setsBoard[r], r
}

func cardsAreEqual(card1 [4]int, card2 [4]int) bool {
  return card1 == card2
}

func addCard(setsBoard [][4]int, card [4]int) [][4]int {
  return append(setsBoard, card)
}

func generateAllCards() [][4]int {
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

func generateSetsBoard(unpickedCards [][4]int, setsBoard [][4]int) ([][4]int, []int) {
  numberOfCards := len(setsBoard)
  indicesOfCardsPicked := []int{}
  for i:=0; i<12-numberOfCards; i++ {
    card, index := pickRandomCard(unpickedCards)
    indicesOfCardsPicked = append(indicesOfCardsPicked, index)
    setsBoard = addCard(setsBoard, card)
  }

  return setsBoard, indicesOfCardsPicked
}

func removeCard(setsBoard [][4]int, i int) (newSetsBoard [][4]int, removedCard [4]int) {
  card := setsBoard[i]
  setsBoard[i] = setsBoard[len(setsBoard) - 1]
  return setsBoard[:len(setsBoard) - 1], card
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

func identifySet(setsBoard [][]int) (int, int, int) {
  var identifySet bool;
  for i:=0; i<len(setsBoard); i++ {
    for j:=0; j<len(setsBoard); j++ {
      for k:=0; k<len(setsBoard); k++ {
        if (i != j && j != k && i != k){
          identifySet = isSet(setsBoard[i], setsBoard[j], setsBoard[k])
          if (identifySet == true){
            return i, j, k
          }
        }
      }
    }
  }
  return 0, 0, 0
}

func allColorsExist(setsBoard [][]int) bool {
  colorsMap := make(map[string]bool)
  for i:=0; i<len(setsBoard); i++ {
    if (setsBoard[i][0] == 1){
      colorsMap["blue"] = true
    }
    if (setsBoard[i][0] == 1){
      colorsMap["red"] = true
    }
    if (setsBoard[i][0] == 1){
      colorsMap["green"] = true
    }
  }

  return (colorsMap["blue"] && colorsMap["red"] && colorsMap["green"])
}

func allNumbersExist(setsBoard [][]int) bool {
  numbersMap := make(map[int]bool)
  for i:=0; i<len(setsBoard); i++ {
    if (setsBoard[i][1] == 1){
      numbersMap[1] = true
    }

    if (setsBoard[i][1] == 0){
      numbersMap[1] = true
    }

    if (setsBoard[i][1] == 1){
      numbersMap[1] = true
    }
  }
}

func main() {
  allCards := generateAllCards()
  fmt.Println(allCards)
  setsBoard, indicesOfCardsPicked := generateSetsBoard(allCards, [][4]int{})
  fmt.Println("setsBoard", setsBoard)
  fmt.Println("indicesOfCardsPicked", indicesOfCardsPicked)

}
