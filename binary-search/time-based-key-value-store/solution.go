/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
package timebasedkeyvaluestore

type TimeMap struct {
	store map[string][]*pair
}

type pair struct {
	val string
	ts  int
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]*pair, 0),
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	if _, found := tm.store[key]; !found {
		tm.store[key] = make([]*pair, 0)
	}

	tm.store[key] = append(tm.store[key], &pair{val: value, ts: timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	pairs, found := tm.store[key]
	if !found {
		return ""
	}

	return search(pairs, timestamp)
}

func search(pairs []*pair, target int) string {
	var val string

	left, right := 0, len(pairs)-1
	for left <= right {
		mid := left + (right-left)/2
		curr := pairs[mid]

		if target == curr.ts {
			return curr.val
		}

		if target > curr.ts {
			val = curr.val
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return val
}

func (tm *TimeMap) Len() int {
	return len(tm.store)
}
