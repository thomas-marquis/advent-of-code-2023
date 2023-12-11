package days

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/thomas-marquis/advent-of-code-2023/utils"
)


var (
	reg = regexp.MustCompile(`^Card\s*\d*:\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s\|\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)$`)
)

const (
	maxCardNb = 193
)


type Card struct {
	Id int
	NbMatches int
	CopiedFromCard *Card
	IsCopy bool
}

func Day4() {
	scanner, _ := utils.ReadFileLines("resources/day4_input")

	var cards []Card
	
	i := 1
	for scanner.Scan() {
		c := Card{
			Id: i,
			IsCopy: false,
		}
		
		lineContent := scanner.Text()
		groups := reg.FindStringSubmatch(lineContent)
		
		winningNbs := utils.ToIntSlice(groups[1:11])
		userNbs := utils.ToIntSlice(groups[11:])
		
		for _, winN := range winningNbs {
			for _, userN := range userNbs {
				if winN == userN {
					c.NbMatches++
				}
			}
		}
		
		// fmt.Printf("Card %v: Winning numbers: %#v\n", c.Id, c.NbMatches)
		
		i++
		cards = append(cards, c)
	}
	
	var wg sync.WaitGroup
	
	res := len(cards)
	cardsChan := make(chan Card)
	for _, card := range cards {
		wg.Add(1)
		go getCardsCopies(card, cards, cardsChan, &wg)
	}

	for card := range cardsChan {
		res++
		if card.IsCopy {
			fmt.Printf("Card %d copied from card %d (%d): Winning numbers: %#v => res=%d\n", card.Id, card.CopiedFromCard.Id, card.CopiedFromCard.NbMatches, card.NbMatches, res)
		} else {
			fmt.Printf("Original card %d: Winning numbers: %#v => res=%d\n", card.Id, card.NbMatches, res)
		}
		if card.NbMatches > 0 {
			wg.Add(1)
			go getCardsCopies(card, cards, cardsChan, &wg)
		}
	}

	close(cardsChan)

	wg.Wait()

	fmt.Println(len(cards))
	fmt.Println(res)
}

func getCardsCopies(card Card, cards []Card, ch chan Card, wg *sync.WaitGroup) {
	defer wg.Done()

	nbMatches := card.NbMatches
	if nbMatches == 0 {
		return
	}

	currCardId := card.Id
	for i := currCardId + 1; i <= min(maxCardNb, currCardId + nbMatches); i++ {
		copiedCard := cards[i-1]
		copiedCard.CopiedFromCard = &card
		copiedCard.IsCopy = true
		ch <- copiedCard
	}
}