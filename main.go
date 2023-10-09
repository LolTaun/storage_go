package main

import (
	"fmt"
	"map_list/storage"
	"map_list/storage/list"
	"map_list/storage/mp"
)

func main() {
	var l storage.Storage = &list.List{}
	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)
	l.Add(4)
	l.Add(5)
	l.Add(5)
	l.Add(6)
	l.Add(7)
	l.Add(7)

	l.Print()
	l.RemoveByIndex(3)
	l.Print()
	l.RemoveByValue(5)
	l.Print()
	l.RemoveAllByValue(7)
	l.Print()
	l_dump, _ := l.GetAll()
	fmt.Println(l_dump)

	var m storage.Storage = &mp.Map{}
	m.Add(1)
	m.Add(2)
	m.Add(3)
	m.Add(4)
	m.Add(4)
	m.Add(5)
	m.Add(5)
	m.Add(6)
	m.Add(7)
	m.Add(7)

	m.Print()
	m.RemoveByIndex(3)
	m.Print()
	m.RemoveByValue(5)
	m.Print()
	m.RemoveAllByValue(7)
	m.Print()
	m_dump, _ := m.GetAll()
	fmt.Println(m_dump)

	m.Add(8)
	m.Print()

}
