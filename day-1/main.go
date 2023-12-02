package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addsum(ans []string) (sum int) {
	var temp string
	for _, val := range ans {

		for _, char := range val {
			if char == '1' || char == '2' || char == '3' || char == '4' || char == '5' || char == '6' || char == '7' || char == '8' || char == '9' {
				temp += string(char)
			}
		}
		if len(temp) > 0 {
			temp = string(temp[0]) + string(temp[len(temp)-1])
			value, _ := strconv.Atoi(temp)

			sum += value
		}

		temp = ""
	}

	return
}

func main() {
	data, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println(err)
	}
	ans := strings.Split(string(data), "\n")

	sum := addsum(ans)
	println(sum)

}
