package leaderboard

import ()

//Item is an interface for a leaderboard item. It only exposes a comparison method
type Item interface {
	LessThan(Item) bool
}

//LeaderBoard implements a leaderboard
type LeaderBoard []*Item

//Collect adds an item to a leaderboard, at the proper position
func (lb *LeaderBoard) Collect(newItem Item) {
	lblen := len(*lb)
	lbcap := cap(*lb)
	if lblen == lbcap && newItem.LessThan(*(*lb)[lblen-1]) {
		//if leaderboard is full and the new item is smaller than the last item in leaderboard, we can just discard it
		return
	}
	for i, item := range *lb {
		if (*item).LessThan(newItem) {
			if lblen < lbcap {
				*lb = append(*lb, nil)
			}
			copy((*lb)[i+1:], (*lb)[i:])
			(*lb)[i] = &newItem
			return
		}
	}

	if lblen < lbcap {
		*lb = append(*lb, &newItem)
		return
	}

}

//New creates a new leaderboard of a size
func New(size int) LeaderBoard {
	lb := make([]*Item, 0, size)
	return lb
}
