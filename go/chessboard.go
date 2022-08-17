package chessboard

type Rank []bool

type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	boardRank, exists := cb[rank]
	if !exists {
		return 0
	}

	count := 0
	for _, v := range boardRank {
		if v {
			count += 1
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	if file < 1 || file > 8 {
		return 0
	}

	count := 0
	for k, rank := range cb {
		if rank[file-1] {
			count += 1
		}
		_ = k
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for _, rank := range cb {
		count += len(rank)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, rank := range cb {
		for _, square := range rank {
			if square {
				count += 1
			}
		}
	}
	return count
}
