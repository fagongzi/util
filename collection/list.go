package collection

import (
	"container/list"
)

// Get return element at index of list
func Get(l *list.List, index int) *list.Element {
	if nil == l || l.Len() == 0 {
		return nil
	}

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i == index {
			return e
		}

		i++
	}

	return nil
}

// IndexOf return index of element in list
func IndexOf(l *list.List, value interface{}) int {
	i := 0
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		if iter.Value == value {
			return i
		}

		i++
	}

	return -1
}

// Remove remove from list
func Remove(l *list.List, value interface{}) {
	var el *list.Element
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			el = e
			break
		}
	}

	if nil != el {
		l.Remove(el)
	}
}

// Truncate trancate list, keeping the first size elements
func Truncate(l *list.List, size int) *list.List {
	if size >= l.Len() {
		return l
	}

	newL := list.New()
	keeping := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if keeping < size {
			keeping++
			newL.PushBack(e.Value)
			continue
		}

		break
	}

	return newL
}

// Iterate iterate list
func Iterate(l *list.List, skip int, f func(interface{}) bool) {
	skipped := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if skipped < skip {
			skipped++
			continue
		}

		if !f(e.Value) {
			return
		}
	}
}

// ToStringArray return string array
func ToStringArray(l *list.List) []string {
	if nil == l {
		return nil
	}

	values := make([]string, l.Len())

	i := 0
	for iter := l.Front(); iter != nil; iter = iter.Next() {
		s, _ := iter.Value.(string)
		values[i] = s

		i++
	}

	return values
}
