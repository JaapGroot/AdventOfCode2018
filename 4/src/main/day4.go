package main

import (
	//"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"io"
)

type input struct {
	date, time, full string
	value int
}


const raw_input = "input/inputraw"
const sort_input = "input/inputsort"

func main() {
	var inputraw []input
	var inputsorted []input

	/* Get the raw data and fill raw input */
	inputraw = fill_inputraw()
	/* Sort the list with a quicksort algorithm (thanks Weiss) */
	inputsorted = sort(inputraw)

	sorted, _ := os.Create(sort_input)
	defer sorted.Close()
	for i := 0; i < len(inputsorted); i++ {
		str := inputsorted[i].full + "\n"
		io.WriteString(sorted, str)
	}
}

func sort(inputraw []input) []input {
	var sorted []input
	if len(inputraw) > 1 {
		var smaller, same, larger []input

		chosenItem := inputraw[len(inputraw) / 2]
		for _, i := range inputraw {
			if i.value < chosenItem.value {
				smaller = append(smaller, i)
			} else if i.value > chosenItem.value {
				larger = append(larger, i)
			} else {
				same = append(same,i)
			}
		}

		smaller = sort(smaller)
		larger = sort(larger)

		sorted = append(sorted,smaller...)
		sorted = append(sorted,same...)
		sorted = append(sorted,larger...)
	} else {
		sorted = append(sorted,inputraw...)
	}
	return sorted
}

func fill_inputraw() []input {
	var inputfill []input
	rawfile, _ := os.Open(raw_input)
	defer rawfile.Close()

	scanner := bufio.NewScanner(rawfile)
	for scanner.Scan() {//i := 0; i < 5; i++ {//scanner.Scan() {
		var temp input
		//scanner.Scan()
		//var raw input
		str := scanner.Text()
		temp.full = str
		split := strings.FieldsFunc(str, func(r rune) bool {
			switch r {
				case '[', ']', ' ':
					return true
				}
			return false
		})
		temp.date = split[0]
		temp.time = split[1]
		temp.value = calc_value(temp)
		inputfill = append(inputfill, temp)
	}


	return inputfill
}

func calc_value(data input) int {
	var value int

	split := strings.Split(data.date, "-")
	year, _  := strconv.Atoi(split[0])
	month, _ := strconv.Atoi(split[1])
	day, _ := strconv.Atoi(split[2])

	split = strings.Split(data.time, ":")
	hour, _ := strconv.Atoi(split[0])
	minute, _ := strconv.Atoi(split[1])

	value = ((year * 12 * 30) + (month * 30) + day) * 24 * 60 + hour * 60 + minute

	return value
}

