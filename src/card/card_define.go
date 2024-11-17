package card

type ECardShape int

const (
	ECardShape_None ECardShape = iota
	ECardShape_Spade
	ECardShape_Heart
	ECardShape_Club
	ECardShape_Diamond
)

const SmallJoker int = 14
const BigJoker int = 15

type ECardType int

const (
	ECardType_None ECardType = iota
	ECardType_Single
	ECardType_Triplet
	ECardType_Sequence
)
