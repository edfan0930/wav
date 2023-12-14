package playwav

import (
	"strconv"
	"strings"
)

func HokkienRules(callNumber, counterNumber int) []string {

	n := strconv.Itoa(callNumber)
	c := strconv.Itoa(counterNumber)
	callNumberSlice := strings.Split(n, "")
	counterNumberSlice := strings.Split(c, "")
	all := []string{"來賓"}

	PutMiddle(&callNumberSlice)
	PutMiddle(&counterNumberSlice)
	all = append(all, callNumberSlice...)

	all = append(all, AfterNumber(counterNumberSlice)...)
	//	Put(HokkienPath, all)
	PutWAVExtension(all)
	return all
}

// AfterNumber
func AfterNumber(counter []string) []string {

	return append([]string{"號", "請到"}, append(counter, []string{"號", "櫃台"}...)...)
}

// PutMiddle
func PutMiddle(n *[]string) {

	number := *n
	length := len(number)
	//the length great than and equal 2, and
	if (length >= 2) && (number[length-2] != "0") {
		if number[length-2] == "1" {
			number[length-2] = "拾"
		} else {

			number = append(number[:length-1], append([]string{"拾"}, number[length-1:]...)...)
		}
	}

	if length >= 3 && (number[length-3] != "0") {

		number = append(number[:length-2], append([]string{"佰"}, number[length-2:]...)...)
	}

	if length == 4 {

		number = append(number[:length-3], append([]string{"仟"}, number[length-3:]...)...)
	}

	after := len(number)
	if number[after-1] == "0" && number[after-2] == "0" && number[after-3] == "0" {
		number = number[:after-3]
	} else if number[after-1] == "0" && number[after-2] == "0" {
		number = number[:after-2]

	} else if number[after-1] == "0" {
		number = number[:after-1]
	}

	*n = number
}
