package main

import ()

const tiny = 0.0000001

type Class byte

const (
	Ham Class = iota
	Spam
	MAXCLASS
)

func (c Class) String() string {
	switch c {
	case Ham:
		return "Ham"
	case Spam:
		return "Spam"
	default:
		panic("HELP")
	}
}

type Example struct {
	Document []string
	Class
}

type doc []int

func (d doc) IDs() []int {
	return []int(d)
}

func main() {

}
