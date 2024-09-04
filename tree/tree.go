package tree

import (
	"strconv"
)

type Notice struct {
	high   *Notice
	level  int
	node   int
	value  string
	low    *Notice
	behind *Notice
}

func Plant() {
	var notice *Notice
	for i := 1; i <= 1; i++ {
		notice = search(notice, strconv.Itoa(i), nil)
		insert = false
	}
}

var max_node int = 1
var insert bool = false

func search(notice *Notice, value string, behind *Notice) *Notice {
	var high, low bool = false, false
	var level = 0
	var node_behind int = 0
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
		if behind == nil {
			node_behind = 0
		} else {
			node_behind = behind.node
		}
		print(" -level- ", level, " -behind- ", node_behind, " -node- ", max_node, "\n")
		insert = true
		return &Notice{nil, level, max_node, value, nil, behind}
	}
	if notice.high != nil {
		notice.high = search(notice.high, value, notice)
	} else {
		high = true
	}
	if notice.low != nil {
		notice.low = search(notice.low, value, notice)
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
			notice.high = search(notice.high, value, notice)
		} else {
			if notice.low == nil {
				max_node = next_node
				print(" | ")
				notice.low = search(notice.low, value, notice)
			}
		}
	} else {
		if node < exp+exp-1 {
			if max_node >= exp+exp-1 {
				if notice.high == nil {
					max_node = next_node
					print(" / ")
					notice.high = search(notice.high, value, notice)
				} else {
					if notice.low == nil {
						max_node = next_node
						print(" | ")
						notice.low = search(notice.low, value, notice)
					}
				}
			} else {
				return notice
			}
		} else {
			if notice.high == nil {
				max_node = next_node
				print(" / ")
				notice.high = search(notice.high, value, notice)
			} else {
				if notice.low == nil {
					max_node = next_node
					print(" | ")
					notice.low = search(notice.low, value, notice)
				}
			}
		}
	}
	return notice
}
