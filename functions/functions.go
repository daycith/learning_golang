package functions

import "errors"

func myPrivateFunction() {

}

func Sum(arg1, arg2 int) int {
	return arg1 + arg2
}

func Divide(arg1, arg2 int) (int, error) {

	if arg2 == 0 {
		return 0, errors.New("can not divide by 0")
	}
	sum := arg1 + arg2
	return sum, nil
}
