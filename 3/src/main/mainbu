package main 
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

/* Structure to save the data from the individual fabric pieces. 
 * Integers: data from text file.
 * dirty: if overlaps with other fabric, dirty is true.
 */
type fabric struct {
	id, offset_x ,offset_y, width, height int
	dirty bool
}

/* Structure to save the coordinates. 
 * x, y: coordinate of the struct.
 * value: empty->* one fabric in coordinate->fabric_id multiple->x
 */
type cord struct {
	x, y int
	value string
}

/* Size of square */
const size = 1000

func main() {
	/* slices for fabric data and cord data. */
	var data []fabric
	var cords []cord

	/* Open input file, close file when finished with program. */
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	/* Scanner to read input of file, append input to data with split_string function. */
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str_in := scanner.Text()
		data = append(data, split_string(str_in))
	}

	/* Execute functions to get result. */
	cords = fill_cord()
	insert_fabric(data, cords)
	set_dirty_bit(data, cords)
	defer fmt.Printf("Total overlap: %d\n", calc_overlap(cords))

	/* Search for fabric that the elves can use. */
	for i := 0 ; i < len(data); i++{
		if!data[i].dirty {
			fmt.Printf("Fabric to use: %d\n", data[i].id)
		}
	}
	//print_field(cords)

}
/* Set dirty if the fabric reaches a cord with a "x". */
func set_dirty_bit(data []fabric, cords []cord) {
	for i := 0; i < len(data); i++{
		/* Check if a inch of fabric has overlap in height * width */
		for j:= 0; j < data[i].height; j++ {
			for k:=0; k < data[i].width; k++ {
				offset := data[i].offset_x + data[i].offset_y * size + j * size + k
				if cords[offset].value == "x" {
					data[i].dirty = true
				}
			}
		}
	}
}

/* Increment for every cord with a x (overlap). */
func calc_overlap(cords []cord) int{
	var total int
	for i := 0 ; i < len(cords); i++ {
		if cords[i].value == "x" {
			total++
		}
	}
	return total
}

/* Insert fabric in cords */
func insert_fabric(data []fabric, cords []cord){
	for i := 0; i < len(data); i++{
		/* data.height and width is used for the rectangle. */
		for j:= 0; j < data[i].height; j++ {
			for k:=0; k < data[i].width; k++ {
				/* Offset is location on the grid. */
				offset := data[i].offset_x + data[i].offset_y * size + j * size + k
				/* If cords.value is not *, it is not empty, so duplicate, set "x". */
				if cords[offset].value != "*" {
					cords[offset].value = "x"
				} else {
					/* cord is empty, put ID there. */
					cords[offset].value = strconv.Itoa(data[i].id)
				}
			}
		}
	}
}

/* Init coordinate values. */
func fill_cord() []cord {
	var cords []cord

	for i := 0; i < size; i++ {
		for j:= 0; j < size; j++ {
			var temp cord
			temp.x = j
			temp.y = i
			temp.value = "*"
			cords = append(cords, temp)
		}
	}
	return cords
}

/* Split input file and add the correct values to data. */
func split_string(str string) fabric {
	var data fabric

	/* Split full string in smaller pieces.
	 * [#1 @ 53,238: 26x24] -> [#1] [@] [53,238:] [26x24] */
	f := strings.Split(str, " ")

	/* Remove # from the first, so id is left. Store in data.id */
	id := strings.Split(string(f[0]), "#")
	data.id, _ = strconv.Atoi(id[1])

	/* Remove "," and ":", so x and y offset is left. */
	offset_x := strings.Split(string(f[2]), ",")
	offset_y := strings.Split(string(offset_x[1]), ":")
	data.offset_x, _ = strconv.Atoi(offset_x[0])
	data.offset_y, _ = strconv.Atoi(offset_y[0])

	/* Remove "x", so width and height is left. */
	size := strings.Split(string(f[3]), "x")
	data.width, _ = strconv.Atoi(size[0])
	data.height, _ = strconv.Atoi(size[1])

	/* Return filled data */
	return data
}

/* Current input to big, with data from example it is possible to show the grid with this function, including the fabric. */
func print_field(cords []cord) {
	for i := 0; i < size; i++ {
		for j:= 0; j < size; j++{
			fmt.Print(cords[i * size + j].value)
			if j == size - 1 {
				fmt.Printf("\n")
			}
		}
	}
}

/* Error handling. */
func check(e error){
	if e != nil {
		panic(e)
	}
}
