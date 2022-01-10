package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func printItems(items []*Item) {
	for _, item := range items {
		fmt.Printf("%+v\n", *item)
	}
}

func compareItems(actualItems, expectedItems []*Item) error {
	if len(actualItems) != len(expectedItems) {
		return fmt.Errorf("expected %v items, got %v", len(expectedItems), len(actualItems))
	}
	for i, actual := range actualItems {
		expected := expectedItems[i]
		if diff := cmp.Diff(*actual, *expected, cmp.AllowUnexported(Item{})); diff != "" {
			return fmt.Errorf("Item mismatch (-want +got):\n%s", diff)
		}
	}
	return nil
}

type Test struct {
	name     string
	input    []*Item
	expected []*Item
}

var tests = []Test{
	// NAME
	{
		name: "item name not changed",
		input: []*Item{
			{name: "foo", sellIn: 0, quality: 0},
		},
		expected: []*Item{
			{name: "foo", sellIn: -1, quality: 0},
		},
	},

	// SELLIN & QUALITY
	{
		name: "item 'foo' 'sellIn' decremented",
		input: []*Item{
			{name: "foo", sellIn: 0, quality: 0},
		},
		expected: []*Item{
			{name: "foo", sellIn: -1, quality: 0},
		},
	},
	{
		name: "item  'foo' 'quality' decreases with 'sellIn'",
		input: []*Item{
			{name: "foo", sellIn: 10, quality: 10},
		},
		expected: []*Item{
			{name: "foo", sellIn: 9, quality: 9},
		},
	},
	{
		name: "item  'foo' 'quality' decreases by 2 for each update when 'sellin' < 0",
		input: []*Item{
			{name: "foo", sellIn: 0, quality: 10},
		},
		expected: []*Item{
			{name: "foo", sellIn: -1, quality: 8},
		},
	},
	{
		name: "item  'foo' 'quality' never negative",
		input: []*Item{
			{name: "foo", sellIn: 0, quality: 0},
		},
		expected: []*Item{
			{name: "foo", sellIn: -1, quality: 0},
		},
	},
	// special item: Aged Brie
	{
		name: "item 'Aged Brie' 'quality' increases as 'sellIn' decreases",
		input: []*Item{
			{name: "Aged Brie", sellIn: 5, quality: 5},
		},
		expected: []*Item{
			{name: "Aged Brie", sellIn: 4, quality: 6},
		},
	},
	{
		name: "item 'Aged Brie' 'quality' increases by 2 when 'sellIn' < 0",
		input: []*Item{
			{name: "Aged Brie", sellIn: 0, quality: 0},
		},
		expected: []*Item{
			{name: "Aged Brie", sellIn: -1, quality: 2},
		},
	},
	{
		name: "item 'Aged Brie' 'quality' does not increase over 50",
		input: []*Item{
			{name: "Aged Brie", sellIn: 50, quality: 50},
		},
		expected: []*Item{
			{name: "Aged Brie", sellIn: 49, quality: 50},
		},
	},
	// special item: Sulfuras, Hand of Ragnaros
	{
		name: "item 'Sulfuras, Hand of Ragnaros' 'quality' and 'sellIn' do not change",
		input: []*Item{
			{name: "Sulfuras, Hand of Ragnaros", sellIn: 10, quality: 10},
		},
		expected: []*Item{
			{name: "Sulfuras, Hand of Ragnaros", sellIn: 10, quality: 10},
		},
	},
	// special item: Backstage passes to a TAFKAL80ETC concert
	{
		name: "item 'Backstage passes' 'quality' increases as 'sellIn' decreases",
		input: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 20, quality: 20},
		},
		expected: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 19, quality: 21},
		},
	},
	{
		name: "item 'Backstage passes' 'quality' increases by 2 when 'sellIn' < 10 & > 5",
		input: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 10, quality: 10},
		},
		expected: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 9, quality: 12},
		},
	},
	{
		name: "item 'Backstage passes' 'quality' increases by 3 when 'sellIn' < 5 & > 0",
		input: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 5},
		},
		expected: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 8},
		},
	},
	{
		name: "item 'Backstage passes' 'quality' becomes 0 when 'sellIn' is < 0",
		input: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 0, quality: 5},
		},
		expected: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: -1, quality: 0},
		},
	},
	{
		name: "item 'Backstage passes' 'quality' never does not increase when over 50",
		input: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 5, quality: 50},
		},
		expected: []*Item{
			{name: "Backstage passes to a TAFKAL80ETC concert", sellIn: 4, quality: 50},
		},
	},
}

func TestTable(t *testing.T) {
	for _, test := range tests {
		fmt.Printf("=== Test: %v ===\n", test.name)
		UpdateQuality(test.input)
		if err := compareItems(test.input, test.expected); err != nil {
			t.Fatalf("%v", err)
		}

	}
}
