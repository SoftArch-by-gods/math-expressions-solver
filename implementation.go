package lab2

import (
	"errors"
	st "github.com/golang-collections/collections/stack"
	"math"
	"strconv"
	"strings"
)

func isDigit(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return false
	}
	return true
}

func strRevers(s string) []string {
	runes := strings.Split(s, " ")
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func CalculatePrefix(input string) (string, error) {
	s := st.New()
	str := strRevers(input)

	for _, elem := range str {
		if isDigit(elem) {
			num, _ := strconv.ParseFloat(elem, 64)
			s.Push(num)
		} else {
			val1 := s.Pop()
			val2 := s.Pop()
			if val1 == nil || val2 == nil {
				return "error", errors.New("expression syntax error")
			}
			n1 := val1.(float64)
			n2 := val2.(float64)

			switch elem {
			case "+":
				s.Push(n1 + n2)
				break
			case "-":
				s.Push(n1 - n2)
				break
			case "*":
				s.Push(n1 * n2)
				break
			case "/":
				s.Push(n1 / n2)
				break
			case "^":
				s.Push(math.Pow(n1, n2))
				break
			default:
				return "error", errors.New("undefined symbol")
			}
		}
	}
	if s.Len() != 1 {
		return "error", errors.New("expression syntax error")
	}
	res := s.Pop()
	fl := res.(float64)
	out := strconv.FormatFloat(fl, 'f', -1, 64)
	return out, nil
}
