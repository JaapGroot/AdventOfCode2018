package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type input struct {
	date, time, action, full string
	value, id int
}

var zeroGuard = &guard{}
func (a *guard) Reset() {
	*a = *zeroGuard
}

type guard struct {
	id, sleeptime int
	action, time []string
	count [60]int
}
const raw_input = "input/inputraw"
const sort_input = "input/inputsort"
const small_input = "input/inputsmall"


func main() {
	var inputraw []input
	var inputsorted []input
	var guards []guard
	/* Get the raw data and fill raw input */
	inputraw = fill_inputraw()
	/* Sort the list with a quicksort algorithm (thanks Weiss) */
	inputsorted = sort(inputraw)
	guards = fill_guard(inputsorted)
	longest_sleeptime(guards)
}


func longest_sleeptime(guards []guard) {
	var highest, guard int

	for i := 0; i < len(guards); i++ {
		for j := i + 1; j < len(guards); j++ {
			if guards[i].id == guards[j].id {
				guards[i].sleeptime += guards[j].sleeptime
			}
		}
		if guards[i].sleeptime > highest {
			highest = guards[i].sleeptime
			guard = i
		}
	}

	for ref := 0; ref < len(guards); ref++ {
		for i := 0; i < len(guards); i++ {	
			if(guards[i].id == guards[ref].id) {
				for j := 0; j < 60; j++{
					if is_asleep(j, guards[i]){
						guards[ref].count[j]++
					}
				}
			}
		}
	}

	highest = 0
	var minute int
	for i := 0; i < 60; i++ {
		if guards[guard].count[i] > highest {
			highest = guards[guard].count[i]
			minute = i
		}
	}

	var most_asleep, most_id, most_min int

	for i := 0 ; i < len(guards); i++ {
		for j := 0 ; j < 60; j++ {
			if guards[i].count[j] > most_asleep {
				most_asleep = guards[i].count[j]
				most_min = j
				most_id = i
			}
		}
	}


	fmt.Printf("Guard[%d]. Sleeptime: %d. Favorite minut: %d. minut * guard: %d\n", guards[guard].id, guards[guard].sleeptime, minute ,guards[guard].id * minute)
	fmt.Printf("Guard[%d]. Number of sleeps: %d. At minut: %d. minut * guard: %d\n", guards[most_id].id, most_asleep, most_min, most_min * guards[most_id].id)
}

func set_sleeptime(guards []guard) {
	for i := 0 ; i < len(guards); i++ {
		var sleeptime int
		var starttime int
		for j := 0; j < len(guards[i].action); j++ {
			split := strings.Split(guards[i].time[j], ":")
			minute, _ := strconv.Atoi(split[1])
			hour, _ := strconv.Atoi(split[0])
			if hour == 0 {
				hour = 24
			}
			if j % 2 == 0 {
				/* Get falls asleep time */
				//fmt.Println(guards[i].time[j])
				starttime  =  minute + (hour * 60)
			} else {
				sleeptime += (minute + (hour * 60)) - starttime
			}
		}
		guards[i].sleeptime = sleeptime
	}

}

func is_asleep(time int, data guard) bool {
	var asleep bool
	var start, stop int

	for j := 0; j < len(data.action); j++ {
		split := strings.Split(data.time[j], ":")
		minute, _ := strconv.Atoi(split[1])
		if j % 2 == 0 {
			start = minute
		} else {
			stop = minute
		}
		if time >= start && time < stop {
			asleep = true
		}
	}
	return asleep
}

func fill_guard(data []input) []guard {
	var ret []guard
	var temp guard

	for i := 0; i < len(data); i++ {
		if data[i].id != 0{
			ret = append(ret, temp)
			temp.Reset()
			temp.id = data[i].id
			//fmt.Println(data[i].id)
		} else {
			temp.action = append(temp.action, data[i].action)
			temp.time = append(temp.time, data[i].time)
		}
	}
	/* Insert last one */
	_, ret = ret[0], ret[1:]
	ret = append(ret, temp)
	set_sleeptime(ret)
	return ret
}

func fill_inputraw() []input {
	var inputfill []input
	rawfile, _ := os.Open(raw_input)//(raw_input)
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
				case '[', ']', ' ', '#':
					return true
				}
			return false
		})
		temp.date = split[0]
		temp.time = split[1]
		if id, err := strconv.Atoi(split[3]); err == nil {
			temp.id = id
		} else {
			temp.action = split[2] + split [3]
		}
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
	value = ((year * 12 * 30) + (month * 31) + day) * 24 * 60 + hour * 60 + minute

	return value
}


