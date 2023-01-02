package nth

import (
	"errors"
	"fmt"
)

var (
	ErrZeroth     = errors.New("idk whether 0th is valid or not so you just opened a black hole")
	ErrNegativeth = errors.New("what are you a monster you can't have a negative ordinal number")
	ErrNoneth     = errors.New("why would you get the Nth when it's not in the slice")
)

func Nth(i int) string {
	if i < 0 {
		panic(ErrNegativeth)
	}
	switch i {
	case 0:
		panic(ErrZeroth)
	case 1:
		return "1st"
	case 2:
		return "2nd"
	case 3:
		return "3rd"
	}
	return fmt.Sprintf("%dth", i)
}

func GetNth[T comparable](item T, slice []T) string {
	for i := range slice {
		if item == slice[i] {
			return Nth(i + 1)
		}
	}
	panic(ErrNoneth)
}

// Thank you to cpt who provided the correct code,
// but that's not in the spirit of this little project
// so it's staying private
func indexToString(index int) string {
	switch index % 10 {
	case 1:
		return fmt.Sprintf("%dst", index)
	case 2:
		return fmt.Sprintf("%dnd", index)
	case 3:
		return fmt.Sprintf("%drd", index)
	default:
		return fmt.Sprintf("%dth", index)
	}
}
