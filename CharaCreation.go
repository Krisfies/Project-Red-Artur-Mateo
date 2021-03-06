package main

import (
	"fmt"
	"os"
	"time"
)

func (p *Personnage) CharCreation(a *Equipement, e6 *Monstre) {
	var name string
	var class string
	var level int
	var lpmax int
	var lp int
	var inventory []string
	var skill []string
	var money int
	var damage int
	var chapeau string
	var tunique string
	var bottes string
	var expmax int

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	Slow("Bienvenue dans ", 1)
	Slow(Yellow+"Goldy"+Reset, 1)
	time.Sleep(100 * time.Millisecond)
	Slow("\nVous êtes dans le menu de ", 1)
	Slow(Yellow+"création de personnage\n"+Reset, 1)
	Slow("Pour commencer, choisissez un ", 1)
	Slow(Yellow+"nom "+Reset, 1)
	Slow("pour votre ", 1)
	Slow(Yellow+"avatar"+Reset, 1)
	Slow(": \n", 1)
	fmt.Scanln(&name)
	if name == "Utilisateur" { //Easter egg n°1, le mode développeur du jeu
		e6.lp = 0
		class = "Utilisateur"
		level = 1
		lpmax = 9999
		lp = lpmax
		inventory = []string{"Véritable Couteau", "Objet Suspicieux"}
		money = 10000
		chapeau = "Couronne en or"
		tunique = "Cape en fourrure, ornée de cristaux"
		bottes = "Bottes en cuir"
		Slow(Yellow+"Bienvenue Utilisateur\n"+Reset, 1)
		time.Sleep(300 * time.Millisecond)
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	} else {
		Slow("Votre personnage se nomme donc ", 1)
		fmt.Print(Yellow + "")
		Slow(name, 1)
		fmt.Print("" + Reset)
		time.Sleep(1 * time.Second)
		Slow("\nChoisissez maintenant la ", 1)
		Slow(Yellow+"race "+Reset, 1)
		Slow("de ", 1)
		fmt.Print(Yellow + "")
		Slow(name, 1)
		fmt.Print("" + Reset)
		Slow(" parmi:\n", 1)
		Slow(Yellow+"\n-Humain\n", 1)
		Slow("\n-Elfe\n", 1)
		Slow("\n-Nain\n \n"+Reset, 1)
		fmt.Scanln(&class)
		if class != "Humain" && class != "Elfe" && class != "Nain" {
			Slow("Erreur, veuillez entrer une valeur correcte:\n", 1)
			Slow("Humain, Elfe ou Nain\n", 1)
			fmt.Scanln(&class)
		}
		switch class {
		case "Humain":
			class = "Humain"
			lpmax = 100
			damage = 8
			chapeau = "Chapeau de paille"
			tunique = "Vieux manteau"
			bottes = "Vieille claquette"
		case "Elfe":
			class = "Elfe"
			lpmax = 80
			damage = 3
			chapeau = "Chapeau d'érudit"
			tunique = "Robe de sage"
			bottes = "Chaussure pointue"
		case "Nain":
			class = "Nain"
			lpmax = 120
			damage = 5
			chapeau = "Casque de mineur"
			tunique = "Salopette rapiécée"
			bottes = "Sabot renforcé"
		}
		if class != "Humain" && class != "Elfe" && class != "Nain" { //Easter egg n°3
			class = "Troll"
			lpmax = 50
			damage = 2
			chapeau = "Bonnet de petite taille"
			tunique = "Veste abîmée"
			bottes = "Sabot en boît"

		}
		lp = lpmax / 2
		money = 100
		inventory = []string{"Potion de vie", "Potion de vie", "Potion de vie"}
		skill = []string{"Coup de poing"}
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Printf(Yellow + "")
		Slow(name, 1)
		Slow(Reset+" est donc un "+Yellow, 1)
		Slow(class, 1)
		Slow(Reset+", il commence avec "+Yellow, 1)
		fmt.Print(lp)
		Slow(" points de vie "+Reset, 1)
		Slow("et "+Yellow, 1)
		fmt.Print(lpmax)
		Slow(" points de vie maximum.\n"+Reset, 1)
		if class == "Troll" {
			Slow("Ça t'apprendras à pas savoir écrire\n", 1)
		}
		fmt.Printf(Yellow + "")
		Slow(name, 1)
		Slow(Reset+" est ", 1)
		Slow(Yellow+"niveau 1, "+Reset, 1)
		Slow("possède le sort ", 1)
		Slow(Yellow+"Coup de poing "+Reset, 1)
		Slow("et a "+Yellow, 1)
		fmt.Print(money)
		Slow(" pièces"+Reset, 1)
		level = 1
		time.Sleep(3 * time.Second)
	}
	Slow("\nLancement de la simulation\n", 1)
	a.Chapeau = chapeau
	a.Tunique = tunique
	a.Bottes = bottes
	expmax = 15
	p.Init(name, class, level, lpmax, lp, inventory, skill, money, damage, expmax)
}

func (p *Personnage) Init(name string, class string, level int, lpmax int, lp int, inventory []string, skill []string, money int, damage int, expmax int) {
	// initialisation de notre personnage
	p.name = name
	p.class = class
	p.level = level
	p.lpmax = lpmax
	p.lp = lp
	p.inventory = inventory
	p.skill = skill
	p.money = money
	p.damage = damage
	p.expmax = expmax
}
