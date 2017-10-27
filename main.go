package main

import (
	"fmt"

	"github.com/juju/errors"
)

var bob int

func ifStatementWithoutErrorx(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil

}

func handledError() error {
	var err error
	result, err := ifStatementWithoutError(12)
	if err != nil {
		return errors.Trace(err)
	}
	fmt.Println(result)
	return nil
}

func untracedError() error {
	var err error
	result, err := ifStatementWithoutError(12)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func unhandledError() error {
	var err error
	result, err := ifStatementWithoutError(42)
	fmt.Println(result)
	fmt.Println(err.Error())
	return nil
}

func unhandledErrorWithIfStatement() error {
	var err error
	_, err = ifStatementWithoutError(23)
	if true {
		fmt.Println("asdf")
		errors.Trace(err)
	}
	return nil
}

func main() {

	// lingo ignore: we don't actually care about this error
	_ = handledError()
	_ = unhandledErrorWithIfStatement()

}
