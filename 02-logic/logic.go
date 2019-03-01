package main

import (
	"fmt"
	"strings"
)

func main() {

	a := 3

	//////////////////////////////////////////////////
	// IF
	// if <condition> {  }
	// Wichtig:
	//   Klammern um die Condition sind optional,
	//   jedoch unkonventionell und werden generell
	//   von gofmt entfernt.

	if 3 > 2 {
		fmt.Printf("%d ist größer als 3\n", a)
	}

	// Natürlich kann man auch Conditions aneinander ketten mit
	// && (and) und || (or).
	if a != 0 && a < 10 || a == 5 {
		fmt.Printf("%d ist nicht 0 und kleiner als 10 oder %d ist 5\n", a, a)
	}

	// Natürlich funktionieren auch Funktionen mit booleans als
	// Rückgabewert in der Condition.
	str := "Hey, was geht ab!"
	if strings.Contains(str, "geht") {
		fmt.Printf("%s beinhaltet \"geht\"\n", str)
	}

	//////////////////////////////////////////////////
	// LOOPS
	// for <condition> { } oder
	// for { }
	// Auch hier gilt wieder bei der Condition das
	// selbe, wie oben bei IF erwähnt.

	// Klassischer gezählter Loop. Die variable i ist
	// natürlich nur im scope des Loops erreichbar.
	for i := 0; i < 3; i++ {
		fmt.Println("Loop 1: ", i)
	}

	// Natürlich kann man auch einen Loop auf eine
	// boolsche Condition beziehen.
	cond := true
	for cond {
		fmt.Println("Loop 2: HI!")
		cond = false
	}

	// Ein 'for' ohne condition ist ein Endlos-Loop,
	// welcher bis zum Abbrechen des Programms oder
	// bis zum call eines 'breaks' läuft.
	i := 0
	for {
		fmt.Println("Loop 3: ", i)
		if i >= 5 {
			// 'break' bricht die Schleife
			// des Scopes ab.
			break
		}
		i++
	}

	//////////////////////////////////////////////////
	// SWITCH
	// switch <var> {
	//     case <condition>, [<condition2>]:
	//         // ...
	// }

	v := 1
	switch v {
	case 0:
		fmt.Println("v ist 1")
	case 1, 2:
		fmt.Println("v ist 1")
	default:
		fmt.Println("default")
	}
}
