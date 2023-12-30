package main

import (
	"fmt"
	"github.com/rotisserie/eris"
)

var (
	// global error values can be useful when wrapping errors or inspecting error types
	ErrInternalServer = eris.New("error internal server")
)

type Request struct {
	ID string
}

func (req Request) Validate() error {
	if req.ID == "" {
		// or return a new error at the source if you prefer
		return eris.Wrapf(ErrInternalServer,
			"error bad request: %v", req)
	}
	return nil
}

func main() {
	err := Request{}.Validate()
	// format the error to a string and print it
	formattedStr := eris.ToString(err, true)
	fmt.Println(formattedStr)
}
