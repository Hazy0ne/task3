package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scan1 int
var scan2 int
var slise = []int{}
var max int

func main() {
	rider := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ведите первое число:")
		rider.Scan()
		read, _ := strconv.Atoi(rider.Text())
		scan1 = read
		fmt.Println("Ведите второе число")
		rider.Scan()
		read, _ = strconv.Atoi(rider.Text())
		scan2 = read
		for i := scan1; i <= scan2; i++ {
			slise = append(slise, i)
		}

		for _, i := range slise {
			max = i
		}

		start()
		slise = slise[:0]
	}

}

func start() {
	sliseCopy := make([]int, len(slise))
	copy(sliseCopy, slise)
	for i := 2; i <= max; i++ {
		for in, ve := range sliseCopy {
			if ve%i == 0 && ve != i {
				slise[in] = 0
			}
		}
	}
	result := []int{}
	for _, v := range slise {
		if v > 1 {
			result = append(result, v)
		}
	}
	fmt.Println(result)

}
