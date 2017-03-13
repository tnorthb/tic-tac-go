package main

import (
  "fmt"
  "math/rand"
  "time"
)

type Player struct {
    name string
    mark string
}

type Coordinate struct {
    row int
    col int
}

var Board [3][3]string
var currentPlayer Player

func main() {
  rand.Seed(time.Now().Unix())
  player1 := Player{"Tom", "O"}
  player2 := Player{"Eli", "X"}

  currentPlayer = player1
  last_move := Coordinate{0, 0}

  for has_won(last_move) == false && len(available_spaces()) != 0 {
    possible_moves := available_spaces()
    n := rand.Int() % len(possible_moves)
    chosen_move := possible_moves[n]
    move(currentPlayer, chosen_move)
    print_board()

    last_move = chosen_move
    if has_won(last_move) {
      fmt.Println(currentPlayer.name + " WINS!")
      return
    }
    switch currentPlayer {
    case player1:
      currentPlayer = player2
    case player2:
      currentPlayer = player1
    }
  }
  fmt.Println("DRAW!")
}

// Returns an array of length 3 with all values in that direction
func traverse_board(direction string, start Coordinate) []string {
  var sli []string
  if direction == "vertical" {
    for i := 0; i < len(Board); i++ {
      sli = append(sli, Board[i][start.col])
    }
  }

  if direction == "horizontal" {
    for i := 0; i < len(Board); i++ {
      sli = append(sli, Board[start.row][i])
    }
  }

  if direction == "diagonal1" {
    for i := 0; i < len(Board); i++ {
      sli = append(sli, Board[i][i])
    }
  }

  if direction == "diagonal2" {
    for i := 0; i < len(Board); i++ {
      sli = append(sli, Board[i][-(i - 2)])
    }
  }
  return sli
}

func all_equal(a []string) bool {
    for i := 0; i < len(a); i++ {
        if a[i] != a[0] || len(a[i]) == 0 {
            return false
        }
    }
    return true
}

func print_board() {
  for _,row := range Board {
    for i, v := range row {
      if len(v) == 0 {
        row[i] = "  "
      }
    }

    fmt.Println(row[0] + " | " + row[1] + " | " + row[2])
    fmt.Println("-----------")
  }
}

func available_spaces() []Coordinate {
  var sli []Coordinate
  for i,row := range Board {
    for col, v := range row {
      if len(v) == 0 {
        sli = append(sli, Coordinate{i, col})
      }
    }
  }
  return sli
}

func move(player Player, spot Coordinate) {
  Board[spot.row][spot.col] = player.mark
}

func has_won(move Coordinate) bool {
  if all_equal(traverse_board("vertical", move)) {
    return true
  } else if all_equal(traverse_board("horizontal", move)) {
    return true
  } else if all_equal(traverse_board("diagonal1", move)) {
    return true
  } else if all_equal(traverse_board("diagonal2", move)) {
    return true
  } else {
    return false
  }
}