package main

func UpdateQuality(items []Item) {
	for i := range items {
		itm := items[i]
		itm.UpdateItem()
	}
}
