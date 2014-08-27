package tools

import ()

//	see if array contains a value
func AryContains(val string, ary []string) int {
	for i, v := range ary {
		if v == val {
			return i
		}
	}
	return -1
}
