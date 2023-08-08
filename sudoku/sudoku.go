package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// Boardは数独の盤面を表す
//
// 0 - 未入力
// 1-9 - 埋まっている
type Board [9][9]int

func pretty(b Board) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa((b[i][j])))
		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}

func isDuplicated(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(b Board) bool {
	// 行チェック
	for i := 0; i < 9; i++ {
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if isDuplicated(c) {
			return false
		}
	}

	// 列チェック
	for i := 0; i < 9; i++ {
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
		}
		if isDuplicated(c) {
			return false
		}
	}

	// 3x3のチェック
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if isDuplicated(c) {
				return false
			}
		}
	}
	return true
}

func isSolved(b Board) bool {
	zeroExists := false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				zeroExists = true
			}
		}
	}
	return !zeroExists && verify(b)
}

func backtrack(b *Board) bool {
	pretty(*b)
	time.Sleep(time.Second * 1)
	fmt.Printf("%+v", pretty(*b))
	if isSolved(*b) {
		fmt.Println("Solved")
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for cand := 1; cand <= 9; cand++ {
					b[i][j] = cand
					if verify(*b) {
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func main() {
	b := Board{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	fmt.Printf("%+v\n", pretty(b))
	if backtrack(&b) {
		fmt.Printf("%+v\n", pretty(b))
	}
}
