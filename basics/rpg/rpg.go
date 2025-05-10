package rpg

import (
	"fmt"
	"slices"
)

type Player struct {
	Name      string
	Inventory []Item
}

type Item struct {
	Name string
	Type string
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("1x %s collected\n", item.Name)
}

func (p *Player) DropItem(itemName string) {
	indexToRemove := -1

	for itemIndex, item := range p.Inventory {
		if item.Name == itemName {
			indexToRemove = itemIndex
			break
		}
	}

	if indexToRemove >= 0 {
		removeItemFromInventory(p, indexToRemove)
		fmt.Printf("1x %s dropped\n", itemName)
	} else {
		fmt.Printf("No item found\n")
	}
}

func (p *Player) UseItem(itemName string) {
	indexToUse := searchItemIndexByName(p, itemName)

	if indexToUse >= 0 {
		item := p.Inventory[indexToUse]

		switch item.Type {
		case "potion":
			fmt.Printf("You have consumed 1x %s \n", itemName)
			removeItemFromInventory(p, indexToUse)
		case "money":
			fmt.Printf("You cannot use %s\n", item.Name)
		default:
			fmt.Printf("Unknown item type %s", item.Type)
		}
	} else {
		fmt.Printf("No item found\n")
	}
}

func removeItemFromInventory(p *Player, itemIndex int) {
	p.Inventory = slices.Delete(p.Inventory, itemIndex, itemIndex+1)
}

func searchItemIndexByName(p *Player, itemName string) int {
	itemIndex := -1

	for itemIndex, item := range p.Inventory {
		if item.Name == itemName {
			return itemIndex
		}
	}

	return itemIndex
}
