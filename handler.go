package lab2

import (
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	buf, inpErr := io.ReadAll(ch.Input)
	if inpErr != nil {
		return inpErr
	}

	res, calcErr := CalculatePrefix(string(buf))
	if calcErr != nil {
		return calcErr
	}

	_, outErr := io.WriteString(ch.Output, res)
	if outErr != nil {
		return outErr
	}

	return nil
}
