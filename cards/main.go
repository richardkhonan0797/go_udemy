package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

	// print(cards.toString())

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// three component loop
	// for i := 0; i < 2; i++ {
	// 	fmt.Println(cards[i])
	// }

	// for-each loop
	// for index, card := range cards {
	// 	fmt.Println(index, card)
	// }
}
