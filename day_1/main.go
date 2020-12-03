package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const (
	inputPath = "day_1/input.txt"
	target    = 2020
)

func main() {
	expenses, err := loadExpenses()
	if err != nil {
		log.Fatal(err)
	}

	var seen IntSet
	for _, expense := range expenses {
		seen.Add(expense)
	}

	for _, expense := range expenses {
		diff := target - expense
		if seen.Has(diff) {
			log.Println(expense * diff)
			break
		}
	}
}

func loadExpenses() ([]int, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var expenses []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return expenses, nil
}

type IntSet struct {
	words []uint32
}

func (s *IntSet) Add(i int) {
	word, bit := i/32, uint(i%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) Has(i int) bool {
	word, bit := i/32, uint(i%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}
