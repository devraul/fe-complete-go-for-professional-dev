package main

import "myModule/rpg"

func main() {
	player := rpg.Player{}

	healthPotion := rpg.Item{
		Name: "Small Health Potion",
		Type: "potion",
	}

	silverCoin := rpg.Item{
		Name: "Silver coin",
		Type: "money",
	}

	player.PickUpItem(healthPotion)
	player.PickUpItem(silverCoin)
	// fmt.Println("Inventory", player.Inventory)

	player.UseItem(healthPotion.Name)
	player.UseItem(silverCoin.Name)
	// fmt.Println("Inventory", player.Inventory)

	player.DropItem(silverCoin.Name)
	// fmt.Println("Inventory", player.Inventory)
}
