package main

import (
	"fmt"
	"strconv"
	"strings"
	"test-be/entities"
	"test-be/services"
)

var (
	ser     = services.MaxService{Array: entities.Array{}}
	strList string
)

func main() {
	fmt.Print("Input List: ")
	fmt.Scanln(&strList)

	tempList := strings.Split(strList, ",")
	for i := range tempList {
		num, _ := strconv.Atoi(tempList[i])
		ser.Array.List = append(ser.Array.List, num)
	}
	fmt.Println("Max value:", ser.FindMax(&ser.Array))
}
