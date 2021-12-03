package main

import "testing"

func Test_Foo(t *testing.T) {
	sulfura := &Sulfuras{item: item{"foo", 0, 0}}
	var items = []Item{sulfura}

	UpdateQuality(items)

	if items[0].GetName() != "foo" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].GetName())
	}
}
