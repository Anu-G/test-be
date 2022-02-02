package services

import (
	"reflect"
	"test-be/entities"
	"testing"
)

var (
	anArray      entities.Array = entities.Array{}
	expectedList []int
	expectedMax  int
)

func Test_FindSorted(t *testing.T) {
	anArray.List = []int{1, 2, 3, 8, 9, 3, 2, 1}
	expectedList = []int{1, 2, 3}

	ser := &MaxService{Array: anArray}
	res := ser.FindSorted(&anArray)

	if !reflect.DeepEqual(res, expectedList) {
		t.Errorf("Wrong Return=%v, Expected=%v", res, expectedList)
	}
}

func Test_FindMax(t *testing.T) {
	anArray.List = []int{1, 2, 3, 8, 9, 3, 2, 1}
	expectedMax = 3

	ser := &MaxService{Array: anArray}
	res := ser.FindMax(&anArray)

	if !reflect.DeepEqual(res, expectedMax) {
		t.Errorf("Wrong Return=%v, Expected=%v", res, expectedMax)
	}
}
