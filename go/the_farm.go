package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (s SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if err == ErrScaleMalfunction {
		if amount >= 0 {
			amount *= 2
		} else {
			return 0, errors.New("negative fodder")
		}
	} else if err != nil {
		return 0, err
	}

	if amount < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows < 0 {
		return 0, SillyNephewError{cows: cows}
	}
	return amount / float64(cows), nil
}
