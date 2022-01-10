package main

type Item struct {
	name            string
	sellIn, quality int
}

const MAX_QUALITY = 50

func UpdateQuality(items []*Item) {
	for _, item := range items {

		//UPDATE SELLIN: same for all items except Sulfuras
		if item.name != "Sulfuras, Hand of Ragnaros" {
			item.sellIn--
		}

		// UPDATE QUALITY: differs significantly, this switch statement treats each item type
		// separately so allowing a single, independent policy with simple logic for each
		// at the cost of some minor duplication
		switch item.name {
		case "Sulfuras, Hand of Ragnaros":
			//make no changes for this case
		case "Aged Brie":
			if item.quality >= MAX_QUALITY {
				continue
			}
			if item.sellIn < 0 {
				item.quality += 2
			} else {
				item.quality++
			}
		case "Backstage passes to a TAFKAL80ETC concert":
			if item.quality >= MAX_QUALITY {
				continue
			}
			if item.sellIn > 0 {
				if item.sellIn < 5 {
					item.quality += 3
				} else if item.sellIn < 10 {
					item.quality += 2
				} else {
					item.quality++
				}
			} else if item.sellIn < 0 {
				item.quality = 0
			}
		case "Conjured":
			// Conjured items decrease in quality at twice the speed of default items
			// meaning a loss of 2 quality when sellIn > 0, and of 4 when sellIn > 0
			// but as with default, quality should never be negative
			if item.quality <= 0 {
				continue
			}
			if item.sellIn < 0 {
				if item.quality >= 4 {
					item.quality -= 4
				} else {
					item.quality = 0
				}
			} else {
				if item.quality >= 2 {
					item.quality -= 2
				} else {
					item.quality = 0
				}
			}
		default:
			if item.quality <= 0 {
				continue
			}
			if item.quality > 0 {
				item.quality--
			}
			if item.sellIn < 0 && item.quality > 0 {
				item.quality--
			}
		}
	}

}
