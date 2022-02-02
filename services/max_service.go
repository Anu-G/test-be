package services

import (
	"reflect"
	"test-be/entities"
)

type MaxService struct {
	Array entities.Array
}

func (M MaxService) FindSorted(num *entities.Array) []int {
	arr := append(M.Array.List, 0)
	tempArr2D := [][]int{}

	tempArr := []int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i]+1 == arr[i+1] {
			tempArr = append(tempArr, arr[i])
		} else if i == 0 {
			continue
		} else if arr[i]-1 == arr[i-1] {
			tempArr = append(tempArr, arr[i])

			tempArr2D = append(tempArr2D, tempArr)
		}
	}

	list := append([]int{0}, M.Array.List...)
	tempArr2DRev := [][]int{}

	tempArrRev := []int{}
	for i := len(list) - 1; i > 0; i-- {
		if list[i]+1 == list[i-1] {
			tempArrRev = append(tempArrRev, list[i])
		} else if i == len(list)-1 {
			continue
		} else if list[i]-1 == list[i+1] {
			tempArrRev = append(tempArrRev, list[i])

			tempArr2DRev = append(tempArr2DRev, tempArrRev)
		}
	}

	resArr := []int{}
	for i := range tempArr2D {
		for j := range tempArr2DRev {
			if reflect.DeepEqual(tempArr2D[i], tempArr2DRev[j]) {
				resArr = tempArr2D[i]
			}
		}
	}

	return resArr
}

func (M MaxService) FindMax(num *entities.Array) int {
	resArr := M.FindSorted(num)
	var tempMax int

	for i := 0; i < len(resArr)-1; i++ {
		if resArr[i] < resArr[i+1] {
			tempMax = resArr[i+1]
		}
	}

	return tempMax
}
