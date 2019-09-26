package leaderboard
import(
"fmt"
)
//Item is an interface for a leaderboard item. It only exposesa comparison method
type Item interface{
  LessThan(*Item) bool
}

//LeaderBoard implements a leaderboard
type LeaderBoard []*Item


//Collect adds an item to a leaderboard, at the proper position
func (lb *LeaderBoard)Collect(newItem *Item){
  fmt.Printf("Collecting %v\n", *newItem)
  if newItem == nil {
    return
  }
  lblen := len(*lb)
  lbcap := cap(*lb)
  if lblen == lbcap && (*newItem).LessThan((*lb)[lblen-1]){
    //if leaderboard is full and the new item is smaller than the last item in leaderboard, we can just discard it
    return
  }
  if lblen == 0 || ( lblen < lbcap && (*newItem).LessThan((*lb)[lblen-1])) {
    *lb = append(*lb, newItem)
    return
  }
  for i, item := range(*lb){
    if (*item).LessThan(newItem){
      if lblen < lbcap{
        *lb = append(*lb, nil)
      }
      copy((*lb)[i:i+lblen], (*lb)[i+1:])
      (*lb)[i] = newItem
      return
    }
  }
  


}
//NewLeaderBoard creates a new leaderboard of a size
func NewLeaderBoard(size int) LeaderBoard{
  lb := make([]*Item, 0, size)
  return lb
}


