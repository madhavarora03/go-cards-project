package main

func main() {
	cards := deck{"Ace of Spades", getCard()}

	cards.print()
}

func getCard() string {
	return "Five of Diamonds"
}
