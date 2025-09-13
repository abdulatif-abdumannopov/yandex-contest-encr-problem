package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var quantity int
var data []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		firstLine := scanner.Text()
		val, err := strconv.Atoi(firstLine)
		if err != nil {
			fmt.Println("Invalid quantity")
		}
		quantity = val
	}

	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	if err := scanner.Err(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, "reading standard input:", err)
		if err != nil {
			return
		}
	}

	if quantity != len(data) {
		log.Fatal("Quantity is not equal number of data")
	}

	for i := 0; i < quantity; i++ {
		parts := strings.Split(data[i], ",")
		if len(parts) < 5 {
			fmt.Println("Некорректная строка")
			return
		}
		surname := parts[0]
		name := parts[1]
		patronymic := parts[2]

		day, _ := strconv.Atoi(parts[3])
		month, _ := strconv.Atoi(parts[4])

		sum := sumDigits(day) + sumDigits(month)

		fio := surname + name + patronymic

		unique := make(map[rune]bool)
		for _, r := range fio {
			unique[r] = true
		}

		letters := make([]rune, 0, len(unique))
		for r := range unique {
			letters = append(letters, r)
		}
		idx := letterIndex(rune(surname[0]))
		enc := len(letters) + sum*64 + idx*256
		hexStr := fmt.Sprintf("%X", enc)
		fmt.Println(hexStr[len(hexStr)-3:])
	}
}

func sumDigits(n int) int {
	sum := 0
	if n < 0 {
		n = -n
	}
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func letterIndex(ch rune) int {
	upper := unicode.ToUpper(ch)
	if upper >= 'A' && upper <= 'Z' {
		return int(upper-'A') + 1
	}
	return 0
}
