// Die erste Zeile in einer Go-Datei muss immer der Package-Name sein.
// Da diese Datei in keinem Unterverzeichnis (Package) ist, ist der Package
// name üblicherweise 'main' oder 'core'.
package main

// Als nächstes importieren wir unsere benötigten Dependencies.
// Dabei sind alle Packages ohne Pfadangabe ('/' davor) globale
// Go Standart Packages
import "fmt"

// Main Funktion, wo das Prgramm immer einsteigt, wenn es startet.
func main() {
	fmt.Println("Hello World!")
}