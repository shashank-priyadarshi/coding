Conceptual Example
Say, you have two types of computers: Mac and Windows. Also, two types of printers: Epson and HP. Both computers and printers need to work with each other in any combination. The client doesn’t want to worry about the details of connecting printers to computers.

If we introduce new printers, we don’t want our code to grow exponentially. Instead of creating four structs for the 2*2 combination, we create two hierarchies:

Abstraction hierarchy: this will be our computers
Implementation hierarchy: this will be our printers
These two hierarchies communicate with each other via a Bridge, where the Abstraction (computer) contains a reference to the Implementation (printer). Both the abstraction and implementation can be developed independently without affecting each other.

### computer.go: Abstraction
```go
package main

type Computer interface {
    Print()
    SetPrinter(Printer)
}
```
### mac.go: Refined abstraction
 ```go
package main

import "fmt"

type Mac struct {
    printer Printer
}

func (m *Mac) Print() {
    fmt.Println("Print request for mac")
    m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
    m.printer = p
}
```
### windows.go: Refined abstraction
```go
package main

import "fmt"

type Windows struct {
    printer Printer
}

func (w *Windows) Print() {
    fmt.Println("Print request for windows")
    w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
    w.printer = p
}
```
### printer.go: Implementation
```go
package main

type Printer interface {
    PrintFile()
}
```
### epson.go: Concrete implementation
 ```go
package main

import "fmt"

type Epson struct {
}

func (p *Epson) PrintFile() {
    fmt.Println("Printing by a EPSON Printer")
}
```
### hp.go: Concrete implementation
 ```go
package main

import "fmt"

type Hp struct {
}

func (p *Hp) PrintFile() {
    fmt.Println("Printing by a HP Printer")
}
```
### main.go: Client code
 ```go
package main

import "fmt"

func main() {

    hpPrinter := &Hp{}
    epsonPrinter := &Epson{}

    macComputer := &Mac{}

    macComputer.SetPrinter(hpPrinter)
    macComputer.Print()
    fmt.Println()

    macComputer.SetPrinter(epsonPrinter)
    macComputer.Print()
    fmt.Println()

    winComputer := &Windows{}

    winComputer.SetPrinter(hpPrinter)
    winComputer.Print()
    fmt.Println()

    winComputer.SetPrinter(epsonPrinter)
    winComputer.Print()
    fmt.Println()
}
```
### output.txt: Execution result
```
Print request for mac
Printing by a HP Printer

Print request for mac
Printing by a EPSON Printer

Print request for windows
Printing by a HP Printer

Print request for windows
```