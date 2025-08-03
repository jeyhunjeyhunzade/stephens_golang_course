package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { // receiver function
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // returning multiple values on one func
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)

	if err != nil {
		// Option1: log the error, and return a call to newDeck()

		// Option2: log the error, and quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	stringSlice := strings.Split((string(byteSlice)), ",")
	return deck(stringSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // custom rand

	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // random number between 0 and last items's index in deck
		d[i], d[newPosition] = d[newPosition], d[i] //swapping
	}
}
