package main

import "fmt"

type client struct {
}

func (c *client) insertSimCardInPhone(ph phone) {
	ph.insertSim()
}

type phone interface {
	insertSim()
}

type nokiaSim struct{}

func (n *nokiaSim) insertSim() {
	fmt.Println("Insert a sim card into your phone")
}

type sumsungSim struct{}

func (s *sumsungSim) insertMiniSim() {
	fmt.Println("Insert a mini sim card into your phone")
}

type sumsungSimAdapter struct {
	sumsungMiniSim *sumsungSim
}

func (s *sumsungSimAdapter) insertSim() {
	s.sumsungMiniSim.insertMiniSim()
}

func main() {
	client := &client{}

	nokiaSim := &nokiaSim{}
	client.insertSimCardInPhone(nokiaSim)

	sumsungMiniSim := &sumsungSim{}
	simAdapter := &sumsungSimAdapter{
		sumsungMiniSim: sumsungMiniSim,
	}
	client.insertSimCardInPhone(simAdapter)

}

