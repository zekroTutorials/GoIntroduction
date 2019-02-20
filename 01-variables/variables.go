// On Playground: https://play.golang.com/p/4y1Iq3iHj4i

package main

import "fmt"

func main() {

	// Erstellt einen Integer (int64)
	// initialisiert mit dem Wert 1
	var myInt int

	// String mit dem initialisiertem Wert ""
	var myString string

	// Initialisiert eine Variable implizit
	// als den Typ des Wertes, hier float64
	pi := 3.1415

	// Es ist auch möglich Variablen in einer
	// art Chain zu initialisieren
	x, y, z := 3, 5, 1

	// Definiert ein Integer Array mit der
	// Größe 0 und dem initialisiertem Wert
	// [ ] (len = 0)
	var myArray []int

	// Es ist auch mölglich, ein leeres,
	// initialisiertes Array fester Größe
	// zu erstellen.
	var mySizedArray [5]int

	// Das Selbe ist auch mit folgender
	// Funktion möglich
	mySizedArray2 := make([]int, 5)

	// Oder man initialisiert implizit ein
	// Aray mit vordefinierten Werten.
	myStringArray := []string{"hey", "ho"}

	// Erstellt eine Map mit dem Key-Type string
	// und dem Value-Type int
	var myMap map[string]int

	// Ist im Prinzip das selbe wie
	// folgende Funktion
	myMap2 := make(map[string]int)

	// Alle Werte der Variablen ausgegeben, damit
	// der Compiler nicht meckert ;)
	fmt.Println(myInt, myString, pi, x, y, z, myArray, mySizedArray,
		mySizedArray2, myStringArray, myMap, myMap2)
}
