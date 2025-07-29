package common

import (
	"context"
	"github.com/gin-gonic/gin"
)

func InvokeUseCase[Input any, Output any](
	GetInput func(*gin.Context) (*Input, error),
	Invoke func(context.Context, *Input) (*Output, error),
	WriteOutput func(*gin.Context, *Output, error),
) gin.HandlerFunc {
	return func(c *gin.Context) {
		input, err := GetInput(c)
		if err != nil {
			// panic(err)
			WriteOutput(c, nil, err)
			return
		}

		output, err := Invoke(c.Request.Context(), input)

		if err != nil {
			// panic(err)
			WriteOutput(c, nil, err)
			return
		}

		WriteOutput(c, output, nil)
	}
}
