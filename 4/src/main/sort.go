package main

import (
	"os"
	"io"
)


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
	write_file(sorted)
	return sorted
}

func write_file(data []input) {
	sorted, _ := os.Create(sort_input)
	defer sorted.Close()

	for i:= 0; i < len(data); i++ {
		str := data[i].full + "\n"
		io.WriteString(sorted, str)
	}
}




