package card

import "fmt"

type CardType struct {
	Type     ECardType
	MinLevel int
	Count    int
}

func NewCardType(cardType ECardType, minLevel int, count int) *CardType {
	return &CardType{
		Type:     cardType,
		MinLevel: minLevel,
		Count:    count,
	}
}

func (ct *CardType) String() string {
	return fmt.Sprintf("%d-%d-%d", ct.Type, ct.MinLevel, ct.Count)
}

func (ct *CardType) Clone() *CardType {
	return &CardType{
		Type:     ct.Type,
		MinLevel: ct.MinLevel,
		Count:    ct.Count,
	}
}
