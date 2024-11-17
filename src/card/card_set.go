package card

import "strings"

type CardSet struct {
	Cards    []*Card
	CardType *CardType // 一手牌确定的牌型就一种
}

func NewCardSet(cap int) *CardSet {
	cardType := &CardType{
		Type: ECardType_None,
	}
	cardSet := &CardSet{
		Cards:    make([]*Card, 0, cap),
		CardType: cardType,
	}

	return cardSet
}

func (cs CardSet) String() string {
	ret := strings.Builder{}
	ret.WriteString("[")
	for i, v := range cs.Cards {
		ret.WriteString(v.String())
		if i != cs.Count()-1 {
			ret.WriteString(",")
		} else {
			ret.WriteString("]")
		}
	}

	return ret.String()
}

func (cs *CardSet) Clone() *CardSet {
	newCardSet := &CardSet{
		Cards:    make([]*Card, 0, cs.Count()),
		CardType: cs.CardType.Clone(),
	}
	for _, v := range cs.Cards {
		newCardSet.Cards = append(newCardSet.Cards, v.Clone())
	}

	return newCardSet
}

func (cs *CardSet) Count() int {
	return len(cs.Cards)
}

func (cs *CardSet) AddCard(card *Card) {
	cs.Cards = append(cs.Cards, card)
}

func (cs *CardSet) RemoveCard(card *Card) {
	for i, v := range cs.Cards {
		if v.Shape == card.Shape && v.Number == card.Number {
			cs.Cards = append(cs.Cards[:i], cs.Cards[i+1:]...)
		}
	}
}

func (cs *CardSet) Hash() string {
	str := strings.Builder{}
	str.WriteString(cs.CardType.String())
	return str.String()
}
