package main

import "fmt"

// Board represents a surface we can work on
type Board struct {
	NailsNeeded int
	NailsDriven int
}

// NailDriver represents behavior to drive nails into a board.
type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

// NailPuller represents behavior to remove nails from board.
type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

// NailDrivePuller represents behavior to drive/remove nails into/from a board
// Any concrete type value that implements both  Driver and Puller behaviors will also implement NailDrivePuller interface
type NailDrivePuller interface {
	NailDriver
	NailPuller
}

// Hammer is a tool to pound in nails
type Hammer struct{}

// DriveNail pounds a nail into the specified board and now Hammer implements NailDriver
func (Hammer) DriveNail(nailSupply *int, b *Board) {
	*nailSupply--
	b.NailsDriven++
	fmt.Println("Hammer: pounded nail into the board.")
}

// Ploskogybzu is a tool that removes nails
type Ploskogybzu struct{}

// PullNail removes a nail from the specified board and now Ploskogybzu implements NailPuller
func (Ploskogybzu) PullNail(nailSupply *int, b *Board) {
	b.NailsDriven--
	*nailSupply++
	fmt.Println("Ploskogybzu: removed nail from the board.")
}

// Contractor carries out the task of securing boards
type Contractor struct{}

// Fasten will drive nails into a board
func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) { // First parameter is a value that implements the NailDriver interface
	for b.NailsDriven < b.NailsNeeded { // This value represents the tool the contractor will use to execute this behavior
		d.DriveNail(nailSupply, b)
	}
}

// Unfasten will remove nails from a board.
func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

// ProccessBoards works against boards.
func (c Contractor) ProccessBoards(dp NailDrivePuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(dp, nailSupply, b)

		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(dp, nailSupply, b)
		}
	}
}

// Toolbox can contain any number of tools
// Since toolbox embeds both NailDriver and NailPuller interfaces, this means Toolbox also implements  NailDrivePuller interface
type Toolbox struct {
	NailDriver
	NailPuller
	nails int
}

func displayState(tb *Toolbox, boards []Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}
	fmt.Println()
}

func main() {
	boards := []Board{
		{NailsDriven: 3},
		{NailsDriven: 1},
		{NailsDriven: 6},
		{NailsNeeded: 6},
		{NailsNeeded: 9},
		{NailsNeeded: 4},
	}

	tb := Toolbox{
		NailDriver: Hammer{},      // create a value of the struct type Hammer and assign it to the inner interface type NailDriver
		NailPuller: Ploskogybzu{}, // create a value of struct type Ploskogybzu and assign it to the inner interface type NailPuller
		nails:      10,
	}

	var c Contractor
	c.ProccessBoards(&tb, &tb.nails, boards)

	displayState(&tb, boards)
}
