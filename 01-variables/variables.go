package main

import "fmt"

func main() {

	// Definition einer uninitialisierten Integer Variable
	var myInt int
	// Initialisierung dieser Variable mit dem Wert 5
	myInt = 5

	// Hier noch ein paar andere Datentypen aufgelistet:
	var myString string
	var myFloat32 float32
	var myFloat64 float64

	// So definiert man uninitialisierte Arrays:
	var myStringArray []string
	var myIntArray []int
	// ...
	// Diese Arrays sind 0-Arrays, welche keine Indices haben.
	// myStringArray[0] = 1 -> Das würde einen 'index out of range' error geben
	// Deswegen muss man an der stelle das Array erweitern mit:
	myIntArray = append(myIntArray, 3)

	// Man kann auch, wie bei Java oder C++ gewohnt, Arrays
	// mit einer festgelegten Größe erstellen:
	var myFloatArray [5]float32
	// Dieses Array ist nun initialisiert mit den Wreten [0 0 0 0 0]
	// Wie gewoht kann man hier den Wert eines Index bestimmen:
	myFloatArray[2] = 3.1415

	// Es gibt auch noch folgende Möglichkeit, variablen direkt zu
	// definieren und zu initialisieren:
	myInitializedString := "Hey!"
	myInitializedInt := 3
	// Go erkennt, mit welchen Typ die Variable erstellt werden soll
	// und deren wert wird direkt festgesetzt.

	// Dies funktioniert auch beim erstellen eines Arrays mit
	// festgelegter Größe:
	myInitializedStringArray := make([]string, 5)
	// Dieses Array beinhaltet dann auch 5 leere ("") Strings.

	// Es gibt auch einen Typ 'interface{}', welcher
	// einen unbestimmten Typ ausdrückt:
	var myVar interface{}
	myVar = 5
	myVar = "5"
	myVar = 5.0

	fmt.Println(myInt, myString, myFloat32,
		myFloat64, myStringArray, myIntArray,
		myFloatArray, myInitializedString,
		myInitializedInt, myInitializedStringArray, myVar)
}
