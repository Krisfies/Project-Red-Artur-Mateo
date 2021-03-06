package main

func (p *Personnage) AddInventory(item string, price int) {
	if p.money >= price {
		p.money -= price
		p.inventory = append(p.inventory, item)
	} else {
		Slow("Tu n'as pas assez ", 2)
		Slow(Yellow+"d'argent "+Reset, 2)
		Slow("pour acheter cet objet", 1)
	}
}

func (p *Personnage) RemoveInventory(item string) {
	var index int = -1
	for i := range p.inventory {
		if p.inventory[i] == item {
			index = i
		}
	}
	if index != -1 {
		p.inventory = append(p.inventory[:index], p.inventory[index+1:]...)
	}
}
