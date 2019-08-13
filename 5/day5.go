package main 
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

const example = "input/example"
const full = "input/input"
const dir = full

func main() {

	file, _ := os.Open(dir)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Can't read file!")
		return
	}
	input := scanner.Text()
	input = strings.TrimSpace(input)
	smallest := input
	for i:= 65; i < 91; i++{
		temp := input
		temp = removeChar(temp, byte(i))
		temp = checkLetters(temp)
		if len(temp) < len(smallest) {
			smallest = temp
		}
	}
	fmt.Println(smallest)

	//s := checkLetters(input)
	fmt.Println(len([]rune(smallest)))
}

func checkLetters(s string) string {
	for i := 0; i < len(s) - 1; {
		removed := false
		if s[i] < 95 { // Capital
			if s[i] == s[i + 1] - 32 {
				s = shiftString(s, i)
				removed = true
			}
		} else if s[i] > 95 { // Small
			if s[i] == (s[i + 1] + 32) {
				s = shiftString(s, i)
				removed = true
			}
		}
		if removed {
			i = 0
		}else {
			i++
		}
	}
	return s
}

func removeChar(s string, char byte) string{
	new_str := []byte(s)
	for i := 0; i < len(new_str); {
		dirty := false
		if new_str[i] == char || new_str[i] == char + 32{
			new_str = append(new_str[:i], new_str[i+1:]...)
			dirty = true
		}
		if !dirty {
			i++
		}
	}
	return string(new_str)
}

func shiftString(s string, i int) string{
	new_str := []byte(s)
	new_str = append(new_str[:i], new_str[i+2:]...)
	return string(new_str)
}

