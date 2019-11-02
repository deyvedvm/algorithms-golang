package main

import (
	"container/list"
	"testing"
)

func TestList(t *testing.T) {
	var mockNumber = 11
	var integerList list.List

	element := integerList.PushBack(mockNumber)

	if element.Value != mockNumber {
		t.Errorf("Expected %d, but got %d", mockNumber, element.Value)
	}
}
