# leaderboard
Trivial leaderboard structure for golang.
Just a simple, no-frills leaderboard structure, as most that I found for go were relying on redis to provide the same feature.

To use, just define a type that implements the Item interface. Then obtain a new board with leaderboard.New(size) and call its Collect method to add items.

