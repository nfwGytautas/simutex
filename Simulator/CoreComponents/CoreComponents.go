package core_components

var index uint64 = 0

func getIndex() uint64 {
	index++
	return index
}
