package main

import (
	"fmt"
	"sort"
)

type phones []int

type adressBook struct {
	name string
	phones  phones
}

type byName []adressBook

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].name < a[j].name }

func main() {
	aBook := []adressBook{
		{ "Миша", phones{78293467382} },
		{ "Никита", phones{89167253764, 89635437382} },
		{ "Борис", phones{89167253764, 89635437382} },
		{ "Алексей", phones{89167253764, 89635437382} },
	}

	fmt.Println(aBook)

	sort.Sort(byName(aBook))
	
	fmt.Println(aBook)
}

