package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestExist(t *testing.T) {
	tests := []struct {
		board  [][]byte
		word   string
		expect bool
	}{
		{
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE"),
			},
			word:   "ABCCED",
			expect: true,
		},
		{
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE"),
			},
			word:   "SEE",
			expect: true,
		},
		{
			board: [][]byte{
				[]byte("ABCE"),
				[]byte("SFCS"),
				[]byte("ADEE"),
			},
			word:   "ABCB",
			expect: false,
		},
		{
			board: [][]byte{
				[]byte("aa"),
			},
			word:   "aaa",
			expect: false,
		},
		{
			board: [][]byte{
				[]byte("aaaa"),
				[]byte("aaaa"),
				[]byte("aaaa"),
			},
			word:   "aaaaaaaaaaaaa",
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if r := exist(tt.board, tt.word); r != tt.expect {
				t.Errorf("expect: %v, actual: %v, word: %v", tt.expect, r, tt.word)
			}
		})
	}
}

func exist(board [][]byte, word string) bool {
	c := make(map[string]int)
	for _, row := range board {
		for _, cell := range row {
			c[string(cell)]++
		}
	}
	w := make(map[string]int)
	for _, ww := range word {
		w[string(ww)]++
	}
	for ww, wc := range w {
		if cc, ok := c[ww]; !ok {
			return false
		} else if cc < wc {
			return false
		}
	}

	for i, row := range board {
		for j := range row {
			if r := search(j, i, board, word); r {
				return r
			}
		}
	}
	return false
}

func search(x, y int, board [][]byte, word string) bool {
	if board[y][x] != word[0] {
		return false
	}
	m := make(map[string]bool)
	var sb strings.Builder
	sb.WriteByte(word[0])
	return back(x, y, board, word, &sb, m)
}

func back(x, y int, board [][]byte, word string, cur *strings.Builder, dir map[string]bool) bool {
	xy := strconv.Itoa(x) + strconv.Itoa(y)
	if ok := dir[xy]; ok {
		return false
	}
	s := cur.String()
	if s == word {
		return true
	}
	if s != word[:len(s)] {
		return false
	}
	dir[xy] = true
	if y < len(board)-1 {
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteByte(board[y+1][x])
		if r := back(x, y+1, board, word, &sb, dir); r {
			return true
		}
	}
	if y > 0 {
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteByte(board[y-1][x])
		if r := back(x, y-1, board, word, &sb, dir); r {
			return true
		}
	}
	if x < len(board[0])-1 {
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteByte(board[y][x+1])
		if r := back(x+1, y, board, word, &sb, dir); r {
			return true
		}
	}
	if x > 0 {
		var sb strings.Builder
		sb.WriteString(s)
		sb.WriteByte(board[y][x-1])
		if r := back(x-1, y, board, word, &sb, dir); r {
			return true
		}
	}
	dir[xy] = false
	return false
}
