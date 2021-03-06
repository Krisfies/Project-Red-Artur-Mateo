package main

import (
	"fmt"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

type Personnage struct {
	// creation de la structure de notre personnage
	name      string
	class     string
	level     int
	lpmax     int
	lp        int
	inventory []string
	skill     []string
	money     int
	damage    int
	exp       int
	expmax    int
}

type Equipement struct {
	Chapeau string
	Tunique string
	Bottes  string
	Arme 	string
}

type Monstre struct {
	name   string
	lp     int
	lpmax  int
	attack int
}

func main() {
	// fonction qui execute nos sous fonctions et rentre les valeur ainsi que le menu principal
	fmt.Println("Les deux artistes sont ABBA et Spielberg")
	var a Equipement
	var p1 Personnage
	var e3 Monstre
	var e4 Monstre
	var e5 Monstre
	var e6 Monstre
	var e7 Monstre
	e3.InitGoblin("Python", 160, 8)
	e4.InitGoblin("Java", 200, 10)
	e5.InitGoblin("C++", 300, 15)
	e6.InitGoblin("Golang", 400, 20)
	p1.CharCreation(&a, &e6)
	e7.InitGoblin("L'Homme en Jaune", 1000, 50)
	Slow("Vous vous rendez sur la ", 1)
	Slow(Yellow+"Place Principale"+Reset, 1)
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(". ")
	time.Sleep(1 * time.Second)
	fmt.Printf(".")
	time.Sleep(1 * time.Second)
	p1.menu(&e3, &e4, &e5, &e6, &e7, &a)
}
