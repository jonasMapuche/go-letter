package tree

import (
	"strconv"
)

type Notice struct {
	high  *Notice
	level int
	node  int
	value string
	low   *Notice
}

func Plant() {
	var notice *Notice
	for i := 1; i <= 6; i++ {
		notice = search(notice, strconv.Itoa(i))
		insert = false
	}
}

var max_node int = 1
var insert bool = false

func search(notice *Notice, value string) *Notice {
	var high, low bool = false, false
	var level = 0
	if notice == nil {
		if max_node > 1 {
			var exp int = 0
			for {
				var no_start = 1 << exp
				var no_end = no_start + no_start - 1
				if (max_node >= no_start) && (max_node <= no_end) {
					level = exp
					break
				}
				exp++
			}
		}
		if !insert {
			print(" -- ", level, " - ", max_node, "\n")
			insert = true
			return &Notice{nil, level, max_node, value, nil}
		}
	}
	if notice.high != nil {
		notice.high = search(notice.high, value)
	} else {
		high = true
	}
	if notice.low != nil {
		notice.low = search(notice.low, value)
	} else {
		low = true
	}
	if (((high) && (low)) || ((!high) && (low))) && (!insert) {
		notice = writer(notice, value)
	}
	return notice
}

func writer(notice *Notice, value string) *Notice {
	var level int = notice.level
	var node int = notice.node
	if node > max_node {
		max_node = node
	}
	var next_node = max_node + 1
	var exp = 1 << level
	if exp == 1 {
		if notice.high == nil {
			max_node = next_node
			print(" / ")
			notice.high = search(notice.high, value)
		} else {
			if notice.low == nil {
				max_node = next_node
				print(" | ")
				notice.low = search(notice.low, value)
			}
		}
	} else {
		if node < exp+exp-1 {
			if max_node >= exp+exp-1 {
				if notice.high == nil {
					max_node = next_node
					print(" / ")
					notice.high = search(notice.high, value)
				} else {
					if notice.low == nil {
						max_node = next_node
						print(" | ")
						notice.low = search(notice.low, value)
					}
				}
			} else {
				return notice
			}
		} else {
			if notice.high == nil {
				max_node = next_node
				print(" / ")
				notice.high = search(notice.high, value)
			} else {
				if notice.low == nil {
					max_node = next_node
					print(" | ")
					notice.low = search(notice.low, value)
				}
			}
		}
	}
	return notice
}
