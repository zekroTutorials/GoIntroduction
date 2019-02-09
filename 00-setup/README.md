# SETUP

**DOCS:**
- https://golang.org/doc/install

---

## Installation

### Windows:

Für Windows muss man sich eifnach nur den [Installer](https://golang.org/dl/#featured) für Windows herunterladen und ausführen.  
Danach sollte man in der Console auf den Command `go` zugreifen können.

Genaueres dazu findet man auch [hier](https://golang.org/doc/install#windows).


### Linux, MacOSX & co:

Für Linux und co. läd man sich zuerst das [Binary Package](https://golang.org/dl/#featured) herunter. Dies geht über die Console sehr einfach mit:
```bash
$ wget https://dl.google.com/go/go[version].linux-[arch].tar.gz
```
*Natürlich braucht ihr hier die aktuelle URL von der Website und die Architektur eures Systems.*

Dannn entpackt ihr das package einfach nach `/usr/local`:
```bash
$ tar -C /usr/local -xzf go[version].linux-[arch].tar.gz
```

Nun müsst ihr nur noch die Executable zu der PATH-Variable hinzufügen:
```bash
$ export PATH=$PATH:/usr/local/go/bin
```

Nun solltet ihr auch auf den Command `go` in der Console zugreifen können.

---

## Einrichtung des Workspaces

Bei Go ist eine wichtige Sache zu beachten beim aufsetzen eines Workspaces:
Go artbeitet mit einer Umgebungsvariable namens `GPATH`, von der aus nach Packages und Source Files gesucht wird und wo diese auch installiert werden.

Nun kann man 2 Wege einschlagen:

1. Man setzt einen generellen `GOPATH` wo man alle Packages hin installiert und drauf zugreift (vergleichbar mit PIP bei Python)
2. Man setzt pro Projekt immer den `GOPATH` in die aktuelle Directory

*Ich bevorzuge Nr. 2 und werde es auch anhand dessen zeigen, jedoch wollte ich beide Möglichkeiten mal erwähnen.*


Nun erstellt man den Workspace nach folgenndem Schema:

1. Wir haben eine Root-Directory wo alle unsere Projekte liegen namens `/dev`

2. Wir erstellen einen Ornder mit dem Namen des Projekts und betreten diesen  
```
/dev$ mkdir myProject && cd myProjekt
```

3. Nun erstellen wir folgenden Pfad in dem Ornder des Projekts:
```
/dev/myProject$ mkdir -p src/github.com/myUsername/myProject
```
Das ganze könnte auch folgendermaßen aussehen, wenn ihr nicht vor habt, die Repository auf GitHub zu hosten:
```
/dev/myProject$ mkdir -p src/myWebsite.de/myPorject
```

4. Nun setzen wir hier unseren `GOPATH`. Bei Linux wird das folgendermaßen gemacht:
```
/dev/myProject$ export GOPATH=/dev/myProject
```
Für Windows funktioniert das mit:
```
C:\dev\myProject> setx GOPATH "C:\dev\myProject"
```
*Dabei ist darauf zu achten, dass bei Windows die CMD erst neu gestartet werden muss, bevor die Variable nutzbar ist!*

5. Nun ist unsere eigentliche Work Directory der Ordner `myProject` in dem erstellten Pfad. Dies ist wichtig zu beachten, da Packages ihre Dateien nach dem selben Prinziep ablegen.  
Also ein Package, z.B. `github.com/zekroTJA/argparser`, würde beim installieren seine dateien in folgenden Pfad legen:
```
/dev/myProject/src/github.com/zekroTJA/argparser
```

6. `(optional)` Beim erstellen einer GitHub Repository muss darauf geachtet werden, dass diese in dem Ordner `/dev/myProject/src/github.com/myUsername/myProject` erstellt wird.

---

## Das erste Programm

Nun erstellen wir das erste Program, indem wir eine `.go` Datei anlegen, in der usner Code definiert wird:

> helloworld.go
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World!")
}
```

Diese datei können wir nun temporär kompilieren und ausführen lassen mit
```
$ go run helloworld.go
Hello World!
```

Wollen wir nun die Binary des Programms, so nutzen wir:
```
$ go build -o helloworld
$ ./helloworld
Hello World!
```