package brotli


type symbolList struct {
	storage []uint16
	offset  int
}

func symbolListGet(sl symbolList, i int) uint16 {
	return sl.storage[i+sl.offset]
}

func symbolListPut(sl symbolList, i int, val uint16) {
	sl.storage[i+sl.offset] = val
}
