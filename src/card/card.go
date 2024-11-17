package card

import "fmt"

type Card struct {
	Shape  ECardShape
	Number int
	Level  int
}

func NewCard(shape ECardShape, num int) *Card {
	newCard := &Card{}
	newCard.Shape = shape
	newCard.Number = num

	newCard.Level = NumberToLevel(num)

	return newCard
}

func ECardShapeToString(shape ECardShape) string {
	shapeStrings := []string{"None", "Spade", "Heart", "Club", "Diamond"}
	return shapeStrings[shape]
}

func (c *Card) String() string {
	return fmt.Sprintf("(%s, %d)", ECardShapeToString(c.Shape), c.Number)
}

// 3,...,10,11,12, 13, 1, 2, 14,15   number
// 1,...,8, 9, 10, 11, 12,13,14,15   level
func NumberToLevel(num int) int {
	if num < 3 {
		return num + 11
	} else if num < SmallJoker {
		return num - 2
	} else {
		return num
	}
}

func (c *Card) Clone() *Card {
	return &Card{
		Shape:  c.Shape,
		Number: c.Number,
		Level:  c.Level,
	}
}
