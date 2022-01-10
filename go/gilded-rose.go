package main

type Item struct {
	name            string
	sellIn, quality int
}

const MAX_QUALITY = 50

func UpdateQuality(items []*Item) {
	for _, item := range items {
		//UPDATE SELLIN
		if item.name != "Sulfuras, Hand of Ragnaros" {
			item.sellIn--
		}
		// UPDATE QUALITY
		switch item.name {
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
		case "Sulfuras, Hand of Ragnaros":
		default:
			if item.quality <= 0 {
				continue
			}
			if item.quality > 0 {
				item.quality--
			}
			if item.sellIn < 0 && item.quality > 0 {
				item.quality = item.quality - 1
			}
		}
	}

}
