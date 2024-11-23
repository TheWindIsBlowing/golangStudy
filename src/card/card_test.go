package card

import "testing"

func TestFindAllCardType(t *testing.T) {
	handCards := NewCardSet(9)

	handCards.AddCard(NewCard(ECardShape_Club, 3))
	handCards.AddCard(NewCard(ECardShape_Club, 4))
	handCards.AddCard(NewCard(ECardShape_Club, 5))
	handCards.AddCard(NewCard(ECardShape_Diamond, 5))
	handCards.AddCard(NewCard(ECardShape_Club, 6))
	handCards.AddCard(NewCard(ECardShape_Club, 7))
	handCards.AddCard(NewCard(ECardShape_Club, 8))
	handCards.AddCard(NewCard(ECardShape_Club, 9))
	handCards.AddCard(NewCard(ECardShape_None, 14))
	handCards.AddCard(NewCard(ECardShape_None, 15))

	t.Log(handCards)

	// allCards := make(map[string]*CardSet, 20)
	// SplitCards(handCards, allCards)
}

/*
* 递归拆牌
*   1.递归终止条件：所有牌拆完
*   2.递归过程，确保缩小递归范围：将匹配的牌型从手牌移除
*     1）对每种牌型进行一次分析：单张、刻子（2-6张）、同花顺（13张）
 */
func SplitCards(pile *CardSet, ret map[string]*CardSet) {
	if pile.Count() == 0 {
		return
	}

	for i := ECardType_Single; i <= ECardType_Sequence; i++ {
		switch i {
		case ECardType_Single:
			// 单张：花色不同也认为不同牌型
		case ECardType_Triplet:
			// 刻子：2-6张的情况
			for num := 2; num <= 6; num++ {

			}
		case ECardType_Sequence:
			// 同花顺：四种花色，每种花色都可能有 3-13张的情况

		default:
		}
	}
}
