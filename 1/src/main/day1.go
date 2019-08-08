package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main() {
	var sum int
	var s []int
	var total int
	var first_loop = true
	var found = false

	/* set first value to 0. */
	s = append(s, 0)
	file, err := os.Open("../../input")
	check(err)
	defer file.Close()

	/* Loop till a duplicate is found. */
	for found == false{
		/* Goto begin of file. */
		file.Seek(0,0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			/* Store line in res. */
			res, _ := strconv.Atoi(scanner.Text())
			sum += res
			/* For every element in s(um) slice. */
			for _, element := range s {
				/* If a duplicate is found, break. */
				if element == sum {
					found = true
					break
				}
			}
			/* Also break from this loop. */
			if found == true{
				break
			}
			s = append(s, sum)
		}
		/* While in the first loop, calculate total of all values. */
		if first_loop {
			total = sum
			first_loop = false
		}
	}
	/* Print results. */
	fmt.Printf("Total sum: %d\nFirst duplicate: %d\n", total, sum)
}
